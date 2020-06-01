package server

import (
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"github.com/mitene/terrafire/internal/api"
	"github.com/mitene/terrafire/internal/utils"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net/http"
)

type Server struct {
	handler *Handler
	server  *grpc.Server
	logger  *logrus.Entry
}

func NewServer(handler *Handler) *Server {
	logger := log.WithField("name", "server")
	srv := grpc.NewServer(
		grpc_middleware.WithUnaryServerChain(grpc_logrus.UnaryServerInterceptor(logger)),
		grpc_middleware.WithStreamServerChain(grpc_logrus.StreamServerInterceptor(logger)),
	)
	api.RegisterWebServer(srv, handler)
	return &Server{
		handler: handler,
		server:  srv,
		logger:  logger,
	}
}

func (s *Server) Start(address string) error {
	srv := grpcweb.WrapServer(s.server)
	assets := http.FileServer(newAssetsFileSystem())

	s.logger.Info("start")

	go func() {
		utils.LogError(s.handler.refreshAllProject())
	}()

	return http.ListenAndServe(address, http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
		if srv.IsGrpcWebRequest(req) {
			srv.ServeHTTP(resp, req)
		} else {
			assets.ServeHTTP(resp, req)
		}
	}))
}

func (s *Server) Stop() {
	s.server.Stop()
}
