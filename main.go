package main

import (
	"log"
	"os"

	"github.com/andijayawizard/go-fiber-api/controllers"
	"github.com/andijayawizard/go-fiber-api/database"
	"github.com/andijayawizard/go-fiber-api/middleware"
	"github.com/andijayawizard/go-fiber-api/models"
	"github.com/andijayawizard/go-fiber-api/routes"
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
func LoadConfig() {
	// Opsional: hanya untuk local/dev
	if os.Getenv("ENV") != "production" {
		if err := godotenv.Load(); err != nil {
			log.Println("Warning: error loading .env file, fallback to system env")
		}
	}
}
func main() {
	LoadConfig()

	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }
	if os.Getenv("JWT_SECRET") == "" {
		log.Fatal("🚨 JWT_SECRET tidak ditemukan di .env — HARUS DISET untuk keamanan!")
	}
	app := fiber.New()

	// Connect DB & migrate
	database.Connect()
	database.DB.AutoMigrate(&models.User{})
	database.DB.AutoMigrate(&models.Book{})
	seed.SeedBooks()

	// Aktifkan CORS di sini
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173", // frontend Vue kamu
		AllowHeaders: "Origin, Content-Type, Accept, X-API-Key",
		AllowMethods: "GET,POST,PUT,DELETE",
	}))
	// Global middleware
	// app.Use(middleware.RequireAPIKey)

	// ✳️ Public route
	app.Post("/auth/register", controllers.Register)
	app.Post("/auth/login", controllers.Login)
	app.Post("/auth/logout", middleware.RequireAuth, controllers.Logout)

	// ✅ Group yang butuh JWT
	routes.AdminRoutes(app)
	// ✅ Group yang butuh API key (opsional)
	routes.PublicRoutes(app)
	app.Listen(":8080")
}
