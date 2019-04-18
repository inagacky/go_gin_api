package user

import (
	l "github.com/go_gin_sample/apps/configure/logger"
	"github.com/go_gin_sample/apps/entity"
	"github.com/go_gin_sample/apps/repository"
)
var logger  = l.GetLogger()

type Service struct {}

// IDを元に取得します
func (c *Service) GetById(id int) *entity.User {

	logger.Info("GetById")
	var repo = user.Repository{}
	user := repo.FindByUserId(id)

	return user
}