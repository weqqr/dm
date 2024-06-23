package userstorage

import (
	"context"
	"dm/db"
	"errors"
)

type UserStorage struct {
	database db.Database
}

func New(database db.Database) *UserStorage {
	return &UserStorage{
		database: database,
	}
}

type (
	CreateUserRequest struct {
		Name string
	}

	CreateUserResponse struct {
		ID string
	}
)

func (u *UserStorage) CreateUser(ctx context.Context, user CreateUserRequest) (CreateUserResponse, error) {
	const query = `
		INSERT INTO 
			users.accounts (nickname)
		VALUES ($1)
		RETURNING id`

	if user.Name == "" {
		return CreateUserResponse{}, errors.New("name is empty")
	}

	row := u.database.QueryRow(ctx, query, user.Name)

	var id string
	err := row.Scan(&id)
	if err != nil {
		return CreateUserResponse{}, nil
	}

	return CreateUserResponse{
		ID: id,
	}, nil
}

type (
	GetUserRequest struct {
		ID string
	}

	GetUserResponse struct {
		ID   string
		Name string
	}
)

func (u *UserStorage) GetUser(ctx context.Context, user GetUserRequest) (GetUserResponse, error) {
	const query = `
		SELECT
			nickname
		FROM
			users.accounts a
		WHERE
			a.id = $1`

	if user.ID == "" {
		return GetUserResponse{}, errors.New("id is empty")
	}

	row := u.database.QueryRow(ctx, query, user.ID)

	var name string
	err := row.Scan(&name)

	if err != nil {
		return GetUserResponse{}, err
	}

	return GetUserResponse{
		ID:   user.ID,
		Name: name,
	}, nil
}

type (
	DeleteUserRequest struct {
		ID string
	}

	DeleteUserResponse struct{}
)

func (u *UserStorage) DeleteUser(ctx context.Context, user DeleteUserRequest) (DeleteUserResponse, error) {
	const query = `
		DELETE FROM
			users.accounts a
		WHERE 
			a.id = $1
		`

	if user.ID == "" {
		return DeleteUserResponse{}, errors.New("id is empty")
	}

	_, err := u.database.Query(ctx, query, user.ID)
	if err != nil {
		return DeleteUserResponse{}, err
	}

	return DeleteUserResponse{}, nil
}

type (
	UpdateUserRequest struct {
		ID   string
		Name *string
	}

	UpdateUserResponse struct {
		Name string
	}
)

func (u *UserStorage) UpdateUser(ctx context.Context, user UpdateUserRequest) (UpdateUserResponse, error) {
	const query = `
		UPDATE
			users.accounts
		SET
			nickname = $2
		WHERE
			id = $1
	`
	if user.ID == "" {
		return UpdateUserResponse{}, errors.New("id is empty")
	}

	if *user.Name == "" {
		return UpdateUserResponse{}, nil
	}

	_, err := u.database.Query(ctx, query, user.ID, user.Name)
	if err != nil {
		return UpdateUserResponse{}, err
	}

	return UpdateUserResponse{}, nil
}
