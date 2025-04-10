package usecases

import (
	"context"
	"imobiliaria/internal/models"
	"imobiliaria/internal/usecases/errors"

	"github.com/sirupsen/logrus"
)

type Houses interface {
	GetHouse(ctx context.Context, id uint) (*models.House, error)
	CreateHouse(ctx context.Context, house *models.House) (*models.House, error)
	GetHouses(ctx context.Context, limit uint, offset uint) ([]*models.House, error)
	UpdateHouse(ctx context.Context, house *models.House) (*models.House, error)
	DeleteHouse(ctx context.Context, id uint) error
	GetHousesByUserID(ctx context.Context, id uint, limit uint, offset uint) ([]*models.House, error)
}

func (u *usecases) GetHouse(ctx context.Context, id uint) (*models.House, error) {
	logrus.Trace("GetHouse usecase called")
	// max int number just because
	if id <= 0 && id >= ^uint(0)>>1 {
		logrus.Trace("Invalid ID", id)

		return nil, errors.NewError(
			"id should be defined",
			errors.ErrorCodeInvalid,
			nil,
		)
	}

	house, err := u.repo.GetHouse(ctx, id)
	if err != nil {
		logrus.WithError(err).Trace("Error getting house")

		return nil, errors.NewError(
			"Error getting house",
			errors.ErrorDataBase,
			err,
		)
	}

	if house == nil {
		logrus.Trace("No Houses were found")

		return nil, errors.NewError(
			"House not found",
			errors.ErrorCodeNotFound,
			nil,
		)
	}

	logrus.Trace("Returnning house", house)

	return house, nil
}

func (u *usecases) CreateHouse(ctx context.Context, house *models.House) (*models.House, error) {
	logrus.Trace("CreateHouse usecase called")
	if house == nil {
		logrus.Trace("Empty house passed", house)

		return nil, errors.NewError(
			"house should be defined",
			errors.ErrorCodeInvalid,
			nil,
		)
	}

	// TODO: colocar isso em outros lugare?
	if err := u.val.Struct(house); err != nil {
		logrus.WithError(err).Trace("Error validating house model")

		return nil, errors.NewError(
			"Error validating house model",
			errors.ErrorCodeInvalid,
			err,
		)
	}

	/// TODO: Colocar mais ErrorDatabase
	createdHouse, err := u.repo.CreateHouse(ctx, house)
	if err != nil {
		logrus.WithError(err).Trace("Error creating house")

		return nil, errors.NewError(
			"Error creating house",
			errors.ErrorDataBase,
			err,
		)
	}
	logrus.Trace("House created", createdHouse)

	return createdHouse, nil
}

func (u *usecases) GetHouses(ctx context.Context, limit uint, offset uint) ([]*models.House, error) {
	// paginates the houses too
	// TODO: Retornar o total de casas para o front, mudar o datatype [] -> {[], &n}
	logrus.Trace("GetHouses usecase called")
	houses, err := u.repo.GetHouses(ctx, limit, offset)
	if err != nil {
		logrus.WithError(err).Trace("Error getting houses")

		return nil, err
	}

	if len(houses) == 0 {
		logrus.Trace("No houses were found")

		return nil, errors.NewError(
			"No houses found",
			errors.ErrorCodeNotFound,
			nil,
		)
	}

	logrus.Trace("Returning houses", houses)

	return houses, nil
}

func (u *usecases) UpdateHouse(ctx context.Context, house *models.House) (*models.House, error) {
	logrus.Trace("UpdateHouse usecase called")
	if house == nil {
		logrus.Trace("Empty house passed", house)

		return nil, errors.NewError(
			"house should be defined",
			errors.ErrorCodeInvalid,
			nil,
		)
	}

	if err := u.val.Struct(house); err != nil {
		logrus.WithError(err).Error("Error validating house model")

		return nil, errors.NewError(
			"Error validating house model",
			errors.ErrorCodeInvalid,
			err,
		)
	}

	updatedHouse, err := u.repo.UpdateHouse(ctx, house)
	if err != nil {
		logrus.WithError(err).Trace("Error updating house")

		return nil, errors.NewError(
			"Error updating house",
			errors.ErrorDataBase,
			err,
		)
	}
	logrus.Trace("House updated", updatedHouse)

	return updatedHouse, nil
}

func (u *usecases) DeleteHouse(ctx context.Context, id uint) error {
	logrus.Trace("DeleteHouse usecase called")

	if id <= 0 && id >= ^uint(0)>>1 {
		logrus.Trace("Invalid ID", id)

		return errors.NewError(
			"id should be defined",
			errors.ErrorCodeInvalid,
			nil,
		)
	}

	err := u.repo.DeleteHouse(ctx, id)
	if err != nil {
		logrus.WithError(err).Trace("Error deleting house")

		return errors.NewError(
			"Error deleting house",
			errors.ErrorDataBase,
			err,
		)
	}

	logrus.Trace("House deleted", id)

	return nil
}

func (u *usecases) GetHousesByUserID(ctx context.Context, id uint, limit uint, offset uint) ([]*models.House, error) {
	logrus.Trace("GetHousesByUserID usecase called")

	if id <= 0 && id >= ^uint(0)>>1 {
		logrus.Trace("Invalid ID", id)

		return nil, errors.NewError(
			"id should be defined",
			errors.ErrorCodeInvalid,
			nil,
		)
	}

	houses, err := u.repo.GetHousesByUserID(ctx, id, limit, offset)
	if err != nil {
		logrus.WithError(err).Trace("Error getting houses by user id")

		return nil, errors.NewError(
			"Error gettings user houses",
			errors.ErrorDataBase,
			err,
		)
	}

	if len(houses) == 0 {
		logrus.Trace("No houses were found")

		return nil, errors.NewError(
			"No houses found",
			errors.ErrorCodeNotFound,
			nil,
		)
	}

	logrus.Trace("Returning houses", houses)

	return houses, nil
}
