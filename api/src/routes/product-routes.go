package routes

import (
	"api-product/src/controllers"

	"github.com/gofiber/fiber"
)

// Função responsável por listar todas as rotas da API 'Product'
func ProductRoutes(app *fiber.App) {

	// Rota responsável por listar todos os 'Products': (GET): localhost:3000/api/v1/products
	app.Get("/api/v1/products", controllers.ListAllProducts)

	// Rota responsável por criar um novo 'Product': (POST): localhost:3000/api/v1/products
	app.Post("/api/v1/products", controllers.CreateProduct)

	// Rota responsável por listar um determinado 'Product' pelo Id: (GET): localhost:3000/api/v1/products/:id
	app.Get("/api/v1/products/:id", controllers.FindProductById)

	// Rota responsável por atualizar um determinado 'Product' pelo Id: (PUT): localhost:3000/api/v1/products/:id
	app.Put("/api/v1/products/:id", controllers.UpdateProductById)

	// Rota responsável por excluir um determinado 'Product' pelo Id: (DELETE): localhost:3000/api/v1/products/:id
	app.Delete("/api/v1/products/:id", controllers.DeleteProductById)
}
