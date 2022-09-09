package test

import (
	"log"

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

	user := "webuser"
	password := "webpass"
	host := "localhost"
	port := "3306"
	database_name := "go_mysql8_test"

	dbconf := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + database_name + "?charset=utf8mb4"

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
