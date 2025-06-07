package main

import (
	"assignment/internals/adapter"
	"assignment/internals/core/handler"
	"assignment/internals/core/services"
	"assignment/internals/helper"
	"assignment/internals/repositories"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println(".env file not found or could not be loaded")
	}

	cfg := adapter.Load()
	dsn := cfg.DSN()

	db, err := adapter.NewMySQLDatabase(dsn)
	if err != nil {
		log.Fatal("DB connection failed:", err)
	}

	// InitializeDatabase(db, dsn)

	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://127.0.0.1:5500/",
		AllowCredentials: true,
	}))
	api := app.Group("/api")

	userHandler.RegisterRoutes(api)

	log.Fatal(app.Listen(":8081"))
}

func InitializeDatabase(db *gorm.DB, dsn string) {
	err := helper.GenerateModels(db)
	if err != nil {
		log.Fatal("Model generation failed:", err)
	}

	err = helper.RunSQLFilesInFolder(dsn, "mock")
	if err != nil {
		log.Fatal("Running SQL files failed:", err)
	}
}
