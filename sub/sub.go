package sub

import "context"

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
	// This is an anonymous function, no function name
	defer func() {
		if err != nil {
			s.Stop()
		}
	}()
	// TODO
	return nil
}

func (s *Server) Stop() error {
	s.cancel()
	// TODO
	return nil
}
