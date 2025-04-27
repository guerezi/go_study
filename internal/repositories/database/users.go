package repositories

import (
	"context"
	"imobiliaria/internal/models"
)

type Users interface {
	CreateUser(ctx context.Context, user *models.User) (*models.User, error)
	GetUser(ctx context.Context, id int) (*models.User, error)
	Login(ctx context.Context, email string, password string) (*models.User, error)
}
