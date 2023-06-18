package repository

import (
	"context"
	"database/sql"
	"sluck/model"
	"strconv"
)

type UserRepository interface {
	Create(ctx context.Context, user *model.User) (string, error)
	// Get(ctx context.Context) error
	// Update(ctx context.Context) error
	// Delete(ctx context.Context) error
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) Create(ctx context.Context, user *model.User) (string, error) {
	// save user to db
	result, err := r.db.Exec("INSERT INTO users (name, email, age) VALUES (?, ?, ?)", user.Name, user.Email, user.Age)
	if err != nil {
		return "", err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return "", nil
	}

	s := strconv.FormatInt(id, 10)
	return s, nil
}
