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
