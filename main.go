package main

import (
	"log"
	"github.com/gofiber/fiber/v2"
	"github.com/jaykeHarrison/go-rest-api/database"
	"github.com/jaykeHarrison/go-rest-api/routes"
)

func welcome(c *fiber.Ctx) error {
	return c.SendString("Welcome to my api")
}

func setupRoutes(app *fiber.App) {
	app.Get("/api", welcome)
	app.Post("/api/users", routes.CreateUser)
	app.Get("/api/users", routes.GetUsers)
	app.Get("/api/users/:id", routes.GetUserByID)
	app.Delete("/api/users/:id", routes.DeleteUserByID)
}

func main() {
	database.ConnectDb()

	app := fiber.New()

	setupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}