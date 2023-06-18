package controller

import "sluck/model"

type UserRequest struct {
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
	Age   int    `json:"age" validate:"required,numeric"`
}

func ToModel(r *UserRequest) *model.User {
	return &model.User{
		Name:  r.Name,
		Email: r.Email,
		Age:   r.Age,
	}
}
