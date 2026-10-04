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
	bookHandler *handler.BookHandler,
	cfg config.Config,
) {
	api := app.Group("/api/v1")

	api.Get("/health", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "success",
			"data": fiber.Map{
				"status": "ok",
			},
		})
	})

	auth := api.Group("/auth")
	auth.Post("/register", authHandler.Register)
	auth.Post("/login", authHandler.Login)

	users := api.Group("/users", middleware.RequireAuth(cfg))
	users.Get("/me", userHandler.Me)

	books := api.Group("/books", middleware.RequireAuth(cfg))
	books.Get("/", bookHandler.GetAll)
	books.Get("/:id", bookHandler.GetByID)
	books.Post("/", middleware.RequireAdmin, bookHandler.Create)
	books.Put("/:id", middleware.RequireAdmin, bookHandler.Update)
	books.Patch("/:id", middleware.RequireAdmin, bookHandler.Patch)
	books.Delete("/:id", middleware.RequireAdmin, bookHandler.Delete)
}
