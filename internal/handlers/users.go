package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

// Field names should start with an uppercase letter
type User struct {
	Name string `json:"name" xml:"name" form:"name"`
}

func (h *Handler) CreateUser(c *fiber.Ctx) error {
	u := new(User)
	if err := c.BodyParser(u); err != nil {
		return err
	}

	result, err := h.Usecases.CreateUser(c.Context(), u.Name)
	logrus.Infoln(result)

	if err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(result)
}

func (h *Handler) GetUser(c *fiber.Ctx) error {
	user, err := h.Usecases.GetUser(c.Context(), c.Params("id"))

	logrus.Infoln(user)

	if err != nil {
		return err
	}

	return c.Status(fiber.StatusAccepted).JSON(user)
}
