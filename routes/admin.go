package routes

import (
	"github.com/andijayawizard/go-fiber-api/controllers"
	"github.com/andijayawizard/go-fiber-api/middleware"

	"github.com/gofiber/fiber/v2"
)

func AdminRoutes(app *fiber.App) {
	group := app.Group("/books", middleware.RequireAuth)
	group.Get("/", controllers.GetBooks)
	group.Get("/:id", controllers.GetBook)
	group.Post("/", controllers.CreateBook)
	group.Put("/:id", controllers.UpdateBook)
	group.Delete("/:id", controllers.DeleteBook)

}
