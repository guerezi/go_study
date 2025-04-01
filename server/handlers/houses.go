package handlers

import (
	"imobiliaria/internal/models"
	"imobiliaria/server/handlers/errors"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type House struct {
	ID      int     `json:"id" validate:"required"`
	Street  string  `json:"street" validate:"required"`
	Number  int     `json:"number" validate:"required,gte=0"`
	City    string  `json:"city" validate:"required"`
	State   string  `json:"state" validate:"required"`
	ZipCode string  `json:"zip_code" validate:"required"`
	Price   float64 `json:"price" validate:"required,gte=0"`
	OwnerID *int    `json:"owner_id"`
}

// type FullHouse struct {
// 	House
// 	HouseOwnderID
// }

// type HouseOwnderID struct {
// 	OwnerID *int `json:"owner_id"`
// }

func (h *Handler) CreateHouse(c *fiber.Ctx) error {
	house := new(House)
	if err := c.BodyParser(house); err != nil {
		return err
	}

	if err := h.Validator.Struct(house); err != nil {
		logrus.Println(err)

		return &errors.Error{
			Message: "Invalid house data",
			Status:  http.StatusBadRequest,
		}
	}

	result, err := h.Usecases.CreateHouse(c.Context(), &models.House{
		Street:  house.Street,
		Number:  house.Number,
		City:    house.City,
		State:   house.State,
		ZipCode: house.ZipCode,
		Price:   house.Price,
	})
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(result)
}

func (h *Handler) GetHouse(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return fiber.ErrBadRequest
	}

	house, err := h.Usecases.GetHouse(c.Context(), id)

	logrus.Infoln(house)

	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(house)
}

func (h *Handler) GetHouses(c *fiber.Ctx) error {
	houses, err := h.Usecases.GetHouses(c.Context())
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(houses)
}

func (h *Handler) UpdateHouse(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return fiber.ErrBadRequest
	}

	house := new(House)
	if err := c.BodyParser(house); err != nil {
		return err
	}

	/// Valida o que tem na declaração da struct
	if err := h.Validator.Struct(house); err != nil {
		logrus.Println(err)

		return &errors.Error{
			Message: "Invalid house data",
			Status:  http.StatusBadRequest,
		}
	}

	result, err := h.Usecases.UpdateHouse(c.Context(), &models.House{
		ID:      id,
		Street:  house.Street,
		Number:  house.Number,
		City:    house.City,
		State:   house.State,
		ZipCode: house.ZipCode,
		Price:   house.Price,
		OwnerID: house.OwnerID,
	})
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(result)
}

func (h *Handler) DeleteHouse(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return fiber.ErrBadRequest
	}

	err = h.Usecases.DeleteHouse(c.Context(), id)
	if err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusNoContent)
}

func (h *Handler) GetHousesByUserID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return fiber.ErrBadRequest
	}

	houses, err := h.Usecases.GetHousesByUserID(c.Context(), id)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(houses)
}
