package db

import (
	"database/sql"
	"fmt"

	"github.com/htoyoda18/sample-tweet-api/shared/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB(conf config.Configuration) (*gorm.DB, *sql.DB, error) {
	dbConnection := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?parseTime=true&loc=Local", conf.DBUserTest, conf.DBPassTest, conf.DBHost, conf.DBNameTest)
	db, err := gorm.Open(mysql.Open(dbConnection), &gorm.Config{})
	if err != nil {
		return nil, nil, err
	}

	sqlDB, sqlDBErr := db.DB()
	if sqlDBErr != nil {
		return nil, nil, err
	}

	return db, sqlDB, nil
}
