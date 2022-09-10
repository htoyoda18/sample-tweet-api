package login_test

import (
	"log"
	"os"
	"testing"

	sharedDB "github.com/htoyoda18/sample-tweet-api/golang/src/shared/db"
	helper "github.com/htoyoda18/sample-tweet-api/golang/src/test/helper"
	sharedRes "github.com/htoyoda18/sample-tweet-api/golang/src/test/response"
	"gopkg.in/go-playground/assert.v1"
	"gorm.io/gorm"
)

var db *gorm.DB

var fixturePass = "../../../test/fixtures/default"
var requestPass = "./test/request"

func TestMain(m *testing.M) {
	err := helper.InitFixture(fixturePass)

	if err != nil {
		log.Fatal(err)
	}

	// DBを定義
	gormDB, sqlDB, _ := sharedDB.InitDB()
	db = gormDB

	run := m.Run()
	sqlDB.Close()

	os.Exit(run)
}

func TestLogin(t *testing.T) {
	//ユーザが正しく作成されるのか？
	t.Run("Success", func(t *testing.T) {
		req := helper.Request(t, requestPass+"/login/login_success.json")
		res := sharedRes.NilResponse(req)

		assert.Equal(t, res, nil)

		t.Cleanup(func() {
			helper.TeardownFixture(fixturePass)
		})
	})
}
