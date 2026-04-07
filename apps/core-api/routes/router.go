package routes

import (
	"farming/internal/handler"
	"farming/internal/repository"
	"farming/internal/usecase"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5"
)

func SetupRoutes(app *fiber.App, db *pgx.Conn) {
	api := app.Group("/api")

	// Healthcheck
	api.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "OK"})
	})

	// User module
	userRepo := repository.NewUserRepo(db)
	userUC := usecase.NewUserUsecase(userRepo)
	userHandler := handler.NewUserHandler(userUC)

	api.Post("/users", userHandler.CreateUser)
	api.Get("/users/:id", userHandler.GetUser)
}
