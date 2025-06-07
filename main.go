// @title Assignment API
// @version 1.0
// @description This is an API for the assignment project.
// @host localhost:8081
// @BasePath /api
package main

import (
	_ "assignment/docs"
	"assignment/internals/adapter"
	"assignment/internals/core/handler"
	"assignment/internals/core/services"
	"assignment/internals/helper"
	"assignment/internals/repositories"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	fiberSwagger "github.com/swaggo/fiber-swagger"
	"gorm.io/gorm"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println(".env file not found or could not be loaded")
	}

	appPort := os.Getenv("APP_PORT")

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

	app.Get("/swagger/*", func(c *fiber.Ctx) error {
		fmt.Println("Swagger called:", c.Path())
		return fiberSwagger.WrapHandler(c)
	})

	log.Fatal(app.Listen(fmt.Sprintf(":%s", appPort)))
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
