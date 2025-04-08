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

	GetHousesByUserID(ctx context.Context, id uint) ([]*models.House, error)
}

func (u *usecases) GetHouse(ctx context.Context, id uint) (*models.House, error) {
	// max int number just because
	if id <= 0 && id >= uint(^uint(0)>>1) {
		return nil, errors.NewError(
			"id should be defined",
			errors.ErrorCodeInvalid,
			nil,
		)
	}

	house, err := u.repo.GetHouse(ctx, id)
	if err != nil {
		return nil, errors.NewError(
			"Error getting house",
			errors.ErrorDataBase,
			err,
		)
	}

	if house == nil {
		return nil, errors.NewError(
			"House not found",
			errors.ErrorCodeNotFound,
			nil,
		)
	}

	return house, nil
}

func (u *usecases) CreateHouse(ctx context.Context, house *models.House) (*models.House, error) {
	// VALIDAR AQUI DE NOVO?
	// TODO: Adicionar mais logrus trace nos usecases
	if house == nil {
		return nil, errors.NewError(
			"house should be defined",
			errors.ErrorCodeInvalid,
			nil,
		)
	}

	// TODO: colocar isso em outros lugares
	if err := u.val.Struct(house); err != nil {
		logrus.WithError(err).Error("Error validating house model")

		return nil, errors.NewError(
			"Error validating house model",
			errors.ErrorCodeInvalid,
			err,
		)
	}

	/// TODO: Colocar mais ErrorDatabase
	createdHouse, err := u.repo.CreateHouse(ctx, house)
	if err != nil {
		return nil, errors.NewError(
			"Error creating house",
			errors.ErrorDataBase,
			err,
		)
	}

	return createdHouse, nil
}

func (u *usecases) GetHouses(ctx context.Context, limit uint, offset uint) ([]*models.House, error) {
	// paginates the houses too
	// TODO: Retornar o total de casas para o front, mudar o datatype [] -> {[], &n}
	houses, err := u.repo.GetHouses(ctx, limit, offset)
	if err != nil {
		return nil, err
	}

	if len(houses) == 0 {
		return nil, errors.NewError(
			"No houses found",
			errors.ErrorCodeNotFound,
			nil,
		)
	}

	return houses, nil
}

func (u *usecases) UpdateHouse(ctx context.Context, house *models.House) (*models.House, error) {
	if house == nil {
		return nil, errors.NewError(
			"house should be defined",
			errors.ErrorCodeInvalid,
			nil,
		)
	}

	return u.repo.UpdateHouse(ctx, house)
}

func (u *usecases) DeleteHouse(ctx context.Context, id uint) error {
	if id <= 0 && id >= uint(^uint(0)>>1) {
		return errors.NewError(
			"id should be defined",
			errors.ErrorCodeInvalid,
			nil,
		)
	}

	return u.repo.DeleteHouse(ctx, id)
}

func (u *usecases) GetHousesByUserID(ctx context.Context, id uint) ([]*models.House, error) {
	if id <= 0 && id >= uint(^uint(0)>>1) {
		return nil, errors.NewError(
			"id should be defined",
			errors.ErrorCodeInvalid,
			nil,
		)
	}

	houses, err := u.repo.GetHousesByUserID(ctx, id)
	if err != nil {
		return nil, err
	}

	if len(houses) == 0 {
		return nil, errors.NewError(
			"No houses found",
			errors.ErrorCodeNotFound,
			nil,
		)
	}

	return houses, nil
}
