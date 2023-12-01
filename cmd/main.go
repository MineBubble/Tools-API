package main

import (
	"fmt"

	"github.com/MineBubble/tools_api/db"
	"github.com/MineBubble/tools_api/models"
	"github.com/gin-gonic/gin"
)

func main() {
	db.Init()

	router := gin.Default()
	router.GET("/tools", ListarFerramentas)

	router.Run(":3000")

}

func ListarFerramentas(c *gin.Context) {
	var tools []models.Tool
	if err := db.DB.Find(&tools).Error; err != nil {
		c.AbortWithStatus(500)
		fmt.Println(err)
	} else {
		c.JSON(200, tools)
	}
}
