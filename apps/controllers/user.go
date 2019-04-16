package controllers

import (
	"github.com/go_gin_sample/apps/models"
)


func NewUser() UserController {
	return UserController{}
}

type UserController struct {
}

func (c UserController) Get(id int) interface{} {
	repo := models.NewUserRepository()
	user := repo.GetById(id)

	return user
}