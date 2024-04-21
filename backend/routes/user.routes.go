package routes

import (
	"github.com/Harrieson/golangbackend/controllers"
	"github.com/Harrieson/golangbackend/middleware"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Post("/api/users/register", controllers.Register)
	app.Post("/api/users/login", controllers.Login)

	app.Use(middleware.IsAuthenticate)
	app.Post("/api/posts/add", controllers.CreatePost)
	app.Get("/api/posts", controllers.GetAllPosts)
	app.Get("/api/posts/:id", controllers.GetAllPosts)
	app.Put("/api/posts/update/:id", controllers.UpdatePost)
}
