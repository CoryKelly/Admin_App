package main

import (
	"github.com/CoryKelly/Admin_App/db"
	"github.com/CoryKelly/Admin_App/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"log"
)

func main() {
	//Database Connection
	db.Connect()

	//Server Setup
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))
	routes.Setup(app)
	log.Fatal(app.Listen(":3000"))
}
