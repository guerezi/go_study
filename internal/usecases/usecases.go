package usecases

import (
	"errors"

	"imobiliaria/internal/repositories/cache"
	repositories "imobiliaria/internal/repositories/database"
	"imobiliaria/internal/validator"
)

var (
	ErrUserNotFound  = errors.New("no User was founded")
	ErrUsersNotFound = errors.New("no Users were founded")
)

type usecases struct {
	repo      repositories.Repositories
	cache     cache.Cache
	validator *validator.Validator
}

type Usecases interface {
	Users
	Houses
	// Transactions
}

var _ Usecases = new(usecases)

func NewUsecases(repo repositories.Repositories, val *validator.Validator, cache cache.Cache) Usecases {
	return &usecases{repo: repo, validator: val, cache: cache}
}
