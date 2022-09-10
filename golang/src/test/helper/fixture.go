package test

import (
	"log"

	"github.com/go-testfixtures/testfixtures/v3"
	sharedDB "github.com/htoyoda18/sample-tweet-api/golang/src/shared/db"
)

func InitFixture(pass string) error {
	err := createFixture(pass)

	return err
}

func createFixture(pass string) error {
	var (
		fixtures *testfixtures.Loader
	)

	db, sqlDB, _ := sharedDB.InitDB()

	sqlDB, sqlDBErr := db.DB()
	if sqlDBErr != nil {
		return sqlDBErr
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

	err := fixtures.Load()

	return err
}
