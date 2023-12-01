package db

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	stringConnection := "host=localhost port=5432 user=postgres password=admin dbname=tools_api"

	database, err := gorm.Open(postgres.Open(stringConnection), &gorm.Config{})
	if err != nil {
		fmt.Println("Erro na conexão")
	}

	DB = database

	fmt.Println("Conexão bem sucedida!")
}
