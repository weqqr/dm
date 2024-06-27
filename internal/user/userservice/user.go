package userservice

import (
	"context"

	"dm/internal/domain"
	"dm/internal/rpc"
	"dm/internal/user/usercore"

	"google.golang.org/grpc"
)

// UserService представляет сервис пользователя, реализующий методы gRPC сервера.
type UserService struct {
	// UnsafeUserServer предоставляет методы для реализации небезопасного сервера.
	rpc.UnsafeUserServer

	// core представляет ядро пользовательской логики, с которым взаимодействует сервис.
	core *usercore.Core
}

// New создает новый экземпляр UserService с предоставленным ядром пользовательской логики.
func New(core *usercore.Core) *UserService {
	return &UserService{
		core: core,
	}
}

// Register регистрирует сервис на gRPC сервере.
func (s *UserService) Register(server *grpc.Server) {
	rpc.RegisterUserServer(server, s)
}

// CreateUser создает нового пользователя, используя контекст и запрос на создание пользователя.
func (s *UserService) CreateUser(ctx context.Context, request *domain.CreateUserRequest) (*domain.CreateUserResponse, error) {
	return s.core.CreateUser(ctx, request)
}

// GetUser возвращает информацию о пользователе по предоставленному запросу.
func (s *UserService) GetUser(ctx context.Context, request *domain.GetUserRequest) (*domain.GetUserResponse, error) {
	return s.core.GetUser(ctx, request)
}

// UpdateUser обновляет информацию о пользователе на основе предоставленного запроса.
func (s *UserService) UpdateUser(ctx context.Context, request *domain.UpdateUserRequest) (*domain.UpdateUserResponse, error) {
	return s.core.UpdateUser(ctx, request)
}

// DeleteUser удаляет пользователя на основе предоставленного запроса.
func (s *UserService) DeleteUser(ctx context.Context, request *domain.DeleteUserRequest) (*domain.DeleteUserResponse, error) {
	return s.core.DeleteUser(ctx, request)
}
