package routes

import (
	"farming/internal/handler"
	"farming/internal/repository"
	"farming/internal/usecase"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5"
)

// ====================================================================
// SetupRoutes menerima parameter app (pointer ke Fiber App) dan db (pointer ke pgx.Conn)
// Fungsi ini digunakan untuk mendefinisikan semua route HTTP pada aplikasi
func SetupRoutes(app *fiber.App, db *pgx.Conn) {

	// Membuat grup route dengan prefix "/api"
	// Semua route yang didefinisikan di bawah api akan diawali "/api"
	api := app.Group("/api")

	// =========================================================
	// Endpoint GET "/api/health" untuk memeriksa status server
	api.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "OK"})
	})

	// =========================================================
	// (USER) Membuat repository user dengan koneksi database
	userRepo := repository.NewUserRepo(db)
	// Membuat usecase user yang menggunakan repository
	userUC := usecase.NewUserUsecase(userRepo)
	// Membuat handler user yang menggunakan usecase
	userHandler := handler.NewUserHandler(userUC)

	// ENDPOINT (users CRUD)
	api.Post("/users", userHandler.CreateUser)
	api.Get("/users/:id", userHandler.GetUser)
	api.Put("/users/:id", userHandler.UpdateUser)
	api.Delete("users/:id", userHandler.DeleteUser)
}
