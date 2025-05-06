package repositories

import (
	"context"

	"imobiliaria/internal/models"
)

type Houses interface {
	GetHouse(ctx context.Context, id uint) (*models.House, error)
	CreateHouse(ctx context.Context, house *models.House) (*models.House, error)
	GetHouses(ctx context.Context, limit uint, offset uint) ([]*models.House, error)
	UpdateHouse(ctx context.Context, house *models.House) (*models.House, error)
	DeleteHouse(ctx context.Context, id uint) error

	GetHousesByUserID(ctx context.Context, id uint, limit uint, offset uint) ([]*models.House, error)
}
