package validator

import (
	"Server/models"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var Validate = validator.New()

func validateUser(c *fiber.Ctx) error {

	var errors []*models.ErrorModel
	var body models.UserModel

	if err := c.BodyParser(&body); err != nil {
		return err
	}

	err := Validate.Struct(body)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var el models.ErrorModel
			el.Field = err.Field()
			el.Tag = err.Tag()
			errors = append(errors, &el)

		}
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"errors": errors})
	}
	fmt.Println("everything ok")
	return c.Next()
}
