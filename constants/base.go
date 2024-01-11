package constants

import "github.com/gofiber/fiber/v2"

const (
	CodeSuccess             = fiber.StatusOK
	CodeCreated             = fiber.StatusCreated
	CodeBadRequest          = fiber.StatusBadRequest
	CodeForbidden           = fiber.StatusForbidden
	CodeErrDataNotFound     = fiber.StatusNotFound
	CodeErrRequestNotValid  = fiber.StatusUnprocessableEntity
	CodeInternalServerError = fiber.StatusInternalServerError
)
