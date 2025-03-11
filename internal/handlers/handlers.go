package handlers

import (
	"imobiliaria/internal/usecases"

	"github.com/sirupsen/logrus"
)

// Implements implicito
type Handler struct {
	Usecases usecases.Usecases
	Logger   *logrus.Logger
}

// var _ IHandler = new(Handler)

// type IHandler interface {
// 	GetNothing(c *fiber.Ctx) error
// 	CreateUser(context.Context, any) (any, error)
// 	GetUser(context.Context, string) (any, error)
// }
