package web

import (
	"context"
	"io"

	"x-ui-scratch/config"
	"x-ui-scratch/logger"
	"x-ui-scratch/web/service"

	"github.com/gin-gonic/gin"
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

	_, err = s.initRouter()
	if err != nil {
		return err
	}

	return nil
}

func (s *Server) Stop() error {
	return nil
}

func (s *Server) initRouter() (*gin.Engine, error) {
	logger.Info("initRouter")
	if config.IsDebug() {
		logger.Info("debug mode")

		// gin.SetMode(gin.DebugMode)
	} else {
		logger.Info("release mode")
		gin.DefaultWriter = io.Discard
		gin.DefaultErrorWriter = io.Discard
		gin.SetMode(gin.ReleaseMode)
	}

	engine := gin.Default()

	return engine, nil
}
