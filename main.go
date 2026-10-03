package main

import (
	"log"

	"github.com/GanisSayogyo/UTS_PBL_Book_Management_API_434241060/app/handler"
	"github.com/GanisSayogyo/UTS_PBL_Book_Management_API_434241060/app/repository"
	"github.com/GanisSayogyo/UTS_PBL_Book_Management_API_434241060/app/service"
	"github.com/GanisSayogyo/UTS_PBL_Book_Management_API_434241060/config"
	"github.com/GanisSayogyo/UTS_PBL_Book_Management_API_434241060/database"
	"github.com/GanisSayogyo/UTS_PBL_Book_Management_API_434241060/route"
	"github.com/gofiber/fiber/v2"
)

func main() {
	cfg := config.LoadConfig()

	db, err := database.ConnectPostgres(cfg)
	if err != nil {
		log.Fatal("Database connection failed:", err)
	}
	defer db.Close()

	userRepo := repository.NewUserRepository(db)
	authService := service.NewAuthService(userRepo)
	authHandler := handler.NewAuthHandler(authService)

	app := fiber.New()

	route.SetupRoutes(app, authHandler)

	log.Fatal(app.Listen(":" + cfg.AppPort))
}