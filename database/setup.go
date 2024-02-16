package database

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

func SetupDB() *gorm.DB {
	USER := "root"
	PASS := ""
	HOST := "127.0.0.1"
	PORT := "3306"
	DBNAME := "project_office_monitoring_backend"
	URL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", USER, PASS, HOST, PORT, DBNAME)
	db, err := gorm.Open("mysql", URL)
	if err != nil {
		panic(err.Error())
	}
	return db

}
