package controller

import (
	"sluck/model"
	"time"
)

type PostRequest struct {
	Name  string `json:"name" validate:"required"`
	Age   int    `json:"age" validate:"gte=0,lt=200"`
	Email string `json:"email" validate:"required,email"`
}

type PutRquest struct {
	Name  string `json:"name"`
	Age   int    `json:"age" validate:"gte=0,lt=200"`
	Email string `json:"email" validate:"email"`
}

func post2Model(req PostRequest) *model.User {
	now := time.Now()
	return &model.User{
		Name:      req.Name,
		Age:       req.Age,
		Email:     req.Email,
		CreatedAt: now,
		UpdatedAt: now,
	}
}

func put2Model(req PutRquest) *model.User {
	now := time.Now()
	return &model.User{
		Name:      req.Name,
		Age:       req.Age,
		Email:     req.Email,
		UpdatedAt: now,
	}
}
