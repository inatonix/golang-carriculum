package repository

import (
	"context"
	"database/sql"
	"sluck/model"
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
	return "", nil
}
