package server

import (
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"github.com/mitene/terrafire/internal/api"
	"github.com/mitene/terrafire/internal/utils"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
	"net/http"
)

type Server struct {
	web       *Web
	scheduler *Scheduler
}

func New(projects map[string]*api.Project, db *DB, git utils.Git) *Server {
	actions := make(chan *api.GetActionResponse, 100)
	actionControls := make(chan *api.GetActionControlResponse, 100)

	web := &Web{
		projects:   projects,
		workspaces: map[string]map[string]*api.Workspace{},
		actions:    actions,
		db:         db,
		git:        git,
	}

	scheduler := &Scheduler{
		actions:        actions,
		actionControls: actionControls,
		db:             db,
	}

	return &Server{
		web:       web,
		scheduler: scheduler,
	}
}

func (s *Server) StartWeb(address string) error {
	logger := log.WithField("name", "web")
	server := grpc.NewServer(
		grpc_middleware.WithUnaryServerChain(grpc_logrus.UnaryServerInterceptor(logger)),
		grpc_middleware.WithStreamServerChain(grpc_logrus.StreamServerInterceptor(logger)),
	)
	api.RegisterWebServer(server, s.web)

	srv := grpcweb.WrapServer(server)
	assets := http.FileServer(newAssetsFileSystem())
	handler := http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
		if srv.IsGrpcWebRequest(req) {
			srv.ServeHTTP(resp, req)
		} else {
			assets.ServeHTTP(resp, req)
		}
	})

	utils.LogError(s.web.refreshAllProject())

	logger.Info("start")

	return http.ListenAndServe(address, handler)
}

func (s *Server) StartScheduler(address string) error {
	logger := log.WithField("name", "scheduler")
	server := grpc.NewServer(
		grpc_middleware.WithUnaryServerChain(grpc_logrus.UnaryServerInterceptor(logger)),
		grpc_middleware.WithStreamServerChain(grpc_logrus.StreamServerInterceptor(logger)),
	)
	api.RegisterSchedulerServer(server, s.scheduler)

	logger.Info("start")

	lis, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}

	return server.Serve(lis)
}
