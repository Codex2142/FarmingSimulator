package main

import (
	"context"
	"log"

	"farming/migrations"
	"farming/pkg/database"
	"farming/routes"

	"farming/internal/config"

	"github.com/gofiber/fiber/v2"
)

// @title Farming API
// @version 1.0
// @description API untuk sistem farming
// @host localhost:3000
// @BasePath /api
func main() {

	config := config.LoadConfig()

	db := database.Connect(config.DBUrl)
	defer db.Close(context.Background())

	log.Println("[Running] Migrating Databases")
	migrations.RunMigration()

	app := fiber.New()

	routes.SetupRoutes(app, db)

	log.Println("Server running on http://localhost:" + config.AppPort)
	log.Println("Swagger docs at http://localhost:" + config.AppPort + "/swagger/index.html")

	log.Fatal(app.Listen(":" + config.AppPort))
}
