package usecase

import (
	"context"

	"sluck/model"
	"sluck/repository"
)

type UserUsecase interface {
	Create(ctx context.Context, u *model.User) error
	// Get(ctx context.Context) error
	// Update(ctx context.Context) error
	// Delete(ctx context.Context) error
}

type userUsecase struct {
	r repository.UserRepository
}

func NewUserUsecase(r repository.UserRepository) UserUsecase {
	return &userUsecase{
		r: r,
	}
}

func (u userUsecase) Create(ctx context.Context, user *model.User) error {
	// user作成
	return nil
}
