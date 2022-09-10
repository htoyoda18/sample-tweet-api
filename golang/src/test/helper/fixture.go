package test

import (
	"fmt"
	"log"
	"os"

	"github.com/go-testfixtures/testfixtures/v3"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitFixture(pass string) error {
	err := createFixture(pass)

	return err
}

func createFixture(pass string) error {
	var (
		fixtures *testfixtures.Loader
	)

	dbconf := fmt.Sprintf(
		"%s:%s@tcp(db:3306)/%s?charset=utf8&parseTime=true",
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_DATABASE"),
	)

	db, err := gorm.Open(mysql.Open(dbconf), &gorm.Config{})
	if err != nil {
		return err
	}

	sqlDB, sqlDBErr := db.DB()
	if sqlDBErr != nil {
		return err
	}

	defer sqlDB.Close()

	fixtures, fixturesErr := testfixtures.New(
		testfixtures.Database(sqlDB),
		testfixtures.Dialect("mysql"),
		testfixtures.Directory(pass),
	)

	if fixturesErr != nil {
		log.Fatal(fixturesErr)
	}

	err = fixtures.Load()

	return err
}
