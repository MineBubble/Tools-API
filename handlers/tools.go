package handlers

import (
	"fmt"

	"github.com/MineBubble/tools_api/db"
	"github.com/MineBubble/tools_api/models"
	"github.com/gin-gonic/gin"
)

func ListarFerramentas(c *gin.Context) {
	var tools []models.Tool
	if err := db.DB.Find(&tools).Error; err != nil {
		c.AbortWithStatus(500)
		fmt.Println(err)
	} else {
		c.JSON(200, tools)
	}
}

func DeletarFerramenta(c *gin.Context) {
	var tool []models.Tool
	if err := db.DB.Where("id = ?", c.Param("id")).Delete(&tool).Error; err != nil {
		c.AbortWithStatus(500)
		fmt.Println(err)
	} else {
		c.JSON(200, tool)
	}
}
