package service

import (
	l "github.com/go_gin_sample/apps/configure/logger"
	r "github.com/go_gin_sample/apps/infrastructure/repository"
	us "github.com/go_gin_sample/apps/usecase/user"
)
var logger  = l.GetLogger()

type UserService struct {}

// IDを元にレコードを取得します
func (c *UserService) GetById(id int) *us.GetUserResponse {

	logger.Info("GetById")
	var repo = r.UserRepository{}
	user := repo.FindByUserId(id)

	return &us.GetUserResponse{User:user}
}