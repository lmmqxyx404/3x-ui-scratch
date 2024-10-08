package web

import (
	"context"
)

type Server struct {
	ctx    context.Context
	cancel context.CancelFunc
}

func NewServer() *Server {
	ctx, cancel := context.WithCancel(context.Background())

	return &Server{
		ctx:    ctx,
		cancel: cancel,
	}
}

func (s *Server) Start() (err error) {
	return nil
}
