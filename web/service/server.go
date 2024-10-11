package service

import (
	"time"
	"x-ui-scratch/logger"
)

type Status struct {
	T time.Time `json:"-"`
}

type ServerService struct {
	// xrayService    XrayService
	// inboundService InboundService
}

func (s *ServerService) GetStatus(lastStatus *Status) *Status {
	now := time.Now()
	logger.Info("TODO: GetStatus")
	status := &Status{
		T: now,
	}

	return status
}
