package usercore

import (
	"context"
	"dm/internal/domain"
	"dm/internal/user/userstorage"
)

type Core struct {
	storage *userstorage.UserStorage
}

func New(storage *userstorage.UserStorage) *Core {
	return &Core{
		storage: storage,
	}
}

func (c *Core) CreateUser(ctx context.Context, request *domain.CreateUserRequest) (*domain.CreateUserResponse, error) {
	return &domain.CreateUserResponse{}, nil
}

func (c *Core) GetUser(ctx context.Context, request *domain.GetUserRequest) (*domain.GetUserResponse, error) {
	return &domain.GetUserResponse{}, nil
}

func (c *Core) UpdateUser(ctx context.Context, request *domain.UpdateUserRequest) (*domain.UpdateUserResponse, error) {
	return &domain.UpdateUserResponse{}, nil
}

func (c *Core) DeleteUser(ctx context.Context, request *domain.DeleteUserRequest) (*domain.DeleteUserResponse, error) {
	return &domain.DeleteUserResponse{}, nil
}
