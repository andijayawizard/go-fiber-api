package main

import (
	"log"

	"github.com/andijayawizard/go-fiber-api/database"
	"github.com/andijayawizard/go-fiber-api/handlers"
	"github.com/andijayawizard/go-fiber-api/middleware"
	"github.com/andijayawizard/go-fiber-api/models"
	"github.com/andijayawizard/go-fiber-api/seed"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

// func initConfig() {
// 	env := os.Getenv("APP_ENV")
// 	if env == "" {
// 		env = "dev" // default dev kalau kosong
// 	}

//		viper.SetConfigFile(".env." + env)
//		err := viper.ReadInConfig()
//		if err != nil {
//			log.Fatalf("❌ Error loading config file: %v", err)
//		}
//	}
func main() {
	// initConfig()
	// Load .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	app := fiber.New()

	// Connect DB & migrate
	database.Connect()
	database.DB.AutoMigrate(&models.Book{})
	seed.SeedBooks()

	// Aktifkan CORS di sini
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173", // frontend Vue kamu
		AllowHeaders: "Origin, Content-Type, Accept, X-API-Key",
		AllowMethods: "GET,POST,PUT,DELETE",
	}))
	// Global middleware
	app.Use(middleware.RequireAPIKey)

	// Routes
	app.Get("/books", handlers.GetBooks)
	app.Get("/books/:id", handlers.GetBook)
	app.Post("/books", handlers.CreateBook)
	app.Put("/books/:id", handlers.UpdateBook)
	app.Delete("/books/:id", handlers.DeleteBook)

	app.Listen(":8080")
}
