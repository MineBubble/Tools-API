package handlers

import (
	"fmt"

	"github.com/MineBubble/tools_api/db"
	"github.com/MineBubble/tools_api/models"
	"github.com/gin-gonic/gin"
)

func ListarFiltrar(c *gin.Context) {
	if c.Query("tag") == "" {
		ListarFerramentas(c)
	} else {
		FiltrarFerramentaPorTag(c)
	}
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

func FiltrarFerramentaPorTag(c *gin.Context) {
	var tools []models.Tool
	var filtrados []models.Tool
	if err := db.DB.Find(&tools).Error; err != nil {
		c.AbortWithStatus(500)
		fmt.Println(err)
	}

	tag := c.Request.URL.Query().Get("tag")

	for i := 0; i < len(tools); i++ {
		for j := 0; j < len(tools[i].Tags); j++ {
			if tag == tools[i].Tags[j] {
				filtrados = append(filtrados, tools[i])
			}
		}
	}

	c.JSON(200, filtrados)
}

func CriarFerramenta(c *gin.Context) {
	var tool models.Tool
	c.BindJSON(&tool)

	if err := db.DB.Create(&tool).Error; err != nil {
		c.AbortWithStatus(500)
		fmt.Println(err)
	} else {
		c.JSON(200, tool)
	}

}

func DeletarFerramenta(c *gin.Context) {
	var tool models.Tool
	if err := db.DB.Where("id = ?", c.Param("id")).Delete(&tool).Error; err != nil {
		c.AbortWithStatus(500)
		fmt.Println(err)
	} else {
		c.JSON(200, tool)
	}
}
