package models

import (
	"GoModDemo/config"
	"fmt"
	"log"

	//
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

//DB db
var DB *gorm.DB

func init() {
	DB = initDB()
}

func initDB() *gorm.DB {
	username := config.Conf.GetDefaultValue("mysql", "username", "username")
	password := config.Conf.GetDefaultValue("mysql", "password", "password")
	host := config.Conf.GetDefaultValue("mysql", "host", "127.0.0.1")
	port := config.Conf.GetDefaultValue("mysql", "port", "3306")
	dbName := config.Conf.GetDefaultValue("mysql", "db_name", "test")
	connStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, port, dbName)
	var err error
	db, err := gorm.Open("mysql", connStr)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return "t_" + defaultTableName
	}
	return db
}
