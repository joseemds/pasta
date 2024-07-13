package noodles

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
	"github.com/joseemds/pasta/.gen/pasta/public/model"
)

func RegisterNoodleGroup(app fiber.Router) {
	app.Post("/", createNoodle)

}

func createNoodle(c fiber.Ctx) error{
	type RequestBody struct {
		Noodles []model.Noodle `json:"noodles"`
	}

	req := new(RequestBody)
	if err := c.Bind().JSON(req); err != nil {
		log.Errorf("Error when parsing json %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse JSON",
		})
	}

	log.Info("Creating noodle")
	return c.SendStatus(fiber.StatusCreated)
}
