package usercore

import (
	"context"

	"dm/internal/domain"
	"dm/internal/user/userstorage"
)

// Core представляет основную логику приложения, используя хранилище пользователей.
type Core struct {
	storage *userstorage.UserStorage
}

// New создает новый экземпляр Core.
func New(storage *userstorage.UserStorage) *Core {
	return &Core{
		storage: storage,
	}
}

func (c *Core) CreateUser(ctx context.Context, request *domain.CreateUserRequest) (*domain.CreateUserResponse, error) {
	response, err := c.storage.CreateUser(ctx, userstorage.CreateUserRequest{
		Name: request.GetName(),
	})

	if err != nil {
		return nil, err
	}

	return &domain.CreateUserResponse{
		Id: response.ID,
	}, nil
}

func (c *Core) GetUser(ctx context.Context, request *domain.GetUserRequest) (*domain.GetUserResponse, error) {
	response, err := c.storage.GetUser(ctx, userstorage.GetUserRequest{
		ID: request.GetId(),
	})

	if err != nil {
		return nil, err
	}

	return &domain.GetUserResponse{
		User: &domain.User{
			Id:   request.GetId(),
			Name: response.Name,
		},
	}, nil
}

func (c *Core) UpdateUser(ctx context.Context, request *domain.UpdateUserRequest) (*domain.UpdateUserResponse, error) {
	name := request.GetName()

	_, err := c.storage.UpdateUser(ctx, userstorage.UpdateUserRequest{
		ID:   request.GetId(),
		Name: &name,
	})

	if err != nil {
		return nil, err
	}

	return &domain.UpdateUserResponse{}, nil
}

func (c *Core) DeleteUser(ctx context.Context, request *domain.DeleteUserRequest) (*domain.DeleteUserResponse, error) {
	_, err := c.storage.DeleteUser(ctx, userstorage.DeleteUserRequest{
		ID: request.GetId(),
	})

	if err != nil {
		return nil, err
	}

	return &domain.DeleteUserResponse{}, nil
}
