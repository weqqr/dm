package userstorage

import (
	"context"
	"errors"

	"dm/db"
)

// UserStorage представляет хранилище данных пользователей.
type UserStorage struct {
	database db.Database
}

// New создает новый экземпляр UserStorage.
func New(database db.Database) *UserStorage {
	return &UserStorage{
		database: database,
	}
}

type (
	// CreateUserRequest представляет данные запроса для создания нового пользователя.
	CreateUserRequest struct {
		// Name имя создаваемого пользователя
		Name string
	}

	// CreateUserResponse представляет данные ответа после создания нового пользователя.
	CreateUserResponse struct {
		// ID уникальный идентификатор созданного пользователя
		ID string
	}
)

const sqlCreateUser = `
	insert into 
		users.accounts (nickname)
	values ($1)
	returning id
`

// CreateUser создает нового пользователя в базе данных.
func (u *UserStorage) CreateUser(ctx context.Context, user CreateUserRequest) (CreateUserResponse, error) {
	if user.Name == "" {
		return CreateUserResponse{}, errors.New("name is empty")
	}

	row := u.database.QueryRow(ctx, sqlCreateUser, user.Name)

	var id string

	if err := row.Scan(&id); err != nil {
		return CreateUserResponse{}, err
	}

	return CreateUserResponse{
		ID: id,
	}, nil
}

type (
	// GetUserRequest представляет данные запроса для получения пользователя.
	GetUserRequest struct {
		// ID уникальный идентификатор пользователя для получения
		ID string
	}

	// GetUserResponse представляет данные ответа после получения пользователя.
	GetUserResponse struct {
		// ID уникальный идентификатор полученного пользователя
		ID string
		// Name никнейм полученного пользователя
		Name string
	}
)

const sqlGetUser = `
	select
		nickname
	from
		users.accounts a
	where
		a.id = $1
`

// GetUser получает пользователя из базы данных.
func (u *UserStorage) GetUser(ctx context.Context, user GetUserRequest) (GetUserResponse, error) {
	if user.ID == "" {
		return GetUserResponse{}, errors.New("id is empty")
	}

	row := u.database.QueryRow(ctx, sqlGetUser, user.ID)

	var name string

	if err := row.Scan(&name); err != nil {
		return GetUserResponse{}, err
	}

	return GetUserResponse{
		ID:   user.ID,
		Name: name,
	}, nil
}

type (
	// DeleteUserRequest представляет данные запроса для удаления пользователя.
	DeleteUserRequest struct {
		// ID уникальный идентификатор пользователя для удаления
		ID string
	}

	// DeleteUserResponse представляет данные ответа после удаления пользователя.
	DeleteUserResponse struct{}
)

const sqlDeleteUser = `
	delete from
		users.accounts a
	where 
		a.id = $1
`

// DeleteUser удаляет пользователя из базы данных.
func (u *UserStorage) DeleteUser(ctx context.Context, user DeleteUserRequest) (DeleteUserResponse, error) {
	if user.ID == "" {
		return DeleteUserResponse{}, errors.New("id is empty")
	}

	_, err := u.database.Query(ctx, sqlDeleteUser, user.ID)
	if err != nil {
		return DeleteUserResponse{}, err
	}

	return DeleteUserResponse{}, nil
}

type (
	// UpdateUserRequest представляет данные запроса на обновление пользователя.
	UpdateUserRequest struct {
		// ID уникальный идентификатор обновляемого пользователя
		ID string
		// Name новый никнейм пользователя
		Name *string
	}

	// UpdateUserResponse представляет данные ответа после обновления пользователя.
	UpdateUserResponse struct {
		// Name обновленный никнейм пользователя
		Name string
	}
)

const sqlUpdateUser = `
	update
		users.accounts
	set
		nickname = $2
	where
		id = $1
`

// UpdateUser обновляет пользователя в базе данных.
func (u *UserStorage) UpdateUser(ctx context.Context, user UpdateUserRequest) (UpdateUserResponse, error) {
	if user.ID == "" {
		return UpdateUserResponse{}, errors.New("id is empty")
	}

	if *user.Name == "" {
		return UpdateUserResponse{}, nil
	}

	_, err := u.database.Query(ctx, sqlUpdateUser, user.ID, user.Name)
	if err != nil {
		return UpdateUserResponse{}, err
	}

	return UpdateUserResponse{}, nil
}
