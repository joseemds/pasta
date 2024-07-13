package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v3"
)

func HealthCheck(c fiber.Ctx) {
	return c.Status(http.StatusOK)
}
