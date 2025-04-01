package usecases

import (
	"context"
	"imobiliaria/internal/models"
	"imobiliaria/internal/usecases/errors"
)

type Houses interface {
	GetHouse(ctx context.Context, id int) (*models.House, error)
	CreateHouse(ctx context.Context, house *models.House) (*models.House, error)
	GetHouses(ctx context.Context) ([]*models.House, error)
	UpdateHouse(ctx context.Context, house *models.House) (*models.House, error)
	DeleteHouse(ctx context.Context, id int) error

	GetHousesByUserID(ctx context.Context, id int) ([]*models.House, error)
}

func (u *usecases) GetHouse(ctx context.Context, id int) (*models.House, error) {
	if id <= 0 {
		return nil, errors.NewError(
			"id should be defined",
			errors.ErrorCodeInvalid,
			nil,
		)
	}

	// HTTP 500 é uma escolha de design
	return u.repo.GetHouse(ctx, id)
}

func (u *usecases) CreateHouse(ctx context.Context, house *models.House) (*models.House, error) {
	// VALIDAR AQUI DE NOVO?
	if house == nil {
		return nil, errors.NewError(
			"house should be defined",
			errors.ErrorCodeInvalid,
			nil,
		)
	}

	return u.repo.CreateHouse(ctx, house)
}

func (u *usecases) GetHouses(ctx context.Context) ([]*models.House, error) {
	houses, err := u.repo.GetHouses(ctx)
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

func (u *usecases) DeleteHouse(ctx context.Context, id int) error {
	if id <= 0 {
		return errors.NewError(
			"id should be defined",
			errors.ErrorCodeInvalid,
			nil,
		)
	}

	return u.repo.DeleteHouse(ctx, id)
}

func (u *usecases) GetHousesByUserID(ctx context.Context, id int) ([]*models.House, error) {
	if id <= 0 {
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
