package rest

import (
	"github.com/mongodb/amboy"
	"github.com/pkg/errors"
	"github.com/tychoish/gimlet"
	"github.com/tychoish/grip"
	"github.com/tychoish/sink"
	"golang.org/x/net/context"
)

type Service struct {
	Port int

	// internal settings
	queue amboy.Queue
	app   *gimlet.APIApp
}

func (s *Service) Validate() error {
	var err error

	if s.queue == nil {
		s.queue, err = sink.GetQueue()
		if err != nil {
			return errors.Wrap(err, "problem getting queue")
		}
	}

	if s.app == nil {
		s.app = gimlet.NewApp()
		s.app.SetDefaultVersion(1)
	}

	if s.Port == 0 {
		s.Port = 3000
	}

	if err := s.app.SetPort(s.Port); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (s *Service) Start(ctx context.Context) error {
	if s.queue == nil || s.app == nil {
		return errors.New("application is not valid")
	}

	s.addRoutes()

	if err := s.queue.Start(ctx); err != nil {
		return errors.Wrap(err, "problem starting queue")
	}

	if err := s.app.Resolve(); err != nil {
		return errors.Wrap(err, "problem resolving routes")
	}

	if err := s.app.Run(); err != nil {
		return errors.Wrap(err, "problem running service")
	}

	grip.Noticef("completed sink service; shutting down")

	return nil
}

func (s *Service) addRoutes() {
	s.app.AddRoute("/status").Version(1).Get().Handler(s.statusHandler)
	s.app.AddRoute("/simple_log/{id}").Version(1).Post().Handler(s.simpleLogInjestion)
	s.app.AddRoute("/simple_log/{id}").Version(1).Get().Handler(s.simpleLogRetrieval)
}
