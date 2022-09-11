package user_test

import (
	"log"
	"os"
	"testing"

	response "github.com/htoyoda18/sample-tweet-api/golang/src/controller/handler/user/test/response"
	"github.com/htoyoda18/sample-tweet-api/golang/src/domain/model"
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

func TestAddUser(t *testing.T) {
	//ユーザが正しく作成されるのか？
	t.Run("Success", func(t *testing.T) {
		req := helper.Request(t, requestPass+"/add_users/success.json", 1)
		res := response.AddUsersResponse(req)

		assert.Equal(t, res.Name, "豊田")
		assert.Equal(t, res.Email, "h.toyoda@aaaaa.com")

		t.Cleanup(func() {
			helper.TeardownFixture(fixturePass)
		})
	})
}

func TestDeleteUser(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		req := helper.Request(t, requestPass+"/delete_user/success.json", 1)
		res := sharedRes.NilResponse(req)

		var user []model.User

		db.Find(&user)

		assert.Equal(t, len(user), 1)
		assert.Equal(t, res, nil)

		t.Cleanup(func() {
			helper.TeardownFixture(fixturePass)
		})
	})

	t.Run("FailToUserId", func(t *testing.T) {
		req := helper.Request(t, requestPass+"/delete_user/fail_to_user_id.json", 1)
		resCode := sharedRes.ErrorResponseCode(req)

		var user []model.User

		db.Find(&user)

		assert.Equal(t, len(user), 2)
		assert.Equal(t, resCode, 400)

		t.Cleanup(func() {
			helper.TeardownFixture(fixturePass)
		})
	})
}
