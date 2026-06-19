package constant

import "github.com/gofiber/fiber/v3"

type ErrorCode string

const (
	//COMMON
	ErrBadRequest     ErrorCode = "BAD_REQUEST"
	ErrInternalServer ErrorCode = "INTERNAL_SERVER_ERROR"

	//ROLE
	ErrRoleNotFound  ErrorCode = "ROLE_IS_NOT_FOUND"
	ErrRoleKeyExists ErrorCode = "ROLE_KEY_IS_CONFLICT"
)

type MapError struct {
	Message    string
	StatusCode int
}

var MapErrorCode = map[ErrorCode]MapError{
	ErrRoleNotFound: {
		Message:    "role is not found",
		StatusCode: fiber.StatusNotFound,
	},
	ErrRoleKeyExists: {
		Message:    "role key is exists",
		StatusCode: fiber.StatusConflict,
	},
}
