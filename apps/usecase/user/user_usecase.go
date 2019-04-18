package usecase

import "github.com/go_gin_sample/apps/domain/model"

type GetUserRequest struct {
	Id string `uri:"id" binding:"number,required,min=1"`
}

type GetUserResponse struct {
	User *model.User `json:"user"`
}


type CreateUserRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Status    int    `json:"status"`
}

type CreateUserResponse struct {
	User *model.User `json:"user"`
}