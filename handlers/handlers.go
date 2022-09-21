package handlers

import (
	"encoding/json"

	data "github.com/arahkya/antifake-news/Data"
	"github.com/arahkya/antifake-news/models"
	"github.com/gofiber/fiber/v2"
)

func HelloWorld(c *fiber.Ctx) error {
	return c.SendString("Hello World")
}

func ListContent(c *fiber.Ctx) error {
	contents := data.ListContent()

	bJson, _ := json.Marshal(contents)

	return c.SendString(string(bJson))
}

func PostContent(c *fiber.Ctx) error {
	content := models.ContentModel{}

	if err := c.BodyParser(&content); err != nil {
		return err
	}

	data.CreateContent(&content)

	return c.SendStatus(200)
}
