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
	
	// User routes
	app.Get("/api/users", routes.FindAllUsers)
	app.Post("/api/users", routes.CreateUser)
	app.Get("/api/users/:id", routes.FindUserByID)
	app.Put("/api/users/:id", routes.UpdateUser)
	app.Delete("/api/users/:id", routes.DeleteUser)
	
	// Product routes
	app.Get("/api/products", routes.GetAllProducts)
	app.Post("/api/products", routes.CreateProduct)
	app.Get("/api/products/:id", routes.FindProductByID)
	app.Put("/api/products/:id", routes.UpdateProduct) 
	app.Delete("/api/products/:id", routes.DeleteProduct) 

	// Order routes
	app.Post("/api/orders", routes.CreateOrder)
	app.Get("/api/orders", routes.GetOrders)
	app.Get("/api/orders/:id", routes.GetOrder)
}

func main() {
	database.ConnectDb()
	app := fiber.New()

	setupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}

