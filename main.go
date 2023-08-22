package main

import (
	"log"

	"github.com/ferchox920/fiber-api/database"
	"github.com/ferchox920/fiber-api/routes"
	"github.com/gofiber/fiber/v2"
)

func welcome(c *fiber.Ctx) error {
	return c.SendString("Welcome to my API")
}

func setupRoutes(app *fiber.App) {
	// Welcome endpoint
	app.Get("/api", welcome)
	app.Post("/api/users", routes.CreateUser)
	app.Get("/api/users/:id", routes.FindUserByID)
	app.Get("/api/users", routes.FindAllUsers)
	app.Put("/api/users/:id", routes.UpdateUser)
	app.Delete("/api/users/:id", routes.DeleteUser)

}

func main() {

	database.ConnectDb()
	app := fiber.New()

	setupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
