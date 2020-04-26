package controllers

import (
	"github.com/gofiber/fiber"
)

func ListAllProducts(c *fiber.Ctx) {
	c.Send("List All Products")
}

func FindProductById(c *fiber.Ctx) {
	c.Send("List Product by Id")
}

func CreateProduct(c *fiber.Ctx) {
	c.Send("Create a new Product")
}

func UpdateProductById(c *fiber.Ctx) {
	c.Send("Update Product by Id")
}

func DeleteProductById(c *fiber.Ctx) {
	c.Send("Delete Product by Id")
}
