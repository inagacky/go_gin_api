package repository

import (
	"github.com/go_gin_sample/apps/configure/db"
	l "github.com/go_gin_sample/apps/configure/logger"
	m "github.com/go_gin_sample/apps/domain/model"
)
var logger  = l.GetLogger()

type UserRepository struct {}

// IDを元にユーザーを取得します
func (c *UserRepository) FindByUserId(id uint64) (*m.User, error) {

	var user = m.User{}
	user.Id = id
	db := db.GetDB()
	if err := db.First(&user).Error; err != nil {
		logger.Error("ユーザーの取得処理でエラーが発生しました: %v", err)
		return nil, err
	}

	return &user, nil
}

// ユーザー情報を作成します
func (c *UserRepository) Save(user *m.User) (*m.User, error) {

/*
	c.setInitTime(&user.Common)
	db := db.GetDB()
	_, err := db.Insert(user)
	if err != nil {
		logger.Error("ユーザーの作成処理でエラーが発生しました: %v", err)
		return nil, err
	}
*/
	return nil, nil
}