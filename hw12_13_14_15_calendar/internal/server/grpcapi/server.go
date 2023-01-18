package grpcapi

import (
	"context"
	"fmt"
	"github.com/evgen1067/hw12_13_14_15_calendar/internal/repository"
	"github.com/evgen1067/hw12_13_14_15_calendar/internal/service"
	"net"
	"time"

	"github.com/evgen1067/hw12_13_14_15_calendar/api"
	"github.com/evgen1067/hw12_13_14_15_calendar/internal/config"
	"github.com/evgen1067/hw12_13_14_15_calendar/internal/logger"
	"google.golang.org/grpc"
)

type Deps struct {
	ctx     context.Context
	repo    repository.EventsRepo
	address string
}

type Server struct {
	address  string
	services *service.Services
	api.UnimplementedEventServiceServer
	Srv *grpc.Server
}

func InitGRPC(cfg *config.Config, services *service.Services) *Server {
	srv := grpc.NewServer(grpc.UnaryInterceptor(LoggerInterceptor))
	server := &Server{
		Srv:      srv,
		address:  net.JoinHostPort(cfg.GRPC.Host, cfg.GRPC.Port),
		services: services,
	}
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
	handler grpc.UnaryHandler,
) (interface{}, error) {
	// TODO логгер вообще работает?
	start := time.Now()
	// Calls the handler
	h, err := handler(ctx, req)

	latency := time.Since(start).Nanoseconds()
	message := fmt.Sprintf("%v %v ns %v", info.FullMethod, latency, err)
	logger.Logger.Info(message)
	return h, err
}
