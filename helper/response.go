package helper

import "github.com/gofiber/fiber/v2"

type Response struct {
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func SuccessResponse(c *fiber.Ctx, status int, data any) error {
	return c.Status(status).JSON(Response{
		Message: "success",
		Data:    data,
	})
}

func ErrorResponse(c *fiber.Ctx, status int, message string) error {
	return c.Status(status).JSON(Response{
		Message: message,
	})
}