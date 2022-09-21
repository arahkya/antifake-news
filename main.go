package main

import (
	"github.com/arahkya/antifake-news/handlers"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", handlers.ListContent)
	app.Post("/", handlers.PostContent)

	app.Listen((":3000"))
}
