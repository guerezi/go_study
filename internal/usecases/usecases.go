package usecases

import (
	"errors"
	"imobiliaria/internal/repositories/cache/redis"
	repositories "imobiliaria/internal/repositories/database"

	"github.com/go-playground/validator/v10"
)

var (
	ErrUserNotFound  = errors.New("no User was founded")
	ErrUsersNotFound = errors.New("no Users were founded")
)

type usecases struct {
	repo  repositories.Repositories
	cache *redis.Redis
	val   *validator.Validate
}

type Usecases interface {
	Users
	Houses
	// Transactions
}

var _ Usecases = new(usecases)

func NewUsecases(repo repositories.Repositories, val *validator.Validate, cache *redis.Redis) Usecases {
	return &usecases{repo: repo, val: val, cache: cache}
}
