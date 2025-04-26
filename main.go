package main

import (
	"github.com/andijayawizard/go-fiber-api/database"
	"github.com/andijayawizard/go-fiber-api/handlers"
	"github.com/andijayawizard/go-fiber-api/middleware"
	"github.com/andijayawizard/go-fiber-api/models"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// Connect DB & migrate
	database.Connect()
	database.DB.AutoMigrate(&models.Book{})

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
