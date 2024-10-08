package web

import (
	"context"

	"x-ui-scratch/web/service"
)

type Server struct {
	ctx    context.Context
	cancel context.CancelFunc

	settingService service.SettingService
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

	_, err = s.settingService.GetTimeLocation()
	if err != nil {
		return err
	}
	return nil
}

func (s *Server) Stop() error {
	return nil
}
