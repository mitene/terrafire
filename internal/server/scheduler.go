package server

import (
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	"github.com/mitene/terrafire/internal/api"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
)

type Scheduler struct {
	handler *Handler
	server  *grpc.Server
	logger  *logrus.Entry
}

func NewScheduler(handler *Handler) *Scheduler {
	logger := log.WithField("name", "scheduler")
	srv := grpc.NewServer(
		grpc_middleware.WithUnaryServerChain(grpc_logrus.UnaryServerInterceptor(logger)),
		grpc_middleware.WithStreamServerChain(grpc_logrus.StreamServerInterceptor(logger)),
	)
	api.RegisterSchedulerServer(srv, handler)
	return &Scheduler{
		handler: handler,
		server:  srv,
		logger:  logger,
	}
}

func (s *Scheduler) Start(address string) error {
	s.logger.Info("start")

	lis, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}

	return s.server.Serve(lis)
}

func (s *Scheduler) Stop() {
	s.server.Stop()
}
