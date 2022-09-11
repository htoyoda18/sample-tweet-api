package tweet_test

import (
	"log"
	"os"
	"testing"

	"github.com/htoyoda18/sample-tweet-api/golang/src/controller/handler/tweet/test/response"
	sharedDB "github.com/htoyoda18/sample-tweet-api/golang/src/shared/db"
	helper "github.com/htoyoda18/sample-tweet-api/golang/src/test/helper"
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

func TestAddTweet(t *testing.T) {
	//ユーザが正しく作成されるのか？
	t.Run("Success", func(t *testing.T) {
		req := helper.Request(t, requestPass+"/add_tweet/success.json", 1)
		res := response.TweetResponse(req)

		assert.Equal(t, res.Text, "今日も一日お疲れ様!")

		t.Cleanup(func() {
			helper.TeardownFixture(fixturePass)
		})
	})
}
