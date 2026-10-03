package route

import (
	"github.com/GanisSayogyo/UTS_PBL_Book_Management_API_434241060/app/handler"
	"github.com/GanisSayogyo/UTS_PBL_Book_Management_API_434241060/config"
	"github.com/GanisSayogyo/UTS_PBL_Book_Management_API_434241060/middleware"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(
	app *fiber.App,
	authHandler *handler.AuthHandler,
	userHandler *handler.UserHandler,
	cfg config.Config,
) {
	api := app.Group("/api/v1")

	auth := api.Group("/auth")
	auth.Post("/register", authHandler.Register)
	auth.Post("/login", authHandler.Login)

	users := api.Group("/users", middleware.RequireAuth(cfg))
	users.Get("/me", userHandler.Me)
}