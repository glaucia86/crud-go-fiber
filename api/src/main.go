package src

import "github.com/gofiber/fiber"

// Routes função responsável por lidar com todas as rotas da aplicação
func RouteDefault(app *fiber.App) {
	app.Get("/api", func(c *fiber.Ctx) {
		c.Send("Sejam bem-vindos(as) a API Go + Fiber + PostGreSQL + Azure!")
	})
}
