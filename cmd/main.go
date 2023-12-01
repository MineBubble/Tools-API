package main

import (
	"github.com/MineBubble/tools_api/db"
	"github.com/MineBubble/tools_api/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	db.Init()

	router := gin.Default()
	router.GET("/tools", handlers.ListarFerramentas)
	router.POST("/tools", handlers.CriarFerramenta)
	router.DELETE("/tools/:id", handlers.DeletarFerramenta)

	router.Run(":3000")

}
