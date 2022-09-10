package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/htoyoda18/sample-tweet-api/golang/src/shared/db"
)

// ginのコンテキストにDBの接続情報を格納する
func SetHeaderHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		db, sqlDB, dbErr := db.InitDB()

		defer sqlDB.Close()

		if dbErr != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, dbErr)
			return
		}

		c.Set("db", db)

		c.Next()
	}
}
