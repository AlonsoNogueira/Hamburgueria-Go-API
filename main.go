package main

import (
	"log"

	config "github.com/alnszzx/HamburgueriaGo/Config"
	model "github.com/alnszzx/HamburgueriaGo/Database/Model"
	repository "github.com/alnszzx/HamburgueriaGo/Database/Repository"
	handler "github.com/alnszzx/HamburgueriaGo/Handler"
	routes "github.com/alnszzx/HamburgueriaGo/Routes"
	service "github.com/alnszzx/HamburgueriaGo/Service"
	"github.com/gin-gonic/gin"
)

func main() {
	db := config.ConnectPostgres()

	err := db.AutoMigrate(
		&model.Employer{},
		&model.Snacks{},
	)

	if err != nil {
		log.Fatal("Error to execute migrations")
	}

	log.Println("Migrations applied")

	//inicialiação
	snackRepository := repository.NewSnacksRepository(db)
	snackService := service.NewSnackService(snackRepository)
	snackHandler := handler.NewSnacksHandler(snackService)

	router := gin.Default()

	routes.SnackRoutes(router, snackHandler)

	log.Println("server running on :3000")

	if err := router.Run(":3000"); err != nil {
		log.Fatal(err)
	}
}
