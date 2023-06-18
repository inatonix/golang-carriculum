package usecase

import (
	"context"

	"sluck/model"
	"sluck/repository"
)

type UserUsecase interface {
	Create(ctx context.Context, u *model.User) (string, error)
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

func (u userUsecase) Create(ctx context.Context, user *model.User) (string, error) {
	// user作成
	id, err := u.r.Create(ctx, user)
	if err != nil {
		return "", err
	}

	return id, nil
}
