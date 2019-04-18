package repository

import (
	"github.com/go_gin_sample/apps/configure/db"
	m "github.com/go_gin_sample/apps/domain/model"
	"log"
)

type UserRepository struct {}

// IDを元に取得します
func (c *UserRepository) FindByUserId(id int) *m.User {

	var user = m.User{Id: id}
	db := db.GetDB()
	_, err := db.Get(&user)
	if err != nil {
		log.Fatalf("レコードの取得に失敗しました。: %v", err)
	}

	return &user
}
