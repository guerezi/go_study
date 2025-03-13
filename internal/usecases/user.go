package usecases

import (
	"context"

	"imobiliaria/internal/models"
)

type Users interface {
	CreateUser(context.Context, *models.User) (*models.User, error)
	GetUser(context.Context, int) (*models.User, error)
}

// CreateUser implements Usecases.
func (u *usecases) CreateUser(ctx context.Context, user *models.User) (*models.User, error) {
	return u.repo.CreateUser(ctx, user)
}

// GetUser implements Usecases.
func (u *usecases) GetUser(ctx context.Context, id int) (*models.User, error) {
	return u.repo.GetUser(ctx, id)
}
