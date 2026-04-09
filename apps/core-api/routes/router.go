package routes

import (
	"farming/internal/handler"
	"farming/internal/repository"
	"farming/internal/usecase"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5"

	_ "farming/docs"

	fiberSwagger "github.com/swaggo/fiber-swagger"
)

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
	// route swagger
	app.Get("/swagger/*", fiberSwagger.WrapHandler)

	// =========================================================
	// Membuat repository user dengan koneksi database
	userRepo := repository.NewUserRepo(db)
	farmRepo := repository.NewFarmRepo(db)
	// Membuat usecase user yang menggunakan repository
	userUC := usecase.NewUserUsecase(userRepo)
	farmUC := usecase.NewFarmUsecase(farmRepo)
	// Membuat handler user yang menggunakan usecase
	userHandler := handler.NewUserHandler(userUC)
	farmHandler := handler.NewFarmHandler(farmUC)

	// ENDPOINT (users CRUD)
	api.Get("/users", userHandler.GetAllUsers)
	api.Post("/users", userHandler.CreateUser)
	api.Get("/users/:id", userHandler.GetUser)
	api.Put("/users/:id", userHandler.UpdateUser)
	api.Delete("users/:id", userHandler.DeleteUser)

	// ENDPOINT (farms CRUD)
	api.Post("/farms", farmHandler.CreateFarm)
	api.Get("/farms/:id", farmHandler.GetFarm)
	api.Put("/farms/:id", farmHandler.UpdateFarm)
	api.Delete("farms/:id", farmHandler.DeleteUser)

}
