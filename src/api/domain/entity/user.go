package entity

import (
	_ "github.com/go-sql-driver/mysql"
)

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


// 値のコピーを行う
func (u *User) ValueCopy(org *User) *User {

	u.LastName = org.LastName
	u.FirstName = org.FirstName
	u.Email = org.Email

	return u
}
