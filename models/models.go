package models

import (
	"log"

	"github.com/jinzhu/gorm"

	//
	_ "github.com/go-sql-driver/mysql"
)

//DbVersion DbVersion
var DbVersion = "V1.0.0"

//GetDb GetDb
func GetDb() *gorm.DB {
	db, err := gorm.Open("mysql", "vurtne:123456@tcp(127.0.0.1:3306)/bee_test?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		log.Fatal(err)
	}

	db.SingularTable(true)

	return db
}
