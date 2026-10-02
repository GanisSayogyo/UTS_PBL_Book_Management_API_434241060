package main

import (
	"fmt"
	"log"

	"github.com/GanisSayogyo/UTS_PBL_Book_Management_API_434241060/config"
	"github.com/GanisSayogyo/UTS_PBL_Book_Management_API_434241060/database"
	"github.com/gofiber/fiber/v2"
)

func main() {
	cfg := config.LoadConfig()

	db, err := database.ConnectPostgres(cfg)
	if err != nil {
		log.Fatal("Database connection failed:", err)
	}
	defer db.Close()

	fmt.Println("Database connection successful")

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Book Management API")
	})

	log.Fatal(app.Listen(":" + cfg.AppPort))
}
