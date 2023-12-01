package main

import (
	"github.com/MineBubble/tools_api/db"
)

func main() {
	// router := gin.Default()
	// router.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "pong",
	// 	})
	// })

	// router.Run(":3000")

	db.Init()
}
