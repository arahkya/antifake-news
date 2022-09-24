package handlers

import (
	"encoding/json"
	"strconv"

	data "github.com/arahkya/antifake-news/Data/sqlite"
	"github.com/arahkya/antifake-news/models"
	"github.com/gofiber/fiber/v2"
)

func HelloWorld(c *fiber.Ctx) error {
	return c.SendString("Hello World")
}

func ListContent(c *fiber.Ctx) error {
	db := data.GetDatabase()
	contents := data.ListContent(db)

	bJson, _ := json.Marshal(contents)

	return c.SendString(string(bJson))
}

func GetContent(c *fiber.Ctx) error {
	contentId, _ := strconv.Atoi(c.Query("id"))

	db := data.GetDatabase()
	content := data.GetContent(db, &contentId)

	bJson, _ := json.Marshal(content)

	return c.SendString(string(bJson))
}

func PostContent(c *fiber.Ctx) error {
	content := models.ContentModel{}

	if err := c.BodyParser(&content); err != nil {
		return err
	}

	db := data.GetDatabase()
	data.CreateContent(db, &content)

	return c.SendStatus(200)
}

func PatchContent(c *fiber.Ctx) error {
	content := models.ContentModel{}

	if err := c.BodyParser(&content); err != nil {
		return err
	}

	db := data.GetDatabase()
	data.UpdateContent(db, &content)

	return c.SendStatus(200)
}

func DeleteContent(c *fiber.Ctx) error {
	contentId, _ := strconv.Atoi(c.Query("id"))

	db := data.GetDatabase()
	data.DeleteContent(db, &contentId)

	return c.SendStatus(200)
}
