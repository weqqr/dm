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
