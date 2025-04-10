package usecases

import (
	"errors"
	"imobiliaria/internal/repositories"

	"github.com/go-playground/validator/v10"
)

var (
	ErrUserNotFound  = errors.New("no User was founded")
	ErrUsersNotFound = errors.New("no Users were founded")
)

type usecases struct {
	repo repositories.Repositories
	val  *validator.Validate
}

type Usecases interface {
	Users
	Houses
	// Transactions
}

var _ Usecases = new(usecases)

func NewUsecases(repo repositories.Repositories, val *validator.Validate) Usecases {
	return &usecases{repo: repo, val: val}
}
