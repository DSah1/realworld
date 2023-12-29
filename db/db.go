package db

import (
	"awesomeProject/internal/model"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

const (
	TEST_DB_URL     = "./../database/realworld_test.db"
	SQLITE_DB_URL   = "./database/realworld.db"
	POSTGRES_DB_URL = "host=localhost user=postgres password=dor4420 dbname=gorm port=5432 sslmode=disable TimeZone=Asia/Jerusalem"
	IS_POSTGRES     = false
)

func getDialector() gorm.Dialector {
	if IS_POSTGRES {
		return postgres.Open(POSTGRES_DB_URL)
	}
	return sqlite.Open(SQLITE_DB_URL)
}

func New() *gorm.DB {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Millisecond * 10, // Slow SQL threshold
			LogLevel:                  logger.Warn,           // Log level
			IgnoreRecordNotFoundError: false,                 // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,                  // Disable color
		},
	)
	db, err := gorm.Open(getDialector(), &gorm.Config{
		Logger: newLogger,
	})

	if err != nil {
		fmt.Println("storage error: ", err)
	}

	sqlDB, err := db.DB()

	sqlDB.SetMaxIdleConns(3)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	return db
}

func TestDB() *gorm.DB {
	dsn := TEST_DB_URL

	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("storage error: ", err)
	}

	return db
}

func dropTestDB() error {
	if err := os.Remove(TEST_DB_URL); err != nil {
		return err
	}
	return nil
}

func AutoMigrate(db *gorm.DB) {
	err := db.AutoMigrate(
		&model.User{},
		&model.Follow{},
		&model.Article{},
		&model.Tag{},
		//&model.Comment{},
	)
	if err != nil {
		return
	}
}
