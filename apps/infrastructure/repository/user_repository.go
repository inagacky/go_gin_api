package repository

import (
	"github.com/go_gin_sample/apps/configure/db"
	l "github.com/go_gin_sample/apps/configure/logger"
	m "github.com/go_gin_sample/apps/domain/model"
	"time"
)
var logger  = l.GetLogger()

type UserRepository struct {}

// IDを元にユーザーを取得します
func (c *UserRepository) FindByUserId(id int64) (*m.User, error) {

	var user = m.User{Id: id}
	db := db.GetDB()
	_, err := db.Get(&user)
	if err != nil {
		logger.Error("ユーザーの取得処理でエラーが発生しました: %v", err)
		return nil, err
	}

	return &user, nil
}

// ユーザー情報を作成します
func (c *UserRepository) Save(user *m.User) (*m.User, error) {

	c.setInitTime(&user.Common)
	db := db.GetDB()
	_, err := db.Insert(user)
	if err != nil {
		logger.Error("ユーザーの作成処理でエラーが発生しました: %v", err)
		return nil, err
	}


	return user, nil
}


func (c *UserRepository) setInitTime(model *m.Common) {
	model.CreatedAt = time.Now()
	model.UpdatedAt = time.Now()

}