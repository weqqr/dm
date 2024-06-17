package userstorage

import "dm/internal/db"

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
