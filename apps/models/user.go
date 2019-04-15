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

	dataSource := "user=sample host=go_graphql_app_db password=sample port=33006 dbname=sample sslmode=disable"
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

	return User{
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

// NewUserRepository ...
func NewUserRepository() UserRepository {
	return UserRepository{}
}

// GetByID ...
func (m UserRepository) GetByID(id int) *User {
	var user = User{Id: id}
	has, _ := engine.Get(&user)
	if has {
		return &user
	}

	return nil
}
