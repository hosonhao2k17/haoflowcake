package common

import "github.com/gofiber/fiber/v3"

type CustomResponse struct {
	Data       any    `json:"data"`
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
}

func NewCustomResponse(data any, message string, statusCode int) *CustomResponse {
	return &CustomResponse{
		Data:       data,
		Message:    message,
		StatusCode: statusCode,
	}
}

func ResponseOK(c fiber.Ctx, data any, message string) error {
	response := NewCustomResponse(data, message, c.Response().StatusCode())
	return c.Status(fiber.StatusOK).JSON(response)
}

func ResponseCreated(c fiber.Ctx, data any, message string) error {
	response := NewCustomResponse(data, message, c.Response().StatusCode())
	return c.Status(fiber.StatusCreated).JSON(response)
}

func ResponseNoContent(c fiber.Ctx) error {
	return c.SendStatus(fiber.StatusNoContent)
}
