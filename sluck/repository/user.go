package repository

import (
	"database/sql"
)

type UserRepository interface {
	// Create(ctx context.Context) error
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
