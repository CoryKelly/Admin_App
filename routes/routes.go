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
	app.Put("/api/users/info", controllers.UpdateInfo)
	app.Put("/api/users/password", controllers.UpdatePassword)

	app.Post("/api/logout", controllers.UpdateUser)
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

	app.Get("/api/products/:id", controllers.GetProduct)
	app.Put("/api/products/:id", controllers.UpdateProduct)
	app.Delete("/api/products/:id", controllers.DeleteProduct)
	app.Post("/api/products", controllers.CreateProduct)
	app.Get("/api/products", controllers.GetAllProducts)

	app.Get("/api/orders", controllers.GetAllOrders)
	app.Post("/api/export", controllers.Export)
	app.Get("/api/chart", controllers.Chart)

	app.Post("/api/upload", controllers.Upload)
	app.Static("/api/uploads", "./uploads")
}
