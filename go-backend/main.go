package main

import (
	"log"
	"os"

	"github.com/GlStep/go-sveltekit/db"
	"github.com/gofiber/fiber/v3"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL is required")
	}

	log.Println("Initializing database")
	db.InitDB(dbURL)

	if db.GetDB() == nil {
		log.Fatal("Error initializing database")
	}

	log.Println("Migrating database")
	db.MigrateDB(db.GetDB(), "./migrations")

	app := fiber.New()

	app.Get("/", func(c fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "ok"})
	})

	appPort := os.Getenv("FIBER_PORT")
	if appPort == "" {
		log.Fatal("FIBER_PORT is required")
	}

	log.Println("Starting server on port " + appPort)
	log.Fatal(app.Listen(":" + appPort))
}
