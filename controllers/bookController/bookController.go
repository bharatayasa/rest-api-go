package bookcontrollers

import (
	"github.com/gofiber/fiber/v2"
)

func Index(c *fiber.Ctx) error {
	return c.Status(200).JSON(fiber.Map{
		"Message": "hallo",
	})
}

func Show(c *fiber.Ctx) error {
	return nil
}

func Create(c *fiber.Ctx) error {
	return nil
}

func Update(c *fiber.Ctx) error {
	return nil
}

func Delete(c *fiber.Ctx) error {
	return nil
}
