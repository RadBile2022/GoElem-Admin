package constants

import "github.com/gofiber/fiber/v2"

func JsonRequestNotValid(data interface{}) fiber.Map {
	return fiber.Map{
		"error": data,
	}
}
