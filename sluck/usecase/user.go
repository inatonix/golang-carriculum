package usecase

import (
	"context"
	"sluck/model"
	"sluck/repository"
)

type UserUsecase interface {
	GetByID(ctx context.Context, id string) (*model.User, error)
	Create(ctx context.Context, user *model.User) (string, error)
	Update(ctx context.Context, user *model.User) error
	Delete(ctx context.Context, id string) error
}

type userUsecase struct {
	r repository.UserRepository
}

func NewUserUsecase(r repository.UserRepository) UserUsecase {
	return &userUsecase{r}
}

func (c *userUsecase) Create(ctx context.Context, user *model.User) (string, error) {
	id, err := c.r.Create(ctx, user)
	if err != nil {
		return "", err
	}

	return id, nil
}

func (c *userUsecase) GetByID(ctx context.Context, id string) (*model.User, error) {
	user, err := c.r.Read(ctx, id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (c *userUsecase) Update(ctx context.Context, user *model.User) error {
	err := c.r.Update(ctx, user)
	if err != nil {
		return err
	}

	return nil
}

func (c *userUsecase) Delete(ctx context.Context, id string) error {
	err := c.r.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
