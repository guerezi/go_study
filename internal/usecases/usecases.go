package usecases

import (
	"errors"

	"imobiliaria/internal/repositories"
)

var (
	ErrUserNotFound  error = errors.New("no User was founded")
	ErrUsersNotFound error = errors.New("no Users were founded")
)

type usecases struct {
	repo repositories.Repositories
}

type Usecases interface {
	Users
	// Banks
	// Transactions
}

var _ Usecases = new(usecases)

func NewUsecases(repo repositories.Repositories) Usecases {
	return &usecases{repo: repo}
}
