package db

import (
	"github.com/inagacky/go_gin_api/src/api/util"
	//	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)
var db *gorm.DB


func Init() {
	var err error

	// 環境変数から取得
	dbUser := util.Getenv("go_gin_api_DB_USER", "root")
	dbPass := util.Getenv("go_gin_api_DB_PASS", "sample")
	dbName := util.Getenv("go_gin_api_DB_NAME", "sample")
	dbHostName := util.Getenv("go_gin_api_DB_HOSTNAME", "127.0.0.1")
	dbPort := util.Getenv("go_gin_api_DB_PORT", "3306")
	protocol := "tcp("+dbHostName+":"+dbPort+")"

	// データソースの定義
	dataSource := dbUser+":"+dbPass+"@"+protocol+"/"+dbName + "?parseTime=true&charset=utf8"
	db, err = gorm.Open("mysql", dataSource)
	if err != nil {
		log.Fatalf("データベースの接続に失敗しました。: %v", err)
	}
}
// DBを返却
func GetDB() *gorm.DB {
	return db
}