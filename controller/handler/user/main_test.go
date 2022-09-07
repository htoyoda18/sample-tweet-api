package user_test

import (
	"os"
	"testing"

	response "github.com/htoyoda18/sample-tweet-api/controller/handler/user/test/response"
	"github.com/htoyoda18/sample-tweet-api/shared/config"
	sharedDB "github.com/htoyoda18/sample-tweet-api/shared/db"
	helper "github.com/htoyoda18/sample-tweet-api/test/helper"
	"gopkg.in/go-playground/assert.v1"

	"gorm.io/gorm"
)

var db *gorm.DB

var requestPass = "./test/request"

func TestMain(m *testing.M) {

	con, _ := config.InitConfiguration()

	// DBを定義
	gormDB, sqlDB, _ := sharedDB.InitDB(con)
	db = gormDB

	run := m.Run()
	sqlDB.Close()

	os.Exit(run)
}

func TestUserAdd(t *testing.T) {
	//ユーザが正しく作成されるのか？
	t.Run("Success", func(t *testing.T) {
		req := helper.Request(t, requestPass+"/add_users/success.json", 1)
		res := response.AddUsersResponse(req)

		assert.Equal(t, res.Name, "豊田")
		assert.Equal(t, res.Email, "h.toyoda@aaaaa.com")
	})
}
