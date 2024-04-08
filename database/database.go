package database

import (
	"fmt"
	"meeting3/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "root@tcp(127.0.0.1:3306)/meeting4?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Data Base Gagal Terhubung")
	}
	fmt.Println("Data Base Telah Terhubung")

	DB = db

	db.AutoMigrate(models.Admin{}, models.User{})

}
