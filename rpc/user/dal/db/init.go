package db

import (
	"time"

	"fmt"
	"douyin/pkg/config"
	"gorm.io/driver/mysql"
    "gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/plugin/opentelemetry/logging/logrus"
	"gorm.io/plugin/opentelemetry/tracing"
)

var DB *gorm.DB

// Init init DB
func init() {
	var err error
	gormlogrus := logger.New(
		logrus.NewWriter(),
		logger.Config{
			SlowThreshold: time.Millisecond,
			Colorful:      false,
			LogLevel:      logger.Info,
		},
	)

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

	DB, err = gorm.Open(mysql.Open(dbURL),
		&gorm.Config{
			PrepareStmt: true,
			Logger:      gormlogrus,
		},
	)
	if err != nil {
		panic(err)
	}else {
		fmt.Println("Connected to the database.")
	}

	if err := DB.Use(tracing.NewPlugin()); err != nil {
		panic(err)
	}
}