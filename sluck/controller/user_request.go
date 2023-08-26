package controller

import (
	"sluck/model"
	"time"
)

type UseRequest struct {
	Name  string `json:"name" validate:"required"`
	Age   int    `json:"age" validate:"gte=0,lt=200"`
	Email string `json:"email" validate:"required,email"`
}

func toModel(req UseRequest) model.User {
	now := time.Now()
	return model.User{
		Name:      req.Name,
		Age:       req.Age,
		Email:     req.Email,
		CreatedAt: now,
		UpdatedAt: now,
	}
}
