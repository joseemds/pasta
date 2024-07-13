package main

import (
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/healthcheck"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/joseemds/pasta/internal/noodles"
)

func main() {
	app := fiber.New()
	app.Use(logger.New())
	app.Use(healthcheck.NewHealthChecker())

	api := app.Group("/api")
	noodlesGroup := api.Group("/noodles")

	noodles.RegisterNoodleGroup(noodlesGroup)
	log.Fatal(app.Listen(":8080"))
}
