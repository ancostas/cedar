// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/datalabeling/v1beta1/evaluation.proto

package datalabeling

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "google.golang.org/genproto/googleapis/api/annotations"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Describes an evaluation between a machine learning model's predictions and
// ground truth labels. Created when an [EvaluationJob][google.cloud.datalabeling.v1beta1.EvaluationJob] runs successfully.
type Evaluation struct {
	// Output only. Resource name of an evaluation. The name has the following
	// format:
	//
	// "projects/<var>{project_id}</var>/datasets/<var>{dataset_id}</var>/evaluations/<var>{evaluation_id</var>}'
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Output only. Options used in the evaluation job that created this
	// evaluation.
	Config *EvaluationConfig `protobuf:"bytes,2,opt,name=config,proto3" json:"config,omitempty"`
	// Output only. Timestamp for when the evaluation job that created this
	// evaluation ran.
	EvaluationJobRunTime *timestamp.Timestamp `protobuf:"bytes,3,opt,name=evaluation_job_run_time,json=evaluationJobRunTime,proto3" json:"evaluation_job_run_time,omitempty"`
	// Output only. Timestamp for when this evaluation was created.
	CreateTime *timestamp.Timestamp `protobuf:"bytes,4,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// Output only. Metrics comparing predictions to ground truth labels.
	EvaluationMetrics *EvaluationMetrics `protobuf:"bytes,5,opt,name=evaluation_metrics,json=evaluationMetrics,proto3" json:"evaluation_metrics,omitempty"`
	// Output only. Type of task that the model version being evaluated performs,
	// as defined in the
	//
	// [evaluationJobConfig.inputConfig.annotationType][google.cloud.datalabeling.v1beta1.EvaluationJobConfig.input_config]
	// field of the evaluation job that created this evaluation.
	AnnotationType AnnotationType `protobuf:"varint,6,opt,name=annotation_type,json=annotationType,proto3,enum=google.cloud.datalabeling.v1beta1.AnnotationType" json:"annotation_type,omitempty"`
	// Output only. The number of items in the ground truth dataset that were used
	// for this evaluation. Only populated when the evaulation is for certain
	// AnnotationTypes.
	EvaluatedItemCount   int64    `protobuf:"varint,7,opt,name=evaluated_item_count,json=evaluatedItemCount,proto3" json:"evaluated_item_count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Evaluation) Reset()         { *m = Evaluation{} }
func (m *Evaluation) String() string { return proto.CompactTextString(m) }
func (*Evaluation) ProtoMessage()    {}
func (*Evaluation) Descriptor() ([]byte, []int) {
	return fileDescriptor_9d8b95e8e98c043d, []int{0}
}

func (m *Evaluation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Evaluation.Unmarshal(m, b)
}
func (m *Evaluation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Evaluation.Marshal(b, m, deterministic)
}
func (m *Evaluation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Evaluation.Merge(m, src)
}
func (m *Evaluation) XXX_Size() int {
	return xxx_messageInfo_Evaluation.Size(m)
}
func (m *Evaluation) XXX_DiscardUnknown() {
	xxx_messageInfo_Evaluation.DiscardUnknown(m)
}

var xxx_messageInfo_Evaluation proto.InternalMessageInfo

func (m *Evaluation) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Evaluation) GetConfig() *EvaluationConfig {
	if m != nil {
		return m.Config
	}
	return nil
}

func (m *Evaluation) GetEvaluationJobRunTime() *timestamp.Timestamp {
	if m != nil {
		return m.EvaluationJobRunTime
	}
	return nil
}

func (m *Evaluation) GetCreateTime() *timestamp.Timestamp {
	if m != nil {
		return m.CreateTime
	}
	return nil
}

func (m *Evaluation) GetEvaluationMetrics() *EvaluationMetrics {
	if m != nil {
		return m.EvaluationMetrics
	}
	return nil
}

func (m *Evaluation) GetAnnotationType() AnnotationType {
	if m != nil {
		return m.AnnotationType
	}
	return AnnotationType_ANNOTATION_TYPE_UNSPECIFIED
}

func (m *Evaluation) GetEvaluatedItemCount() int64 {
	if m != nil {
		return m.EvaluatedItemCount
	}
	return 0
}

// Configuration details used for calculating evaluation metrics and creating an
// [Evaluation][google.cloud.datalabeling.v1beta1.Evaluation].
type EvaluationConfig struct {
	// Vertical specific options for general metrics.
	//
	// Types that are valid to be assigned to VerticalOption:
	//	*EvaluationConfig_BoundingBoxEvaluationOptions
	VerticalOption       isEvaluationConfig_VerticalOption `protobuf_oneof:"vertical_option"`
	XXX_NoUnkeyedLiteral struct{}                          `json:"-"`
	XXX_unrecognized     []byte                            `json:"-"`
	XXX_sizecache        int32                             `json:"-"`
}

func (m *EvaluationConfig) Reset()         { *m = EvaluationConfig{} }
func (m *EvaluationConfig) String() string { return proto.CompactTextString(m) }
func (*EvaluationConfig) ProtoMessage()    {}
func (*EvaluationConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_9d8b95e8e98c043d, []int{1}
}

func (m *EvaluationConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EvaluationConfig.Unmarshal(m, b)
}
func (m *EvaluationConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EvaluationConfig.Marshal(b, m, deterministic)
}
func (m *EvaluationConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EvaluationConfig.Merge(m, src)
}
func (m *EvaluationConfig) XXX_Size() int {
	return xxx_messageInfo_EvaluationConfig.Size(m)
}
func (m *EvaluationConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_EvaluationConfig.DiscardUnknown(m)
}

var xxx_messageInfo_EvaluationConfig proto.InternalMessageInfo

type isEvaluationConfig_VerticalOption interface {
	isEvaluationConfig_VerticalOption()
}

type EvaluationConfig_BoundingBoxEvaluationOptions struct {
	BoundingBoxEvaluationOptions *BoundingBoxEvaluationOptions `protobuf:"bytes,1,opt,name=bounding_box_evaluation_options,json=boundingBoxEvaluationOptions,proto3,oneof"`
}

func (*EvaluationConfig_BoundingBoxEvaluationOptions) isEvaluationConfig_VerticalOption() {}

func (m *EvaluationConfig) GetVerticalOption() isEvaluationConfig_VerticalOption {
	if m != nil {
		return m.VerticalOption
	}
	return nil
}

func (m *EvaluationConfig) GetBoundingBoxEvaluationOptions() *BoundingBoxEvaluationOptions {
	if x, ok := m.GetVerticalOption().(*EvaluationConfig_BoundingBoxEvaluationOptions); ok {
		return x.BoundingBoxEvaluationOptions
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*EvaluationConfig) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*EvaluationConfig_BoundingBoxEvaluationOptions)(nil),
	}
}

// Options regarding evaluation between bounding boxes.
type BoundingBoxEvaluationOptions struct {
	// Minimum
	// [intersection-over-union
	//
	// (IOU)](/vision/automl/object-detection/docs/evaluate#intersection-over-union)
	// required for 2 bounding boxes to be considered a match. This must be a
	// number between 0 and 1.
	IouThreshold         float32  `protobuf:"fixed32,1,opt,name=iou_threshold,json=iouThreshold,proto3" json:"iou_threshold,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BoundingBoxEvaluationOptions) Reset()         { *m = BoundingBoxEvaluationOptions{} }
func (m *BoundingBoxEvaluationOptions) String() string { return proto.CompactTextString(m) }
func (*BoundingBoxEvaluationOptions) ProtoMessage()    {}
func (*BoundingBoxEvaluationOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_9d8b95e8e98c043d, []int{2}
}

func (m *BoundingBoxEvaluationOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BoundingBoxEvaluationOptions.Unmarshal(m, b)
}
func (m *BoundingBoxEvaluationOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BoundingBoxEvaluationOptions.Marshal(b, m, deterministic)
}
func (m *BoundingBoxEvaluationOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BoundingBoxEvaluationOptions.Merge(m, src)
}
func (m *BoundingBoxEvaluationOptions) XXX_Size() int {
	return xxx_messageInfo_BoundingBoxEvaluationOptions.Size(m)
}
func (m *BoundingBoxEvaluationOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_BoundingBoxEvaluationOptions.DiscardUnknown(m)
}

var xxx_messageInfo_BoundingBoxEvaluationOptions proto.InternalMessageInfo

func (m *BoundingBoxEvaluationOptions) GetIouThreshold() float32 {
	if m != nil {
		return m.IouThreshold
	}
	return 0
}

type EvaluationMetrics struct {
	// Common metrics covering most general cases.
	//
	// Types that are valid to be assigned to Metrics:
	//	*EvaluationMetrics_ClassificationMetrics
	//	*EvaluationMetrics_ObjectDetectionMetrics
	Metrics              isEvaluationMetrics_Metrics `protobuf_oneof:"metrics"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *EvaluationMetrics) Reset()         { *m = EvaluationMetrics{} }
func (m *EvaluationMetrics) String() string { return proto.CompactTextString(m) }
func (*EvaluationMetrics) ProtoMessage()    {}
func (*EvaluationMetrics) Descriptor() ([]byte, []int) {
	return fileDescriptor_9d8b95e8e98c043d, []int{3}
}

func (m *EvaluationMetrics) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EvaluationMetrics.Unmarshal(m, b)
}
func (m *EvaluationMetrics) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EvaluationMetrics.Marshal(b, m, deterministic)
}
func (m *EvaluationMetrics) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EvaluationMetrics.Merge(m, src)
}
func (m *EvaluationMetrics) XXX_Size() int {
	return xxx_messageInfo_EvaluationMetrics.Size(m)
}
func (m *EvaluationMetrics) XXX_DiscardUnknown() {
	xxx_messageInfo_EvaluationMetrics.DiscardUnknown(m)
}

var xxx_messageInfo_EvaluationMetrics proto.InternalMessageInfo

type isEvaluationMetrics_Metrics interface {
	isEvaluationMetrics_Metrics()
}

type EvaluationMetrics_ClassificationMetrics struct {
	ClassificationMetrics *ClassificationMetrics `protobuf:"bytes,1,opt,name=classification_metrics,json=classificationMetrics,proto3,oneof"`
}

type EvaluationMetrics_ObjectDetectionMetrics struct {
	ObjectDetectionMetrics *ObjectDetectionMetrics `protobuf:"bytes,2,opt,name=object_detection_metrics,json=objectDetectionMetrics,proto3,oneof"`
}

func (*EvaluationMetrics_ClassificationMetrics) isEvaluationMetrics_Metrics() {}

func (*EvaluationMetrics_ObjectDetectionMetrics) isEvaluationMetrics_Metrics() {}

func (m *EvaluationMetrics) GetMetrics() isEvaluationMetrics_Metrics {
	if m != nil {
		return m.Metrics
	}
	return nil
}

func (m *EvaluationMetrics) GetClassificationMetrics() *ClassificationMetrics {
	if x, ok := m.GetMetrics().(*EvaluationMetrics_ClassificationMetrics); ok {
		return x.ClassificationMetrics
	}
	return nil
}

func (m *EvaluationMetrics) GetObjectDetectionMetrics() *ObjectDetectionMetrics {
	if x, ok := m.GetMetrics().(*EvaluationMetrics_ObjectDetectionMetrics); ok {
		return x.ObjectDetectionMetrics
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*EvaluationMetrics) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*EvaluationMetrics_ClassificationMetrics)(nil),
		(*EvaluationMetrics_ObjectDetectionMetrics)(nil),
	}
}

// Metrics calculated for a classification model.
type ClassificationMetrics struct {
	// Precision-recall curve based on ground truth labels, predicted labels, and
	// scores for the predicted labels.
	PrCurve *PrCurve `protobuf:"bytes,1,opt,name=pr_curve,json=prCurve,proto3" json:"pr_curve,omitempty"`
	// Confusion matrix of predicted labels vs. ground truth labels.
	ConfusionMatrix      *ConfusionMatrix `protobuf:"bytes,2,opt,name=confusion_matrix,json=confusionMatrix,proto3" json:"confusion_matrix,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *ClassificationMetrics) Reset()         { *m = ClassificationMetrics{} }
func (m *ClassificationMetrics) String() string { return proto.CompactTextString(m) }
func (*ClassificationMetrics) ProtoMessage()    {}
func (*ClassificationMetrics) Descriptor() ([]byte, []int) {
	return fileDescriptor_9d8b95e8e98c043d, []int{4}
}

func (m *ClassificationMetrics) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClassificationMetrics.Unmarshal(m, b)
}
func (m *ClassificationMetrics) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClassificationMetrics.Marshal(b, m, deterministic)
}
func (m *ClassificationMetrics) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClassificationMetrics.Merge(m, src)
}
func (m *ClassificationMetrics) XXX_Size() int {
	return xxx_messageInfo_ClassificationMetrics.Size(m)
}
func (m *ClassificationMetrics) XXX_DiscardUnknown() {
	xxx_messageInfo_ClassificationMetrics.DiscardUnknown(m)
}

var xxx_messageInfo_ClassificationMetrics proto.InternalMessageInfo

func (m *ClassificationMetrics) GetPrCurve() *PrCurve {
	if m != nil {
		return m.PrCurve
	}
	return nil
}

func (m *ClassificationMetrics) GetConfusionMatrix() *ConfusionMatrix {
	if m != nil {
		return m.ConfusionMatrix
	}
	return nil
}

// Metrics calculated for an image object detection (bounding box) model.
type ObjectDetectionMetrics struct {
	// Precision-recall curve.
	PrCurve              *PrCurve `protobuf:"bytes,1,opt,name=pr_curve,json=prCurve,proto3" json:"pr_curve,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ObjectDetectionMetrics) Reset()         { *m = ObjectDetectionMetrics{} }
func (m *ObjectDetectionMetrics) String() string { return proto.CompactTextString(m) }
func (*ObjectDetectionMetrics) ProtoMessage()    {}
func (*ObjectDetectionMetrics) Descriptor() ([]byte, []int) {
	return fileDescriptor_9d8b95e8e98c043d, []int{5}
}

func (m *ObjectDetectionMetrics) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ObjectDetectionMetrics.Unmarshal(m, b)
}
func (m *ObjectDetectionMetrics) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ObjectDetectionMetrics.Marshal(b, m, deterministic)
}
func (m *ObjectDetectionMetrics) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ObjectDetectionMetrics.Merge(m, src)
}
func (m *ObjectDetectionMetrics) XXX_Size() int {
	return xxx_messageInfo_ObjectDetectionMetrics.Size(m)
}
func (m *ObjectDetectionMetrics) XXX_DiscardUnknown() {
	xxx_messageInfo_ObjectDetectionMetrics.DiscardUnknown(m)
}

var xxx_messageInfo_ObjectDetectionMetrics proto.InternalMessageInfo

func (m *ObjectDetectionMetrics) GetPrCurve() *PrCurve {
	if m != nil {
		return m.PrCurve
	}
	return nil
}

type PrCurve struct {
	// The annotation spec of the label for which the precision-recall curve
	// calculated. If this field is empty, that means the precision-recall curve
	// is an aggregate curve for all labels.
	AnnotationSpec *AnnotationSpec `protobuf:"bytes,1,opt,name=annotation_spec,json=annotationSpec,proto3" json:"annotation_spec,omitempty"`
	// Area under the precision-recall curve. Not to be confused with area under
	// a receiver operating characteristic (ROC) curve.
	AreaUnderCurve float32 `protobuf:"fixed32,2,opt,name=area_under_curve,json=areaUnderCurve,proto3" json:"area_under_curve,omitempty"`
	// Entries that make up the precision-recall graph. Each entry is a "point" on
	// the graph drawn for a different `confidence_threshold`.
	ConfidenceMetricsEntries []*PrCurve_ConfidenceMetricsEntry `protobuf:"bytes,3,rep,name=confidence_metrics_entries,json=confidenceMetricsEntries,proto3" json:"confidence_metrics_entries,omitempty"`
	// Mean average prcision of this curve.
	MeanAveragePrecision float32  `protobuf:"fixed32,4,opt,name=mean_average_precision,json=meanAveragePrecision,proto3" json:"mean_average_precision,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PrCurve) Reset()         { *m = PrCurve{} }
func (m *PrCurve) String() string { return proto.CompactTextString(m) }
func (*PrCurve) ProtoMessage()    {}
func (*PrCurve) Descriptor() ([]byte, []int) {
	return fileDescriptor_9d8b95e8e98c043d, []int{6}
}

func (m *PrCurve) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PrCurve.Unmarshal(m, b)
}
func (m *PrCurve) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PrCurve.Marshal(b, m, deterministic)
}
func (m *PrCurve) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PrCurve.Merge(m, src)
}
func (m *PrCurve) XXX_Size() int {
	return xxx_messageInfo_PrCurve.Size(m)
}
func (m *PrCurve) XXX_DiscardUnknown() {
	xxx_messageInfo_PrCurve.DiscardUnknown(m)
}

var xxx_messageInfo_PrCurve proto.InternalMessageInfo

func (m *PrCurve) GetAnnotationSpec() *AnnotationSpec {
	if m != nil {
		return m.AnnotationSpec
	}
	return nil
}

func (m *PrCurve) GetAreaUnderCurve() float32 {
	if m != nil {
		return m.AreaUnderCurve
	}
	return 0
}

func (m *PrCurve) GetConfidenceMetricsEntries() []*PrCurve_ConfidenceMetricsEntry {
	if m != nil {
		return m.ConfidenceMetricsEntries
	}
	return nil
}

func (m *PrCurve) GetMeanAveragePrecision() float32 {
	if m != nil {
		return m.MeanAveragePrecision
	}
	return 0
}

type PrCurve_ConfidenceMetricsEntry struct {
	// Threshold used for this entry.
	//
	// For classification tasks, this is a classification threshold: a
	// predicted label is categorized as positive or negative (in the context of
	// this point on the PR curve) based on whether the label's score meets this
	// threshold.
	//
	// For image object detection (bounding box) tasks, this is the
	// [intersection-over-union
	//
	// (IOU)](/vision/automl/object-detection/docs/evaluate#intersection-over-union)
	// threshold for the context of this point on the PR curve.
	ConfidenceThreshold float32 `protobuf:"fixed32,1,opt,name=confidence_threshold,json=confidenceThreshold,proto3" json:"confidence_threshold,omitempty"`
	// Recall value.
	Recall float32 `protobuf:"fixed32,2,opt,name=recall,proto3" json:"recall,omitempty"`
	// Precision value.
	Precision float32 `protobuf:"fixed32,3,opt,name=precision,proto3" json:"precision,omitempty"`
	// Harmonic mean of recall and precision.
	F1Score float32 `protobuf:"fixed32,4,opt,name=f1_score,json=f1Score,proto3" json:"f1_score,omitempty"`
	// Recall value for entries with label that has highest score.
	RecallAt1 float32 `protobuf:"fixed32,5,opt,name=recall_at1,json=recallAt1,proto3" json:"recall_at1,omitempty"`
	// Precision value for entries with label that has highest score.
	PrecisionAt1 float32 `protobuf:"fixed32,6,opt,name=precision_at1,json=precisionAt1,proto3" json:"precision_at1,omitempty"`
	// The harmonic mean of [recall_at1][google.cloud.datalabeling.v1beta1.PrCurve.ConfidenceMetricsEntry.recall_at1] and [precision_at1][google.cloud.datalabeling.v1beta1.PrCurve.ConfidenceMetricsEntry.precision_at1].
	F1ScoreAt1 float32 `protobuf:"fixed32,7,opt,name=f1_score_at1,json=f1ScoreAt1,proto3" json:"f1_score_at1,omitempty"`
	// Recall value for entries with label that has highest 5 scores.
	RecallAt5 float32 `protobuf:"fixed32,8,opt,name=recall_at5,json=recallAt5,proto3" json:"recall_at5,omitempty"`
	// Precision value for entries with label that has highest 5 scores.
	PrecisionAt5 float32 `protobuf:"fixed32,9,opt,name=precision_at5,json=precisionAt5,proto3" json:"precision_at5,omitempty"`
	// The harmonic mean of [recall_at5][google.cloud.datalabeling.v1beta1.PrCurve.ConfidenceMetricsEntry.recall_at5] and [precision_at5][google.cloud.datalabeling.v1beta1.PrCurve.ConfidenceMetricsEntry.precision_at5].
	F1ScoreAt5           float32  `protobuf:"fixed32,10,opt,name=f1_score_at5,json=f1ScoreAt5,proto3" json:"f1_score_at5,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PrCurve_ConfidenceMetricsEntry) Reset()         { *m = PrCurve_ConfidenceMetricsEntry{} }
func (m *PrCurve_ConfidenceMetricsEntry) String() string { return proto.CompactTextString(m) }
func (*PrCurve_ConfidenceMetricsEntry) ProtoMessage()    {}
func (*PrCurve_ConfidenceMetricsEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_9d8b95e8e98c043d, []int{6, 0}
}

func (m *PrCurve_ConfidenceMetricsEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PrCurve_ConfidenceMetricsEntry.Unmarshal(m, b)
}
func (m *PrCurve_ConfidenceMetricsEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PrCurve_ConfidenceMetricsEntry.Marshal(b, m, deterministic)
}
func (m *PrCurve_ConfidenceMetricsEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PrCurve_ConfidenceMetricsEntry.Merge(m, src)
}
func (m *PrCurve_ConfidenceMetricsEntry) XXX_Size() int {
	return xxx_messageInfo_PrCurve_ConfidenceMetricsEntry.Size(m)
}
func (m *PrCurve_ConfidenceMetricsEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_PrCurve_ConfidenceMetricsEntry.DiscardUnknown(m)
}

var xxx_messageInfo_PrCurve_ConfidenceMetricsEntry proto.InternalMessageInfo

func (m *PrCurve_ConfidenceMetricsEntry) GetConfidenceThreshold() float32 {
	if m != nil {
		return m.ConfidenceThreshold
	}
	return 0
}

func (m *PrCurve_ConfidenceMetricsEntry) GetRecall() float32 {
	if m != nil {
		return m.Recall
	}
	return 0
}

func (m *PrCurve_ConfidenceMetricsEntry) GetPrecision() float32 {
	if m != nil {
		return m.Precision
	}
	return 0
}

func (m *PrCurve_ConfidenceMetricsEntry) GetF1Score() float32 {
	if m != nil {
		return m.F1Score
	}
	return 0
}

func (m *PrCurve_ConfidenceMetricsEntry) GetRecallAt1() float32 {
	if m != nil {
		return m.RecallAt1
	}
	return 0
}

func (m *PrCurve_ConfidenceMetricsEntry) GetPrecisionAt1() float32 {
	if m != nil {
		return m.PrecisionAt1
	}
	return 0
}

func (m *PrCurve_ConfidenceMetricsEntry) GetF1ScoreAt1() float32 {
	if m != nil {
		return m.F1ScoreAt1
	}
	return 0
}

func (m *PrCurve_ConfidenceMetricsEntry) GetRecallAt5() float32 {
	if m != nil {
		return m.RecallAt5
	}
	return 0
}

func (m *PrCurve_ConfidenceMetricsEntry) GetPrecisionAt5() float32 {
	if m != nil {
		return m.PrecisionAt5
	}
	return 0
}

func (m *PrCurve_ConfidenceMetricsEntry) GetF1ScoreAt5() float32 {
	if m != nil {
		return m.F1ScoreAt5
	}
	return 0
}

// Confusion matrix of the model running the classification. Only applicable
// when the metrics entry aggregates multiple labels. Not applicable when the
// entry is for a single label.
type ConfusionMatrix struct {
	Row                  []*ConfusionMatrix_Row `protobuf:"bytes,1,rep,name=row,proto3" json:"row,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *ConfusionMatrix) Reset()         { *m = ConfusionMatrix{} }
func (m *ConfusionMatrix) String() string { return proto.CompactTextString(m) }
func (*ConfusionMatrix) ProtoMessage()    {}
func (*ConfusionMatrix) Descriptor() ([]byte, []int) {
	return fileDescriptor_9d8b95e8e98c043d, []int{7}
}

func (m *ConfusionMatrix) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfusionMatrix.Unmarshal(m, b)
}
func (m *ConfusionMatrix) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfusionMatrix.Marshal(b, m, deterministic)
}
func (m *ConfusionMatrix) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfusionMatrix.Merge(m, src)
}
func (m *ConfusionMatrix) XXX_Size() int {
	return xxx_messageInfo_ConfusionMatrix.Size(m)
}
func (m *ConfusionMatrix) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfusionMatrix.DiscardUnknown(m)
}

var xxx_messageInfo_ConfusionMatrix proto.InternalMessageInfo

func (m *ConfusionMatrix) GetRow() []*ConfusionMatrix_Row {
	if m != nil {
		return m.Row
	}
	return nil
}

type ConfusionMatrix_ConfusionMatrixEntry struct {
	// The annotation spec of a predicted label.
	AnnotationSpec *AnnotationSpec `protobuf:"bytes,1,opt,name=annotation_spec,json=annotationSpec,proto3" json:"annotation_spec,omitempty"`
	// Number of items predicted to have this label. (The ground truth label for
	// these items is the `Row.annotationSpec` of this entry's parent.)
	ItemCount            int32    `protobuf:"varint,2,opt,name=item_count,json=itemCount,proto3" json:"item_count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConfusionMatrix_ConfusionMatrixEntry) Reset()         { *m = ConfusionMatrix_ConfusionMatrixEntry{} }
func (m *ConfusionMatrix_ConfusionMatrixEntry) String() string { return proto.CompactTextString(m) }
func (*ConfusionMatrix_ConfusionMatrixEntry) ProtoMessage()    {}
func (*ConfusionMatrix_ConfusionMatrixEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_9d8b95e8e98c043d, []int{7, 0}
}

func (m *ConfusionMatrix_ConfusionMatrixEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfusionMatrix_ConfusionMatrixEntry.Unmarshal(m, b)
}
func (m *ConfusionMatrix_ConfusionMatrixEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfusionMatrix_ConfusionMatrixEntry.Marshal(b, m, deterministic)
}
func (m *ConfusionMatrix_ConfusionMatrixEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfusionMatrix_ConfusionMatrixEntry.Merge(m, src)
}
func (m *ConfusionMatrix_ConfusionMatrixEntry) XXX_Size() int {
	return xxx_messageInfo_ConfusionMatrix_ConfusionMatrixEntry.Size(m)
}
func (m *ConfusionMatrix_ConfusionMatrixEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfusionMatrix_ConfusionMatrixEntry.DiscardUnknown(m)
}

var xxx_messageInfo_ConfusionMatrix_ConfusionMatrixEntry proto.InternalMessageInfo

func (m *ConfusionMatrix_ConfusionMatrixEntry) GetAnnotationSpec() *AnnotationSpec {
	if m != nil {
		return m.AnnotationSpec
	}
	return nil
}

func (m *ConfusionMatrix_ConfusionMatrixEntry) GetItemCount() int32 {
	if m != nil {
		return m.ItemCount
	}
	return 0
}

// A row in the confusion matrix. Each entry in this row has the same
// ground truth label.
type ConfusionMatrix_Row struct {
	// The annotation spec of the ground truth label for this row.
	AnnotationSpec *AnnotationSpec `protobuf:"bytes,1,opt,name=annotation_spec,json=annotationSpec,proto3" json:"annotation_spec,omitempty"`
	// A list of the confusion matrix entries. One entry for each possible
	// predicted label.
	Entries              []*ConfusionMatrix_ConfusionMatrixEntry `protobuf:"bytes,2,rep,name=entries,proto3" json:"entries,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                `json:"-"`
	XXX_unrecognized     []byte                                  `json:"-"`
	XXX_sizecache        int32                                   `json:"-"`
}

func (m *ConfusionMatrix_Row) Reset()         { *m = ConfusionMatrix_Row{} }
func (m *ConfusionMatrix_Row) String() string { return proto.CompactTextString(m) }
func (*ConfusionMatrix_Row) ProtoMessage()    {}
func (*ConfusionMatrix_Row) Descriptor() ([]byte, []int) {
	return fileDescriptor_9d8b95e8e98c043d, []int{7, 1}
}

func (m *ConfusionMatrix_Row) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfusionMatrix_Row.Unmarshal(m, b)
}
func (m *ConfusionMatrix_Row) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfusionMatrix_Row.Marshal(b, m, deterministic)
}
func (m *ConfusionMatrix_Row) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfusionMatrix_Row.Merge(m, src)
}
func (m *ConfusionMatrix_Row) XXX_Size() int {
	return xxx_messageInfo_ConfusionMatrix_Row.Size(m)
}
func (m *ConfusionMatrix_Row) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfusionMatrix_Row.DiscardUnknown(m)
}

var xxx_messageInfo_ConfusionMatrix_Row proto.InternalMessageInfo

func (m *ConfusionMatrix_Row) GetAnnotationSpec() *AnnotationSpec {
	if m != nil {
		return m.AnnotationSpec
	}
	return nil
}

func (m *ConfusionMatrix_Row) GetEntries() []*ConfusionMatrix_ConfusionMatrixEntry {
	if m != nil {
		return m.Entries
	}
	return nil
}

func init() {
	proto.RegisterType((*Evaluation)(nil), "google.cloud.datalabeling.v1beta1.Evaluation")
	proto.RegisterType((*EvaluationConfig)(nil), "google.cloud.datalabeling.v1beta1.EvaluationConfig")
	proto.RegisterType((*BoundingBoxEvaluationOptions)(nil), "google.cloud.datalabeling.v1beta1.BoundingBoxEvaluationOptions")
	proto.RegisterType((*EvaluationMetrics)(nil), "google.cloud.datalabeling.v1beta1.EvaluationMetrics")
	proto.RegisterType((*ClassificationMetrics)(nil), "google.cloud.datalabeling.v1beta1.ClassificationMetrics")
	proto.RegisterType((*ObjectDetectionMetrics)(nil), "google.cloud.datalabeling.v1beta1.ObjectDetectionMetrics")
	proto.RegisterType((*PrCurve)(nil), "google.cloud.datalabeling.v1beta1.PrCurve")
	proto.RegisterType((*PrCurve_ConfidenceMetricsEntry)(nil), "google.cloud.datalabeling.v1beta1.PrCurve.ConfidenceMetricsEntry")
	proto.RegisterType((*ConfusionMatrix)(nil), "google.cloud.datalabeling.v1beta1.ConfusionMatrix")
	proto.RegisterType((*ConfusionMatrix_ConfusionMatrixEntry)(nil), "google.cloud.datalabeling.v1beta1.ConfusionMatrix.ConfusionMatrixEntry")
	proto.RegisterType((*ConfusionMatrix_Row)(nil), "google.cloud.datalabeling.v1beta1.ConfusionMatrix.Row")
}

func init() {
	proto.RegisterFile("google/cloud/datalabeling/v1beta1/evaluation.proto", fileDescriptor_9d8b95e8e98c043d)
}

var fileDescriptor_9d8b95e8e98c043d = []byte{
	// 1013 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x56, 0x4f, 0x6f, 0x1b, 0x45,
	0x14, 0xef, 0xda, 0x4d, 0x1c, 0xbf, 0x94, 0xfc, 0x19, 0x52, 0xb3, 0xb5, 0x52, 0xd5, 0x18, 0x81,
	0x2c, 0x0e, 0x5e, 0xec, 0xd6, 0x08, 0x88, 0x04, 0x4a, 0x42, 0x44, 0x00, 0x45, 0x0d, 0xd3, 0x70,
	0xa9, 0x84, 0x56, 0xe3, 0xf1, 0x8b, 0xb3, 0x65, 0xbd, 0xb3, 0xcc, 0xce, 0x3a, 0x89, 0xaa, 0x88,
	0x6b, 0xaf, 0x7c, 0x13, 0xb8, 0x70, 0xe3, 0xc6, 0xd7, 0xe0, 0x8e, 0xf8, 0x14, 0x68, 0x67, 0x67,
	0xbd, 0x6b, 0xd7, 0x4a, 0x1c, 0x44, 0x6f, 0x33, 0xef, 0xcd, 0xef, 0xf7, 0x7b, 0xf3, 0xe6, 0xbd,
	0xb7, 0x0b, 0xdd, 0xa1, 0x10, 0x43, 0x1f, 0x1d, 0xee, 0x8b, 0x78, 0xe0, 0x0c, 0x98, 0x62, 0x3e,
	0xeb, 0xa3, 0xef, 0x05, 0x43, 0x67, 0xdc, 0xe9, 0xa3, 0x62, 0x1d, 0x07, 0xc7, 0xcc, 0x8f, 0x99,
	0xf2, 0x44, 0xd0, 0x0e, 0xa5, 0x50, 0x82, 0xbc, 0x9b, 0x62, 0xda, 0x1a, 0xd3, 0x2e, 0x62, 0xda,
	0x06, 0x53, 0xdf, 0x36, 0xb4, 0x2c, 0xf4, 0x1c, 0x16, 0x04, 0x42, 0x69, 0x7c, 0x94, 0x12, 0xd4,
	0x1f, 0x14, 0xbc, 0x12, 0x23, 0x11, 0x4b, 0x8e, 0xc6, 0xb5, 0x40, 0x3c, 0x39, 0x9f, 0xc1, 0xec,
	0xdc, 0x06, 0xe3, 0x46, 0x21, 0x72, 0x37, 0x42, 0x65, 0xc0, 0x8f, 0x0c, 0x58, 0xef, 0xfa, 0xf1,
	0xa9, 0xa3, 0xbc, 0x11, 0x46, 0x8a, 0x8d, 0xc2, 0xf4, 0x40, 0xf3, 0xaf, 0xbb, 0x00, 0x07, 0x93,
	0x14, 0x10, 0x02, 0x77, 0x03, 0x36, 0x42, 0xdb, 0x6a, 0x58, 0xad, 0x2a, 0xd5, 0x6b, 0xf2, 0x2d,
	0x2c, 0x73, 0x11, 0x9c, 0x7a, 0x43, 0xbb, 0xd4, 0xb0, 0x5a, 0xab, 0xdd, 0xc7, 0xed, 0x1b, 0x33,
	0xd4, 0xce, 0x29, 0xf7, 0x35, 0x94, 0x1a, 0x0a, 0xf2, 0x1d, 0xbc, 0x93, 0x67, 0xdc, 0x7d, 0x21,
	0xfa, 0xae, 0x8c, 0x03, 0x37, 0x89, 0xca, 0x2e, 0x6b, 0xf6, 0x7a, 0xc6, 0x9e, 0x85, 0xdc, 0x3e,
	0xc9, 0x42, 0xa6, 0x5b, 0x39, 0xf4, 0x1b, 0xd1, 0xa7, 0x71, 0x90, 0xb8, 0xc8, 0x0e, 0xac, 0x72,
	0x89, 0x4c, 0x61, 0x4a, 0x73, 0xf7, 0x46, 0x1a, 0x48, 0x8f, 0x6b, 0x30, 0x07, 0x52, 0x88, 0x67,
	0x84, 0x4a, 0x7a, 0x3c, 0xb2, 0x97, 0x34, 0xc7, 0x93, 0x5b, 0x5d, 0xf4, 0x28, 0xc5, 0xd2, 0x4d,
	0x9c, 0x35, 0x91, 0xe7, 0xb0, 0x5e, 0x78, 0x22, 0x75, 0x19, 0xa2, 0xbd, 0xdc, 0xb0, 0x5a, 0x6b,
	0xdd, 0xce, 0x02, 0x0a, 0xbb, 0x13, 0xe4, 0xc9, 0x65, 0x88, 0x74, 0x8d, 0x4d, 0xed, 0xc9, 0x47,
	0x90, 0x65, 0x05, 0x07, 0xae, 0xa7, 0x70, 0xe4, 0x72, 0x11, 0x07, 0xca, 0xae, 0x34, 0xac, 0x56,
	0x99, 0x92, 0x89, 0xef, 0x6b, 0x85, 0xa3, 0xfd, 0xc4, 0xf3, 0xd9, 0x8f, 0xff, 0xec, 0x9e, 0xc1,
	0x07, 0x53, 0x5a, 0x69, 0x14, 0x2c, 0xf4, 0xa2, 0x36, 0x17, 0x23, 0xa7, 0x50, 0x10, 0x9f, 0x87,
	0x52, 0xbc, 0x40, 0xae, 0x22, 0xe7, 0xa5, 0x59, 0x5d, 0xe9, 0x1a, 0x8c, 0x30, 0x31, 0x99, 0xd5,
	0x55, 0xa1, 0x8b, 0x22, 0xe7, 0x65, 0xbe, 0xb9, 0x6a, 0xfe, 0x6a, 0xc1, 0xc6, 0x6c, 0x31, 0x90,
	0x57, 0x16, 0x3c, 0xea, 0x8b, 0x38, 0x18, 0x78, 0xc1, 0xd0, 0xed, 0x8b, 0x0b, 0xb7, 0xf0, 0x04,
	0x22, 0xd4, 0x2c, 0xba, 0x02, 0x57, 0xbb, 0x5f, 0x2c, 0x90, 0xa0, 0x3d, 0xc3, 0xb4, 0x27, 0x2e,
	0x72, 0xa5, 0xa7, 0x29, 0xcd, 0xe1, 0x1d, 0xba, 0xdd, 0xbf, 0xc6, 0xbf, 0xb7, 0x09, 0xeb, 0x63,
	0x94, 0xca, 0xe3, 0xcc, 0x37, 0xd2, 0xcd, 0x7d, 0xd8, 0xbe, 0x8e, 0x92, 0xbc, 0x07, 0x6f, 0x79,
	0x22, 0x76, 0xd5, 0x99, 0xc4, 0xe8, 0x4c, 0xf8, 0x03, 0x1d, 0x6a, 0x89, 0xde, 0xf3, 0x44, 0x7c,
	0x92, 0xd9, 0x9a, 0xaf, 0x4a, 0xb0, 0xf9, 0x5a, 0x6d, 0x90, 0x9f, 0xa0, 0xc6, 0x7d, 0x16, 0x45,
	0xde, 0xa9, 0xc7, 0xa7, 0x2b, 0x2e, 0xbd, 0xee, 0x27, 0x0b, 0x5c, 0x77, 0x7f, 0x8a, 0xc0, 0x30,
	0x1f, 0xde, 0xa1, 0xf7, 0xf9, 0x3c, 0x07, 0x89, 0xc1, 0x16, 0xfd, 0xe4, 0xdd, 0xdc, 0x01, 0x2a,
	0xe4, 0x53, 0xa2, 0x69, 0x3f, 0x7f, 0xba, 0x80, 0xe8, 0x53, 0x4d, 0xf1, 0x65, 0xc6, 0x90, 0xab,
	0xd6, 0xc4, 0x5c, 0xcf, 0x5e, 0x15, 0x2a, 0x46, 0xa5, 0xf9, 0x87, 0x05, 0xf7, 0xe7, 0x06, 0x4d,
	0x0e, 0x60, 0x25, 0x94, 0x2e, 0x8f, 0xe5, 0x18, 0x4d, 0x02, 0x3e, 0x5c, 0x20, 0x96, 0x63, 0xb9,
	0x9f, 0x20, 0x68, 0x25, 0x4c, 0x17, 0xe4, 0x07, 0xd8, 0x48, 0xa6, 0x4b, 0x1c, 0xe9, 0xbb, 0x31,
	0x25, 0xbd, 0x0b, 0x73, 0xb5, 0xee, 0x22, 0xf9, 0xcc, 0xa0, 0x47, 0x1a, 0x49, 0xd7, 0xf9, 0xb4,
	0xa1, 0xe9, 0x42, 0x6d, 0xfe, 0xf5, 0xff, 0xa7, 0xf8, 0x9b, 0xbf, 0x2d, 0x41, 0xc5, 0x18, 0x67,
	0x46, 0x45, 0x32, 0xcd, 0x0d, 0xf3, 0xed, 0x46, 0xc5, 0xb3, 0x10, 0x79, 0x71, 0x54, 0x24, 0x7b,
	0xd2, 0x82, 0x0d, 0x26, 0x91, 0xb9, 0x71, 0x30, 0xc0, 0x2c, 0xec, 0x92, 0xae, 0xdd, 0xb5, 0xc4,
	0xfe, 0x7d, 0x62, 0x4e, 0xa3, 0xf8, 0x19, 0xea, 0x7a, 0x5e, 0x0f, 0x30, 0xe0, 0x98, 0x95, 0x8b,
	0x8b, 0x81, 0x92, 0x1e, 0x46, 0x76, 0xb9, 0x51, 0x6e, 0xad, 0x76, 0x77, 0x17, 0xbf, 0xaa, 0xce,
	0x71, 0x4a, 0x66, 0x52, 0x77, 0x10, 0x28, 0x79, 0x49, 0x6d, 0x3e, 0xcf, 0xee, 0x61, 0x44, 0x9e,
	0x40, 0x6d, 0x84, 0x2c, 0x70, 0xd9, 0x18, 0x25, 0x1b, 0xa2, 0x1b, 0x4a, 0xe4, 0x5e, 0xf2, 0x26,
	0x7a, 0xbc, 0x97, 0xe8, 0x56, 0xe2, 0xdd, 0x4d, 0x9d, 0xc7, 0x99, 0xaf, 0xfe, 0x77, 0x09, 0x6a,
	0xf3, 0xa5, 0x48, 0x07, 0xb6, 0x0a, 0x37, 0x9a, 0xed, 0xdd, 0xb7, 0x73, 0xdf, 0xa4, 0x85, 0x49,
	0x0d, 0x96, 0x25, 0x72, 0xe6, 0xfb, 0x26, 0x49, 0x66, 0x47, 0xb6, 0xa1, 0x9a, 0x87, 0x53, 0xd6,
	0xae, 0xdc, 0x40, 0x1e, 0xc0, 0xca, 0x69, 0xc7, 0x8d, 0xb8, 0x90, 0x68, 0x62, 0xad, 0x9c, 0x76,
	0x9e, 0x25, 0x5b, 0xf2, 0x10, 0x20, 0xa5, 0x70, 0x99, 0xea, 0xe8, 0x6f, 0x4c, 0x89, 0x56, 0x53,
	0xcb, 0xae, 0xea, 0x24, 0x73, 0x65, 0x42, 0xa3, 0x4f, 0x2c, 0xa7, 0x73, 0x65, 0x62, 0x4c, 0x0e,
	0x35, 0xe0, 0x5e, 0x46, 0xaf, 0xcf, 0x54, 0xf4, 0x19, 0x30, 0x12, 0xc9, 0x89, 0xa2, 0x4a, 0xcf,
	0x5e, 0x99, 0x56, 0xe9, 0xcd, 0xaa, 0xf4, 0xec, 0xea, 0x6b, 0x2a, 0xbd, 0x19, 0x95, 0x9e, 0x0d,
	0x33, 0x2a, 0xbd, 0xe6, 0xef, 0x65, 0x58, 0x9f, 0xe9, 0x1c, 0x72, 0x08, 0x65, 0x29, 0xce, 0x6d,
	0x4b, 0x97, 0xc7, 0xc7, 0xb7, 0x6f, 0xbd, 0x36, 0x15, 0xe7, 0x34, 0xa1, 0xa8, 0xff, 0x62, 0xc1,
	0xd6, 0x8c, 0x33, 0x7d, 0xc6, 0x37, 0xd9, 0x1e, 0x0f, 0x01, 0x0a, 0xdf, 0xcf, 0xe4, 0xcd, 0x97,
	0x68, 0xd5, 0xcb, 0x3e, 0x9b, 0xf5, 0x3f, 0x2d, 0x28, 0x53, 0x71, 0xfe, 0x46, 0x43, 0x60, 0x50,
	0xc9, 0x9a, 0xac, 0xa4, 0xb3, 0xf8, 0xd5, 0x7f, 0xc8, 0xe2, 0xbc, 0xc4, 0xd1, 0x8c, 0x77, 0xef,
	0x02, 0xde, 0xe7, 0x62, 0x74, 0x33, 0xed, 0xb1, 0xf5, 0xfc, 0xc8, 0x1c, 0x1a, 0x0a, 0x9f, 0x05,
	0xc3, 0xb6, 0x90, 0x43, 0x67, 0x88, 0x81, 0xfe, 0xa1, 0x72, 0xf2, 0x3f, 0x86, 0x6b, 0x7e, 0x4c,
	0x77, 0x8a, 0xc6, 0xfe, 0xb2, 0x46, 0x3e, 0xfe, 0x37, 0x00, 0x00, 0xff, 0xff, 0x31, 0x5d, 0x23,
	0xd3, 0x95, 0x0b, 0x00, 0x00,
}