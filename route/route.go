package route

import (
	"github.com/GanisSayogyo/UTS_PBL_Book_Management_API_434241060/app/handler"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, authHandler *handler.AuthHandler) {
	api := app.Group("/api/v1")

	auth := api.Group("/auth")
	auth.Post("/register", authHandler.Register)
}