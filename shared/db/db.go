package db

import (
	"database/sql"
	"fmt"

	"github.com/htoyoda18/sample-tweet-api/shared/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB(conf config.Configuration) (*gorm.DB, *sql.DB, error) {
	user := "webuser"
	password := "webpass"
	host := "localhost"
	port := "3306"
	database_name := "go_mysql8_development"

	dbconf := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + database_name + "?charset=utf8mb4"

	// dbConnection := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?parseTime=true&loc=Local", conf.DBUserTest, conf.DBPassTest, conf.DBHost, conf.DBNameTest)
	fmt.Println(dbconf)
	db, err := gorm.Open(mysql.Open(dbconf), &gorm.Config{})
	if err != nil {
		return nil, nil, err
	}

	sqlDB, sqlDBErr := db.DB()
	if sqlDBErr != nil {
		return nil, nil, err
	}

	return db, sqlDB, nil
}
