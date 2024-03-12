package main

import (
	bookcontrollers "github.com/bharatayasa/rest-api-go/controllers/bookController"
	"github.com/bharatayasa/rest-api-go/models"

	"github.com/gofiber/fiber/v2"
)

func main() {
	models.ConnectDatabase()
	app := fiber.New()

	api := app.Group("/api")
	book := api.Group("/books")

	book.Get("/", bookcontrollers.Index)
	book.Get("/:id", bookcontrollers.Show)
	book.Post("/", bookcontrollers.Create)
	book.Put("/:id", bookcontrollers.Update)
	book.Delete("/:id", bookcontrollers.Delete)

	app.Listen(":3000")
}
