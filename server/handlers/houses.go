package handlers

import (
	"fmt"
	"net/http"

	"imobiliaria/internal/models"
	"imobiliaria/server/handlers/errors"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type House struct {
	ID      int     `json:"id"`
	Street  string  `json:"street" validate:"required,street" `
	Number  string  `json:"number" validate:"required,number" `
	City    string  `json:"city" validate:"required,city" `
	State   string  `json:"state" validate:"required,state" `
	ZipCode string  `json:"zip_code" validate:"required,zipcode" `
	Price   float64 `json:"price" validate:"required,gte=0,lte=99999999"`
	OwnerID *int    `json:"owner_id"`
}

func (h *Handler) CreateHouse(c *fiber.Ctx) error {
	logrus.Trace("CreateHouse handler called")

	house := new(House)
	if err := c.BodyParser(house); err != nil {
		logrus.WithError(err).Trace("Error parsing house")

		return err
	}

	if err := h.validator.Validate(house); err != nil {
		logrus.WithError(err).Trace("Error validating house")

		return &errors.Error{
			Message: fmt.Sprintf("Invalid house data: %s", err.Error()),
			Status:  http.StatusBadRequest,
		}
	}

	// TODO: verificar se o owner existe
	owner := 1

	result, err := h.usecases.CreateHouse(c.Context(), &models.House{
		Street:  house.Street,
		Number:  house.Number,
		City:    house.City,
		State:   house.State,
		ZipCode: house.ZipCode,
		Price:   house.Price,
		OwnerID: &owner,
	})
	if err != nil {
		logrus.WithError(err).Trace("Error creating house")

		return err
	}

	return c.Status(fiber.StatusCreated).JSON(result)
}

func (h *Handler) GetHouse(c *fiber.Ctx) error {
	logrus.Trace("GetHouse handler called")

	id, err := c.ParamsInt("id")
	if err != nil || id < 0 {
		logrus.WithError(err).Trace("Invalid ID", id)

		return fiber.ErrBadRequest
	}

	house, err := h.usecases.GetHouse(c.Context(), uint(id))
	if err != nil {
		logrus.WithError(err).Trace("Error getting house")

		return err
	}

	return c.Status(fiber.StatusOK).JSON(house)
}

func (h *Handler) GetHouses(c *fiber.Ctx) error {
	logrus.Trace("GetHouses handler called")

	// gets limit and offset from query
	limit := c.QueryInt("limit", 10)
	offset := c.QueryInt("offset", 0)

	if limit <= 0 || offset < 0 {
		logrus.Trace("Invalid limit or offset", limit, offset)

		return fiber.ErrBadRequest
	}

	houses, err := h.usecases.GetHouses(c.Context(), uint(limit), uint(offset))
	if err != nil {
		logrus.WithError(err).Trace("Error getting houses")

		return err
	}

	return c.Status(fiber.StatusOK).JSON(houses)
}

func (h *Handler) UpdateHouse(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil || id < 0 {
		logrus.WithError(err).Trace("Invalid ID", id)

		return fiber.ErrBadRequest
	}

	house := new(House)
	if err := c.BodyParser(house); err != nil {
		logrus.WithError(err).Trace("Error parsing house")

		return err
	}

	/// Valida o que tem na declaração da struct
	if err := h.validator.Validate(house); err != nil {
		logrus.WithError(err).Trace("Error validating house")

		return &errors.Error{
			Message: "Invalid house data",
			Status:  http.StatusBadRequest,
		}
	}

	result, err := h.usecases.UpdateHouse(c.Context(), &models.House{
		ID:      id,
		Street:  house.Street,
		Number:  house.Number,
		City:    house.City,
		State:   house.State,
		ZipCode: house.ZipCode,
		Price:   house.Price,
		// OwnerID: house.OwnerID,
	})
	if err != nil {
		logrus.WithError(err).Trace("Error updating house")

		return err
	}

	return c.Status(fiber.StatusOK).JSON(result)
}

func (h *Handler) DeleteHouse(c *fiber.Ctx) error {
	logrus.Trace("DeleteHouse handler called")

	id, err := c.ParamsInt("id")
	if err != nil || id < 0 {
		logrus.WithError(err).Trace("Invalid ID", id)

		return fiber.ErrBadRequest
	}

	err = h.usecases.DeleteHouse(c.Context(), uint(id))
	if err != nil {
		logrus.WithError(err).Trace("Error deleting house")

		return err
	}

	return c.SendStatus(fiber.StatusNoContent)
}

func (h *Handler) GetHousesByUserID(c *fiber.Ctx) error {
	logrus.Trace("GetHousesByUserID handler called")

	id, err := c.ParamsInt("id")
	if err != nil || id < 0 {
		logrus.WithError(err).Trace("Invalid ID", id)

		return fiber.ErrBadRequest
	}

	// gets limit and offset from query
	limit := c.QueryInt("limit", 10)
	offset := c.QueryInt("offset", 0)

	if limit <= 0 || offset < 0 {
		logrus.Trace("Invalid limit or offset", limit, offset)

		return fiber.ErrBadRequest
	}

	houses, err := h.usecases.GetHousesByUserID(c.Context(), uint(id), uint(limit), uint(offset))
	if err != nil {
		logrus.WithError(err).Trace("Error getting houses by user id")

		return err
	}

	return c.Status(fiber.StatusOK).JSON(houses)
}
