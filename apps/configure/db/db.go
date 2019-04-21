package db

import (
	//	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)
var db *gorm.DB


func Init() {
	var err error

	// データソースの定義
	dataSource := "root:sample@tcp(127.0.0.1:33006)/sample?parseTime=true&charset=utf8"
	db, err = gorm.Open("mysql", dataSource)
	if err != nil {
		log.Fatalf("データベースの接続に失敗しました。: %v", err)
	}
}
// DBを返却
func GetDB() *gorm.DB {
	return db
}