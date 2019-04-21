package model

import (
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type Common struct {
	CreatedAt time.Time `json:"createdAt" xorm:"'created_at'"`
	UpdatedAt time.Time `json:"updatedAt" xorm:"'updated_at'"`
}

type User struct {
	Id        int64 `json:"id" xorm:"'user_id' pk autoincr"`
	FirstName string `json:"firstName" xorm:"'first_name'"`
	LastName  string `json:"lastName" xorm:"'last_name'"`
	Email     string `json:"email" xorm:"'email' not null unique"`
	Status    int8    `json:"status" xorm:"'status' not null"`
	Common
}

const (
	UserStatusValid = 1   // ステータス：有効
	UserStatusInValid = 9 // ステータス：無効
)