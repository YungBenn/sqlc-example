package main

import (
	"context"
	"log"

	"github.com/gofiber/fiber/v2"

	"sqlc-example/internal/routes"
	"sqlc-example/pkg/config"
	"sqlc-example/pkg/database"
)

func buildServer() error {
	env, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Cannot load .env")
	}

	db, err := database.Connect(context.Background(), env.DBurl)

	if err != nil {
		return err
	}

	app := fiber.New()

	routes.SetupRoutes(app, db)

	log.Fatal(app.Listen(":" + env.Port))

	return nil
}

func main() {
	if err := buildServer(); err != nil {
		log.Fatal(err)
	}
}
