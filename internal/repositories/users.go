package repositories

import (
	"context"

	"imobiliaria/internal/models"
)

type Users interface {
	CreateUser(context.Context, *models.User) (*models.User, error)
	GetUser(context.Context, int) (*models.User, error)
}
