package main

import (
	"api-livro/src"

	"github.com/gofiber/fiber"
)

// main função responsável por levantar o servidor da aplicação em Go
func main() {
	app := fiber.New()
	src.Routes(app)
	app.Listen(3000)
}
