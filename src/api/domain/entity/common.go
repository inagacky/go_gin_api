package entity

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
