package main

import (
	"github.com/bharatayasa/rest-api-go/controllers/bookcontroller"
	"github.com/bharatayasa/rest-api-go/models"

	"github.com/gofiber/fiber/v2"
)

func main() {
	models.ConnectDatabase()
	app := fiber.New()

	api := app.Group("/api")
	book := api.Group("/books")

	book.Get("/", bookcontroller.Index)
	book.Get("/:id", bookcontroller.Show)
	book.Post("/", bookcontroller.Create)
	book.Put("/:id", bookcontroller.Update)
	book.Delete("/:id", bookcontroller.Delete)

	app.Listen(":8000")
}
