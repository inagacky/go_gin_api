package service

import (
	l "github.com/go_gin_sample/apps/configure/logger"
	m "github.com/go_gin_sample/apps/domain/model"
	r "github.com/go_gin_sample/apps/infrastructure/repository"
)
var logger  = l.GetLogger()

type UserService struct {}

// IDを元にレコードを取得します
func (c *UserService) GetById(id int) *m.User {

	logger.Info("GetById")
	var repo = r.UserRepository{}
	user := repo.FindByUserId(id)

	return user
}