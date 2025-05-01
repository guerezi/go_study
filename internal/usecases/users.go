package usecases

import (
	"context"

	"imobiliaria/internal/models"
	"imobiliaria/internal/repositories/cache"
	"imobiliaria/internal/usecases/errors"

	"github.com/sirupsen/logrus"
)

type Users interface {
	CreateUser(context.Context, *models.User) (*models.User, error)
	GetUser(context.Context, int) (*models.User, error)
	Login(context.Context, string, string) (*models.User, error)
}

const userCacheKey = "user"

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

	if user, err := cache.Get[models.User](u.cache, cache.BuildKey(userCacheKey, id)); err == nil {
		logrus.Trace("user found in cache")

		return user, nil
	}

	user, err := u.repo.GetUser(ctx, id)
	if err != nil {
		logrus.WithError(err).Error("error getting user from repository")

		return nil, errors.NewError(
			"error getting user from repository",
			errors.ErrorDataBase,
			err,
		)
	}

	if err := cache.Set(u.cache, cache.BuildKey(userCacheKey, id), user, cache.DefaultSetExpiration); err != nil {
		logrus.WithError(err).Error("error setting user in cache")
		// Not returning error here
	}

	return user, nil
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
