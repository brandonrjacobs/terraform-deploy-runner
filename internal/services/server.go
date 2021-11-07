package services

import (
	"context"
	"deploy-runner/internal"
	"github.com/go-chi/chi/v5"
	"net"
)

const (
	networkProtocol = "tcp"
)

type server struct {
	log        internal.BackgroundLog
	router     *chi.Mux
	address    string
}

func (s *server) Start(ctx context.Context) error {
	lis, err := net.Listen(networkProtocol, s.address)
	if err != nil {
		s.log.ErrorfCtx(ctx, "Unable to initialize listener for tracksapi server")
		return err
	}
	s.log.Infof("Starting grpc server on port: %s", lis.Addr().String())
	return nil
}

func (s *server) Stop(ctx context.Context) error {
	return nil
}

func (s *server) Disabled() bool {
	return false
}
