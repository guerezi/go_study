package usecases

import (
	"context"
	"fmt"

	"imobiliaria/internal/models"
)

type Users interface {
	CreateUser(context.Context, *models.User) (*models.User, error)
	GetUser(context.Context, int) (*models.User, error)
}

// CreateUser implements Usecases.
func (u *usecases) CreateUser(ctx context.Context, user *models.User) (*models.User, error) {
	if user.Name == "" {
		return nil, fmt.Errorf("name should be defined")
	}

	if user.Age == nil || *user.Age == 0 {
		return nil, fmt.Errorf("age should be defined")
	}

	return u.repo.CreateUser(ctx, user)
}

// GetUser implements Usecases.
func (u *usecases) GetUser(ctx context.Context, id int) (*models.User, error) {
	if id == 0 {
		return nil, fmt.Errorf("id should be defined")
	}

	return u.repo.GetUser(ctx, id)
}
