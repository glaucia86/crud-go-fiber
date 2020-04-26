package main

import "github.com/gofiber/fiber"

func main() {
	app := fiber.New()

	app.Get("/api", func(c *fiber.Ctx) {
		c.Send("Sejam bem-vindos(as) a API Go + Fiber + PostGreSQL + Azure!")
	})

	app.Listen(3000)
}
