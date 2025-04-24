package handlers

import (
	"imobiliaria/internal/models"
	"imobiliaria/server/handlers/errors"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/sirupsen/logrus"
)

// Field names should start with an uppercase letter
type User struct {
	Name     string `json:"name"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
	Age      *int   `json:"age" validate:"gte=0,lte=130"`
}

type UserLogin struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
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
		Name:         u.Name,
		Email:        u.Email,
		Age:          u.Age,
		PasswordHash: u.Password,
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

func (h *Handler) Login(c *fiber.Ctx) error {
	u := new(UserLogin)
	if err := c.BodyParser(u); err != nil {
		return err
	}

	/// Valida o que tem na declaração da struct
	if err := h.Validator.Struct(u); err != nil {
		logrus.Println(err)

		return &errors.Error{
			Message: "Invalid login data",
			Status:  http.StatusBadRequest,
		}
	}

	result, err := h.Usecases.Login(c.Context(), u.Email, u.Password)
	if err != nil {
		logrus.WithError(err).Error("Error logging in")

		return err
	}

	store := c.Locals("sessionStorage").(*session.Store)
	sess, err := store.Get(c)

	
	if err != nil {
		logrus.WithError(err).Error("Error getting session")
		
		return err
	}
	
	sess.Set("user", result.Email)
	if err := sess.Save(); err != nil {
		logrus.WithError(err).Error("Error saving session")

		return err
	}

	c.Locals("user", result)

	return c.Status(fiber.StatusOK).JSON(result)
}
