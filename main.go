package main

import (
	"github.com/arahkya/antifake-news/handlers"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", handlers.ListContent)
	app.Get("/{id}", handlers.GetContent)
	app.Delete("/{id}", handlers.DeleteContent)
	app.Post("/", handlers.PostContent)
	app.Patch("/", handlers.PatchContent)

	app.Listen((":3000"))
}
