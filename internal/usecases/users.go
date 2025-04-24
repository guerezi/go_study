package usecases

import (
	"context"

	"imobiliaria/internal/models"
	"imobiliaria/internal/usecases/errors"

	"github.com/sirupsen/logrus"
)

type Users interface {
	CreateUser(context.Context, *models.User) (*models.User, error)
	GetUser(context.Context, int) (*models.User, error)
	Login(context.Context, string, string) (*models.User, error)
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

func (u *usecases) Login(ctx context.Context, email string, password string) (*models.User, error) {
	// TODO: validator for email and password
	if email == "" {
		logrus.Trace("email is empty at login")

		return nil, errors.NewError(
			"email should be defined",
			errors.ErrorCodeInvalid,
			nil,
		)
	}

	if password == "" {
		logrus.Trace("password is empty at login")

		return nil, errors.NewError(
			"password should be defined",
			errors.ErrorCodeInvalid,
			nil,
		)
	}

	return u.repo.Login(ctx, email, password)
}
