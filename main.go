package main

import (
	"log"
	"os"

	"busapp/database"
	"busapp/routers"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func getenv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}

func main() {
	godotenv.Load()
	app := fiber.New()
	DB := database.Init()
	routers.Initalize(app, DB)
	log.Fatal(app.Listen(":" + getenv("PORT", "3000")))
}
