package dbconn

import (
	"log"

	"github.com/jinzhu/gorm"

	_ "github.com/go-sql-driver/mysql"
)

var Db *gorm.DB

func DbConn() {
	database, err := gorm.Open("mysql", "root:<password>@tcp(127.0.0.1:3306)/gormTest")

	if err != nil {
		log.Fatal("Error connect db", err)
	}

	Db = database

}
