package database

import (
	"assisment2/models"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

const (
	host     = "localhost"
	port     = "33061"
	dbname   = "assisment2"
	username = "root"
	password = "root"
)

func InitDb() {
	mysqlInfo := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&collation=utf8mb4_unicode_ci&parseTime=true",
		username,
		password,
		host,
		port,
		dbname)
	db, err := gorm.Open(mysql.Open(mysqlInfo), &gorm.Config{PrepareStmt: true})
	if err != nil {
		panic("failed to connect database")
	}
	DB = db
	DB.Debug().AutoMigrate(&models.Order{}, &models.Item{})
}
