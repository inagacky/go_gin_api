package user

import (
	"github.com/go_gin_sample/apps/configure/db"
	"github.com/go_gin_sample/apps/entity"
	"log"
)

type Repository struct {}

// IDを元に取得します
func (c *Repository) FindByUserId(id int) *entity.User {

	var user = entity.User{Id: id}
	db := db.GetDB()
	_, err := db.Get(&user)
	if err != nil {
		log.Fatalf("レコードの取得に失敗しました。: %v", err)
	}

	return &user
}
