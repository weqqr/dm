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
	response, err := c.storage.CreateUser(ctx, userstorage.CreateUserRequest{
		Name: request.Name,
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
		ID: request.Id,
	})

	if err != nil {
		return nil, err
	}

	return &domain.GetUserResponse{
		User: &domain.User{
			Id:   request.Id,
			Name: response.Name,
		},
	}, nil
}

func (c *Core) UpdateUser(ctx context.Context, request *domain.UpdateUserRequest) (*domain.UpdateUserResponse, error) {
	_, err := c.storage.UpdateUser(ctx, userstorage.UpdateUserRequest{
		ID:   request.Id,
		Name: request.Name,
	})

	if err != nil {
		return nil, err
	}

	return &domain.UpdateUserResponse{}, nil
}

func (c *Core) DeleteUser(ctx context.Context, request *domain.DeleteUserRequest) (*domain.DeleteUserResponse, error) {
	_, err := c.storage.DeleteUser(ctx, userstorage.DeleteUserRequest{
		ID: request.Id,
	})

	if err != nil {
		return nil, err
	}

	return &domain.DeleteUserResponse{}, nil
}
