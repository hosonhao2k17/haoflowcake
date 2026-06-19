package middleware

import (
	"errors"
	"haoflowcake/common/dto"

	"github.com/gofiber/fiber/v3"
)

func ErrorHandler(ctx fiber.Ctx, err error) error {
	errorRes := new(dto.ErrorResDto)
	if errors.As(err, &errorRes) {
		return ctx.Status(errorRes.StatusCode).JSON(errorRes)
	}

	return ctx.Status(fiber.StatusInternalServerError).JSON(errorRes)
}
