package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/htoyoda18/sample-tweet-api/golang/src/shared/logger"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie := c.Request.Header.Get("Authorization")
		arrCookie := strings.Split(cookie, "Bearer ")

		token, err := jwt.Parse(arrCookie[1], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				c.Status(http.StatusUnauthorized)
			}

			return []byte("SECRET_KEY"), nil
		})

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			c.Set("user_id", int(claims["user_id"].(float64)))
		} else {
			logger.Error("AuthMiddleware", err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, err)
		}
	}
}
