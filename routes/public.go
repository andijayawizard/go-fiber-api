package routes

import (
	"github.com/andijayawizard/go-fiber-api/controllers"
	"github.com/andijayawizard/go-fiber-api/middleware"
	"github.com/gofiber/fiber/v2"
)

func PublicRoutes(app *fiber.App) {
	group := app.Group("/public", middleware.RequireAPIKey)
	group.Get("/books", controllers.GetBooks)    // Read only
	group.Get("/books/:id", controllers.GetBook) // Read only

}
