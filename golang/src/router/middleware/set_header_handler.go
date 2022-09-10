package middleware

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/htoyoda18/sample-tweet-api/golang/src/shared/config"
	"github.com/htoyoda18/sample-tweet-api/golang/src/shared/db"
)

// ginのコンテキストにDBの接続情報を格納する
func SetHeaderHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		conf, err := config.InitConfiguration()

		if err != nil {
			log.Fatal(err.Error())
		}

		db, sqlDB, dbErr := db.InitDB(conf)

		defer sqlDB.Close()

		if dbErr != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, dbErr)
			return
		}

		c.Set("db", db)
	}
}
