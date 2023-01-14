package grpcApi

import (
	"context"
	"fmt"
	"github.com/evgen1067/hw12_13_14_15_calendar/api"
	"github.com/evgen1067/hw12_13_14_15_calendar/internal/config"
	"github.com/evgen1067/hw12_13_14_15_calendar/internal/logger"
	"github.com/evgen1067/hw12_13_14_15_calendar/internal/repository"
	"google.golang.org/grpc"
	"net"
	"time"
)

type Deps struct {
	ctx     context.Context
	repo    repository.EventsRepo
	address string
}

type Server struct {
	Deps
	api.UnimplementedEventServiceServer
	Srv *grpc.Server
}

func InitGRPC(_ctx context.Context, _repo repository.EventsRepo, cfg *config.Config) *Server {
	srv := grpc.NewServer(grpc.UnaryInterceptor(LoggerInterceptor))
	server := &Server{
		Deps: Deps{
			ctx:     _ctx,
			repo:    _repo,
			address: net.JoinHostPort(cfg.GRPC.Host, cfg.GRPC.Port),
		},
		Srv: nil,
	}
	server.Srv = srv
	api.RegisterEventServiceServer(srv, server)
	return server
}

func (s *Server) ListenAndServe() error {
	lis, err := net.Listen("tcp", s.address)
	if err != nil {
		return err
	}
	if err := s.Srv.Serve(lis); err != nil {
		return err
	}
	return nil
}

func (s *Server) Stop() {
	s.Srv.GracefulStop()
}

func LoggerInterceptor(ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (interface{}, error) {
	start := time.Now()
	// Calls the handler
	h, err := handler(ctx, req)
	latency := time.Since(start).Nanoseconds()
	message := fmt.Sprintf("%v %v ns", info.FullMethod, latency)
	logger.Logger.Info(message)
	return h, err
}
