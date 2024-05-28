package main

import "github.com/gin-gonic/gin"

func main() {

	router := gin.Default()
	// var err error

	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"code": 200,
			"msg":  "pong",
		})
	})

	router.Run(":8888")
}
