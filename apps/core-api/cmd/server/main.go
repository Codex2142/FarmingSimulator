package main

import (
	"context"
	"log"

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
	defer db.Close(context.Background())

	log.Fatal(app.Listen(":3000"))
}
