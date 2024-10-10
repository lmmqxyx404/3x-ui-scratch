package web

import (
	"context"

	"x-ui-scratch/web/service"

	"github.com/robfig/cron/v3"
)

type Server struct {
	ctx    context.Context
	cancel context.CancelFunc

	settingService service.SettingService

	cron *cron.Cron
}

func NewServer() *Server {
	ctx, cancel := context.WithCancel(context.Background())

	return &Server{
		ctx:    ctx,
		cancel: cancel,
	}
}

func (s *Server) Start() (err error) {
	// This is an anonymous function, no function name
	defer func() {
		if err != nil {
			s.Stop()
		}
	}()

	loc, err := s.settingService.GetTimeLocation()
	if err != nil {
		return err
	}

	s.cron = cron.New(cron.WithLocation(loc), cron.WithSeconds())
	s.cron.Start()

	return nil
}

func (s *Server) Stop() error {
	return nil
}
