package server

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
)

type Server struct {
	config     Config
	grpcServer *grpc.Server
}

type Config struct {
	Address string `toml:"address"`
}

func NewServer(config Config) *Server {
	grpcServer := grpc.NewServer()

	return &Server{
		config:     config,
		grpcServer: grpcServer,
	}
}

type Service interface {
	Register(server *grpc.Server)
}

func (s *Server) AddService(service Service) {
	service.Register(s.grpcServer)
}

func (s *Server) Run(ctx context.Context) error {
	listenConfig := net.ListenConfig{}
	listener, err := listenConfig.Listen(ctx, "tcp", s.config.Address)
	if err != nil {
		return err
	}

	log.Printf("Listening on %v", listener.Addr())

	return s.grpcServer.Serve(listener)
}
