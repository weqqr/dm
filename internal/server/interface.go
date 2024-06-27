package server

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
)

// Server представляет сервер gRPC, который конфигурируется и запускается.
type Server struct {
	// config содержит конфигурацию сервера.
	config Config
	// grpcServer это инстанс gRPC сервера.
	grpcServer *grpc.Server
}

// Config содержит параметры конфигурации для сервера.
type Config struct {
	// Address это адрес, на котором сервер будет прослушивать соединения.
	Address string `toml:"address"`
}

// NewServer создает новый экземпляр сервера с заданной конфигурацией.
func NewServer(config Config) *Server {
	grpcServer := grpc.NewServer()

	return &Server{
		config:     config,
		grpcServer: grpcServer,
	}
}

// Service представляет сервис, который может быть зарегистрирован на gRPC сервере.
type Service interface {
	Register(server *grpc.Server)
}

// AddService добавляет и регистрирует сервис на gRPC сервере.
func (s *Server) AddService(service Service) {
	service.Register(s.grpcServer)
}

// Run запускает сервер, чтобы он начал прослушивание на указанном адресе.
func (s *Server) Run(ctx context.Context) error {
	listenConfig := net.ListenConfig{}

	listener, err := listenConfig.Listen(ctx, "tcp", s.config.Address)
	if err != nil {
		return err
	}

	log.Printf("Listening on %v", listener.Addr())

	return s.grpcServer.Serve(listener)
}
