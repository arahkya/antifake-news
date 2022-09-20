package handlers

import (
	data "github.com/arahkya/antifake-news/Data"
	"github.com/arahkya/antifake-news/models"
	"github.com/gofiber/fiber/v2"
)

func HelloWorld(c *fiber.Ctx) error {
	return c.SendString("Hello World")
}

func PostContent(c *fiber.Ctx) error {
	content := models.ContentModel{}

	if err := c.BodyParser(&content); err != nil {
		return err
	}

	data.CreateContent(&content)

	return c.SendStatus(200)
}
