package constants

import (
	"github.com/gofiber/fiber/v2"
)

func JsonRes(code int, data interface{}) fiber.Map {
	switch code {
	case CodeSuccess:
		return JsonDataSuccess(data)
	case CodeForbidden:
		return JsonForbidden(data)
	case CodeErrDataNotFound:
		return JsonNotFound(data)
	default:
		return JsonErrorInternal(data)
	}
}

func JsonDataSuccess(data interface{}) fiber.Map {
	res := fiber.Map{
		"data": data,
	}
	return res
}

func JsonForbidden(data interface{}) fiber.Map {
	return fiber.Map{
		"error": data,
	}
}

func JsonNotFound(data interface{}) fiber.Map {
	return fiber.Map{
		"error": data,
	}
}

func JsonErrorInternal(data interface{}) fiber.Map {
	return fiber.Map{
		"error": data,
	}
}
