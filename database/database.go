// database/database.go
package database

import (
	"douyin/pkg/config"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func init() {
	ConnectDB()
}

func ConnectDB() error {
	dbURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		"root",
		"root",
		"127.0.0.1",
		config.AppConfigInstance.DBPort,
		// "3306",
		"douyin",
		// config.AppConfigInstance.DBUser,
		// config.AppConfigInstance.DBPassword,
		// config.AppConfigInstance.DBHost,
		// config.AppConfigInstance.DBPort,
		// config.AppConfigInstance.DatabaseName,
	)

	var err error
	// DB, err = gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/douyin?charset=utf8&parseTime=true")

	DB, err = gorm.Open("mysql", dbURL)
	if err != nil {
		panic("Failed to connect to the database")
	} else {
		fmt.Println("Connected to the database")
		
	}

	return nil
}