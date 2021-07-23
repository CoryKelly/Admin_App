package routes

import (
	"github.com/CoryKelly/Admin_App/controllers"
	"github.com/CoryKelly/Admin_App/middleware"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	//Public
	app.Post("/api/register", controllers.Register)
	app.Post("/api/login", controllers.Login)

	//Middleware
	app.Use(middleware.IsAuthenticated)

	//Private
	app.Post("/api/logout", controllers.Logout)
	app.Get("/api/user", controllers.User)

	app.Get("/api/user/:id", controllers.GetUser)
	app.Put("/api/user/:id", controllers.UpdateUser)
	app.Delete("/api/user/:id", controllers.DeleteUser)
	app.Post("/api/users", controllers.CreateUser)
	app.Get("/api/users", controllers.GetAllUsers)

	app.Get("/api/roles/:id", controllers.GetRole)
	app.Put("/api/roles/:id", controllers.UpdateRole)
	app.Delete("/api/roles/:id", controllers.DeleteRole)
	app.Post("/api/roles", controllers.CreateRole)
	app.Get("/api/roles", controllers.GetAllRoles)

	app.Get("/api/permissions", controllers.GetAllPermissions)
}
