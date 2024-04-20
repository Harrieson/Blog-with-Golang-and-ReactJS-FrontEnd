package routes

import (
	"github.com/Harrieson/golangbackend/controllers"
	"github.com/Harrieson/golangbackend/middleware"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Use(middleware.IsAuthenticate)
	app.Post("/api/users/register", controllers.Register)
}
