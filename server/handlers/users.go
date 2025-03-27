package handlers

import (
	"imobiliaria/internal/models"
	"imobiliaria/server/handlers/errors"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

// Field names should start with an uppercase letter
type User struct {
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
	Age   *int   `json:"age" validate:"required,gte=0,lte=130"`
}

func (h *Handler) CreateUser(c *fiber.Ctx) error {
	u := new(User)
	if err := c.BodyParser(u); err != nil {
		return err
	}

	/// Valida o que tem na declaração da struct
	if err := h.Validator.Struct(u); err != nil {
		logrus.Println(err)

		return &errors.Error{
			Message: "Invalid user data",
			Status:  http.StatusBadRequest,
		}
	}

	result, err := h.Usecases.CreateUser(c.Context(), &models.User{
		Name:  u.Name,
		Email: u.Email,
		Age:   u.Age,
	})
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(result)
}

func (h *Handler) GetUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return fiber.ErrBadRequest
	}

	user, err := h.Usecases.GetUser(c.Context(), id)

	logrus.Infoln(user)

	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(user)
}
