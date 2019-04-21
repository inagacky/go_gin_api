package model

import (
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type CommonModelFields struct {
	Id          uint64     `gorm:"primary_key" json:"id"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at"`
}

type User struct {
	CommonModelFields
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Status    uint8  `json:"status"`
}

const (
	UserStatusValid = 1   // ステータス：有効
	UserStatusInValid = 9 // ステータス：無効
)