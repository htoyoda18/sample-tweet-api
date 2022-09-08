package test

import (
	"database/sql"
	"log"

	"github.com/go-testfixtures/testfixtures/v3"
	"github.com/htoyoda18/sample-tweet-api/shared/config"
)

func InitFixture(pass string) error {
	err := createFixture(pass)

	return err
}

func createFixture(pass string) error {
	var con config.Configuration
	con, _ = config.InitConfiguration()

	var (
		db       *sql.DB
		fixtures *testfixtures.Loader
	)

	// dbと接続
	db, errDB := sql.Open("mysql", con.DBUserTest+":"+con.DBPassTest+"@tcp("+con.DBHost+":3306)/"+con.DBNameTest+"?parseTime=true")
	if errDB != nil {
		log.Fatal(errDB)
	}

	defer db.Close()

	fixtures, fixturesErr := testfixtures.New(
		testfixtures.Database(db),
		testfixtures.Dialect("mysql"),
		testfixtures.Directory(pass),
	)
	if fixturesErr != nil {
		log.Fatal(fixturesErr)
	}

	err := fixtures.Load()

	return err
}
