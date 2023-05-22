package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Dbinstance struct {
	Db *gorm.DB
}

var DB Dbinstance

func ConnectDB() {

	dsn := "root@tcp(127.0.0.1:3306)/paymentGateway?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("Failed to connect to Database. \n", err)
		os.Exit(2)
	} else {
		fmt.Println("Mysql Connected Successfully!")
	}
	// defer db.Close()

	DB = Dbinstance{
		Db: db,
	}
}
