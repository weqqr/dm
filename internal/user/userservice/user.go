package userservice

import (
	"context"
	"dm/internal/domain"
	"dm/internal/rpc"
	"dm/internal/user/usercore"

	"google.golang.org/grpc"
)

type UserService struct {
	rpc.UnsafeUserServer

	core *usercore.Core
}

func New(core *usercore.Core) *UserService {
	return &UserService{
		core: core,
	}
}

func (s *UserService) Register(server *grpc.Server) {
	rpc.RegisterUserServer(server, s)
}

func (s *UserService) CreateUser(ctx context.Context, request *domain.CreateUserRequest) (*domain.CreateUserResponse, error) {
	return s.core.CreateUser(ctx, request)
}

func (s *UserService) GetUser(ctx context.Context, request *domain.GetUserRequest) (*domain.GetUserResponse, error) {
	return s.core.GetUser(ctx, request)
}

func (s *UserService) UpdateUser(ctx context.Context, request *domain.UpdateUserRequest) (*domain.UpdateUserResponse, error) {
	return s.core.UpdateUser(ctx, request)
}

func (s *UserService) DeleteUser(ctx context.Context, request *domain.DeleteUserRequest) (*domain.DeleteUserResponse, error) {
	return s.core.DeleteUser(ctx, request)
}
