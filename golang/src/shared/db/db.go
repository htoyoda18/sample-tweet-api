package db

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/htoyoda18/sample-tweet-api/golang/src/domain/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() (*gorm.DB, *sql.DB, error) {
	dbconf := fmt.Sprintf(
		"%s:%s@tcp(db:3306)/%s?charset=utf8&parseTime=true",
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_DATABASE"),
	)

	db, err := gorm.Open(mysql.Open(dbconf), &gorm.Config{})
	if err != nil {
		return nil, nil, err
	}

	migrate(db)

	sqlDB, sqlDBErr := db.DB()
	if sqlDBErr != nil {
		return nil, nil, err
	}

	return db, sqlDB, nil
}

func migrate(db *gorm.DB) {
	db.AutoMigrate(&model.User{})
}
