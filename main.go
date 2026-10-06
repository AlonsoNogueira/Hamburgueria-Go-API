package main

import (
	"log"
	"os"

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
	port := os.Getenv("PORT")

	err := db.AutoMigrate(
		&model.Employer{},
		&model.Snacks{},
	)

	if err != nil {
		log.Fatal("Error to execute migrations")
	}

	log.Println("Migrations applied")

	//inicialiação - Snack
	snackRepository := repository.NewSnacksRepository(db)
	snackService := service.NewSnackService(snackRepository)
	snackHandler := handler.NewSnacksHandler(snackService)

	//inicialização - Employer
	employerRepository := repository.NewEmployerRepository(db)
	employerService := service.NewEmployerService(employerRepository)
	employerHandler := handler.NewEmployerHandler(employerService)

	router := gin.Default()

	routes.SnackRoutes(router, snackHandler)
	routes.EmployerRoutes(router, employerHandler)

	log.Println("server running on :3000")

	if err := router.Run(port); err != nil {
		log.Fatal(err)
	}
}
