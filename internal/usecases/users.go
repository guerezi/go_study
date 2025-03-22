package usecases

import (
	"context"
	"imobiliaria/internal/models"
	"imobiliaria/internal/usecases/errors"
)

type Users interface {
	CreateUser(context.Context, *models.User) (*models.User, error)
	GetUser(context.Context, int) (*models.User, error)
}

// CreateUser implements Usecases.
func (u *usecases) CreateUser(ctx context.Context, user *models.User) (*models.User, error) {
	if user.Name == "" {
		return nil, errors.NewError(
			"name should be defined",
			errors.ErrorCodeInvalid,
			nil,
		)
	}

	if user.Age == nil || *user.Age == 0 {
		return nil, errors.NewError(
			"age should be defined",
			errors.ErrorCodeInvalid,
			nil,
		)
	}

	return u.repo.CreateUser(ctx, user)
}

// GetUser implements Usecases.
func (u *usecases) GetUser(ctx context.Context, id int) (*models.User, error) {
	if id == 0 {
		return nil, errors.NewError(
			"id should be defined",
			errors.ErrorCodeInvalid,
			nil,
		)
	}

	return u.repo.GetUser(ctx, id)
}
