package repositories

import (
	"context"

	"imobiliaria/internal/models"
)

type Houses interface {
	GetHouse(ctx context.Context, id int) (*models.House, error)
	CreateHouse(ctx context.Context, house *models.House) (*models.House, error)
	GetHouses(ctx context.Context) ([]*models.House, error)
	UpdateHouse(ctx context.Context, house *models.House) (*models.House, error)
	DeleteHouse(ctx context.Context, id int) error

	GetHousesByUserID(ctx context.Context, id int) ([]*models.House, error)
}
