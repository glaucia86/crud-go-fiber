package main

import (
	"api-product/src"

	"api-product/src/routes"

	"api-product/src/database"

	"github.com/gofiber/fiber"
)

// main função responsável por levantar o servidor da aplicação em Go
func main() {
	app := fiber.New()

	database.ReturnDatabase()

	src.RouteDefault(app)

	routes.ProductRoutes(app)

	app.Listen(3000)

}
