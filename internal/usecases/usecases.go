package usecases

import (
	"errors"
	"imobiliaria/internal/repositories"
)

var ErrUserNotFound error = errors.New("no User was founded")
var ErrUsersNotFound error = errors.New("no Users were founded")

type usecases struct {
	repo repositories.Repositories
}

type Usecases interface {
	Users
	Nothings
	// Banks
	// Transactions
}

var _ Usecases = new(usecases)

func NewUsecases(repo repositories.Repositories) Usecases {
	return &usecases{repo: repo}
}
