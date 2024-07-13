package main

import (
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/healthcheck"
	"github.com/gofiber/fiber/v3/middleware/logger"
)

func main() {
	app := fiber.New()
	app.Use(logger.New())
	app.Use(healthcheck.NewHealthChecker())
	log.Fatal(app.Listen(":8080"))
}
