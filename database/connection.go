package database

import (
	"log"

	// Removed handlers import to avoid import cycle
	"github.com/andijayawizard/go-fiber-api/middleware"
	"github.com/andijayawizard/go-fiber-api/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB is a global variable to hold the database connection
var DB *gorm.DB

func Connect() {
	dsn := "host=localhost user=postgres password=postgres dbname=bookdb port=5432 sslmode=disable"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ Failed to connect to database:", err)
	}
	log.Println("✅ Database connected successfully!")
}

func database() {
	app := fiber.New()
	DB.AutoMigrate(&models.Book{})
	// Connect DB & migrate
	Connect()
	DB.AutoMigrate(&models.Book{})

	// Global middleware
	app.Use(middleware.RequireAPIKey)

	// Routes
	// Define routes without directly using handlers to avoid import cycle
	app.Get("/books", func(c *fiber.Ctx) error {
		// Call the appropriate handler logic here
		return c.SendString("GetBooks handler logic")
	})
	app.Get("/books/:id", func(c *fiber.Ctx) error {
		// Call the appropriate handler logic here
		return c.SendString("GetBook handler logic")
	})
	app.Post("/books", func(c *fiber.Ctx) error {
		// Call the appropriate handler logic here
		return c.SendString("CreateBook handler logic")
	})
	app.Put("/books/:id", func(c *fiber.Ctx) error {
		// Call the appropriate handler logic here
		return c.SendString("UpdateBook handler logic")
	})
	app.Delete("/books/:id", func(c *fiber.Ctx) error {
		// Call the appropriate handler logic here
		return c.SendString("DeleteBook handler logic")
	})

	app.Listen(":8080")
}
