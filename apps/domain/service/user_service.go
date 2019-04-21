package service

import (
	l "github.com/go_gin_sample/apps/configure/logger"
	"github.com/go_gin_sample/apps/domain/model"
	r "github.com/go_gin_sample/apps/infrastructure/repository"
)
var logger  = l.GetLogger()

type UserService struct {}

// IDを元にレコードを取得します
func (c *UserService) GetById(id uint64) (*model.User, error) {

	repo := &r.UserRepository{}
	user, err := repo.FindByUserId(id)

	return user, err
}


// Userの作成を行います
func (c *UserService) CreateUser(user *model.User) (*model.User, error) {

	user.Status = model.UserStatusValid
	repo := &r.UserRepository{}
	user, err := repo.Save(user)

	return user, err
}