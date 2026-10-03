package middleware

import "github.com/gofiber/fiber/v2"

func RequireAdmin(c *fiber.Ctx) error {
	role := c.Locals("role")

	if role != "admin" {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"message": "admin access required",
		})
	}

	return c.Next()
}