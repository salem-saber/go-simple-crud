package main

import (
	"github.com/salemsaber/03_crud/initializers"
	"github.com/salemsaber/03_crud/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})

}
