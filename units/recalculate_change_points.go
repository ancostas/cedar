package units

import (
	"context"
	"fmt"
	"sort"

	"github.com/aclements/go-moremath/stats"
	"github.com/evergreen-ci/cedar"
	"github.com/evergreen-ci/cedar/model"
	"github.com/evergreen-ci/cedar/perf"
	"github.com/evergreen-ci/utility"
	"github.com/mongodb/amboy"
	"github.com/mongodb/amboy/dependency"
	"github.com/mongodb/amboy/job"
	"github.com/mongodb/amboy/registry"
	"github.com/mongodb/grip"
	"github.com/mongodb/grip/message"
	"github.com/mongodb/grip/sometimes"
	"github.com/pkg/errors"
)

type recalculateChangePointsJob struct {
	*job.Base           `bson:"metadata" json:"metadata" yaml:"metadata"`
	env                 cedar.Environment
	conf                *model.CedarConfig
	PerformanceResultId model.PerformanceResultSeriesID `bson:"time_series_id" json:"time_series_id" yaml:"time_series_id"`
	changePointDetector perf.PerformanceAnalysisService
}

func init() {
	registry.AddJobType("recalculate-change-points", func() amboy.Job { return makeChangePointsJob() })
}

func makeChangePointsJob() *recalculateChangePointsJob {
	j := &recalculateChangePointsJob{
		Base: &job.Base{
			JobType: amboy.JobType{
				Name:    "recalculate-change-points",
				Version: 1,
			},
		},
		env: cedar.GetEnvironment(),
	}
	j.SetDependency(dependency.NewAlways())
	return j
}

func NewRecalculateChangePointsJob(timeSeriesId model.PerformanceResultSeriesID) amboy.Job {
	j := makeChangePointsJob()
	// Every ten minutes at most
	timestamp := utility.RoundPartOfHour(10)
	j.SetID(fmt.Sprintf("%s.%s.%s.%s.%s.%s", j.JobType.Name, timeSeriesId.Project, timeSeriesId.Variant, timeSeriesId.Task, timeSeriesId.Test, timestamp))
	j.PerformanceResultId = timeSeriesId
	return j
}

func (j *recalculateChangePointsJob) makeMessage(msg string, id model.PerformanceResultSeriesID) message.Fields {
	return message.Fields{
		"job_id":  j.ID(),
		"message": msg,
		"project": id.Project,
		"variant": id.Variant,
		"task":    id.Task,
		"test":    id.Test,
	}
}

func (j *recalculateChangePointsJob) Run(ctx context.Context) {
	defer j.MarkComplete()
	if j.conf == nil {
		j.conf = model.NewCedarConfig(j.env)
	}
	if j.conf.Flags.DisableSignalProcessing {
		grip.InfoWhen(sometimes.Percent(10), j.makeMessage("signal processing is disabled, skipping processing", j.PerformanceResultId))
		return
	}
	if j.env == nil {
		j.env = cedar.GetEnvironment()
	}
	if j.changePointDetector == nil {
		err := j.conf.Find()
		if err != nil {
			j.AddError(errors.Wrap(err, "Unable to get cedar configuration"))
			return
		}
		j.changePointDetector = perf.NewPerformanceAnalysisService(j.conf.ChangeDetector.URI, j.conf.ChangeDetector.User, j.conf.ChangeDetector.Token)
	}
	performanceData, err := model.GetPerformanceData(ctx, j.env, j.PerformanceResultId)
	if err != nil {
		j.AddError(errors.Wrapf(err, "Unable to aggregate time perfData %s", j.PerformanceResultId))
		return
	}
	if performanceData == nil {
		j.AddError(model.MarkPerformanceResultsAsAnalyzed(ctx, j.env, j.PerformanceResultId))
		return
	}
	for _, perfData := range performanceData.Data {
		sort.Slice(perfData.TimeSeries, func(i, j int) bool {
			return perfData.TimeSeries[i].Order < perfData.TimeSeries[j].Order
		})
		series := perf.TimeSeriesModel{
			Project:     performanceData.PerformanceResultId.Project,
			Variant:     performanceData.PerformanceResultId.Variant,
			Task:        performanceData.PerformanceResultId.Task,
			Test:        performanceData.PerformanceResultId.Test,
			Measurement: perfData.Measurement,
		}
		for k, v := range performanceData.PerformanceResultId.Arguments {
			series.Arguments = append(series.Arguments, perf.ArgumentsModel{
				Name:  k,
				Value: v,
			})
		}
		for i, item := range perfData.TimeSeries {
			series.Data = append(series.Data, perf.TimeSeriesDataModel{
				PerformanceResultID: item.PerfResultID,
				Order:               item.Order,
				Value:               item.Value,
				Version:             ,
			})
		}

		result, err := j.changePointDetector.ReportUpdatedTimeSeries(ctx, series)
		if err != nil {
			j.AddError(errors.Wrapf(err, "Unable to detect change points in time perfData %s", j.PerformanceResultId))
			return
		}
	}

}
