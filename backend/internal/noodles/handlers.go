package noodles

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
	"github.com/joseemds/pasta/.gen/pasta/public/model"
)

func RegisterNoodleGroup(app fiber.Router) {
	app.Post("/", createNoodle)

}

func createNoodle(c fiber.Ctx) error{
	type RequestBody struct {
		Noodles []model.Noodle `json:"noodles" validate:"min=1"`
	}


	validate := validator.New()
	req := new(RequestBody)


	if err := c.Bind().JSON(req); err != nil {
		log.Errorf("Error when parsing json %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse JSON",
		})
	}


	if err:= validate.Struct(req); err != nil  {
		log.Errorf("Error when validating struct %v", req)
		errors := err.(validator.ValidationErrors)
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"error": fmt.Sprintf("Validation error: %s", errors),
		})

	}

	log.Info("Creating noodle")
	return c.SendStatus(fiber.StatusCreated)
}
