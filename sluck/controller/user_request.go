package controller

import "sluck/model"

type UserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

func ToModel(r *UserRequest) model.User {
	return model.User{
		Name:  r.Name,
		Email: r.Email,
		Age:   r.Age,
	}
}
