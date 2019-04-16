package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"log"
	"time"
)
var engine *xorm.Engine

func init() {
	var err error

	dataSource := "root:sample@tcp(127.0.0.1:33006)/sample?parseTime=true&charset=utf8"
	engine, err = xorm.NewEngine("mysql", dataSource)
	if err != nil {
		log.Fatalf("データベースの接続に失敗しました。: %v", err)
	}
}

type User struct {
	Id        int `json:"id" xorm:"'user_id'"`
	FirstName string `json:"firstName" xorm:"'first_name'"`
	LastName  string `json:"lastName" xorm:"'last_name'"`
	Email     string `json:"email" xorm:"'email'"`
	Status    int    `json:"status" xorm:"'status'"`
	CreatedAt time.Time `json:"createdAt" xorm:"'created_at'"`
	UpdatedAt time.Time `json:"updatedAt" xorm:"'updated_at'"`
}

func NewUser(id int, firstName string, lastName string,
	email string, status int, createdAt time.Time, updatedAt time.Time) User {

	return User {
		Id:       id,
		FirstName: firstName,
		LastName: lastName,
		Email: email,
		Status: status,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}

type UserRepository struct {

}

// UserRepositoryの定義
func NewUserRepository() UserRepository {
	return UserRepository{}
}

// IDを元に取得します
func (c *UserRepository) GetById(id int) *User {
	var user = User{Id: id}
	_, err := engine.Get(&user)
	if err != nil {
		log.Fatalf("レコードの取得に失敗しました。: %v", err)
	}

	return &user
}
