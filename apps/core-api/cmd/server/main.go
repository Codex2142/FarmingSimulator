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

func main() {

	config := config.LoadConfig()

	db := database.Connect(config.DBUrl)
	defer db.Close(context.Background())

	log.Println("[Running] Migrating Databases")
	migrations.RunMigration()

	app := fiber.New()

	routes.SetupRoutes(app, db)

	log.Fatal(app.Listen(":" + config.AppPort))
}
