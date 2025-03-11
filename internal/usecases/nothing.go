package usecases

import (
	"context"
	"errors"

	"github.com/sirupsen/logrus"
)

type Nothings interface {
	GetNothing(context.Context) (bool, error)
}

// GetNothing implements IUsecases.
func (u *usecases) GetNothing(c context.Context) (bool, error) {
	logrus.Println("Got nothing")

	return true, errors.New("errei, fui garoto")
}
