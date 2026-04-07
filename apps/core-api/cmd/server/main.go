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
	app := fiber.New()

	routes.SetupRoutes(app)

	config := config.LoadConfig()

	db := database.Connect(config.DBUrl)

	log.Println("[Running] Migrating Databases")
	migrations.RunMigration()
	defer db.Close(context.Background())

	log.Fatal(app.Listen(":3000"))
}
