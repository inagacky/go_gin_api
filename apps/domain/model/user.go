package model

import (
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type User struct {
	Id        int `json:"id" xorm:"'user_id'"`
	FirstName string `json:"firstName" xorm:"'first_name'"`
	LastName  string `json:"lastName" xorm:"'last_name'"`
	Email     string `json:"email" xorm:"'email'"`
	Status    int    `json:"status" xorm:"'status'"`
	CreatedAt time.Time `json:"createdAt" xorm:"'created_at'"`
	UpdatedAt time.Time `json:"updatedAt" xorm:"'updated_at'"`
}