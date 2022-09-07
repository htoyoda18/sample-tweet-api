package main

import "github.com/htoyoda18/sample-tweet-api/router"

// @title Tweet API
// @description 簡易的なTwitterに近い REST APIです

// @host localhost:8080
// @BasePath /v1
func main() {
	r := router.SetupRouter()
	r.Run(":8080")
}
