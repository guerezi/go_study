package server

import (
	"errors"
	errorsUsecase "imobiliaria/internal/usecases/errors"
	"imobiliaria/server/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/sirupsen/logrus"
)

type Server struct {
	Handler *handlers.Handler
}

func (s *Server) Listen(port string) error {
	app := fiber.New(fiber.Config{
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			// NOTE: this mf should be a trace or a debug
			logrus.WithError(err).Infoln("Got an exception")

			// Retrieve the custom status code if it's a *fiber.Error
			var e *errorsUsecase.Error
			if !errors.As(err, &e) {
				// grafa.add(err)
				return ctx.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
			}

			switch e.Code {
			case errorsUsecase.ErrorCodeInvalid:
				return ctx.Status(fiber.StatusBadRequest).SendString(e.Message)
			case errorsUsecase.ErrorCodeNotFound:
				return ctx.Status(fiber.StatusNotFound).SendString(e.Message)
			default:
				// sentry.add(err)
				return ctx.Status(fiber.StatusNotImplemented).SendString("Unknow Error")
			}
		},
	})
	app.Use(requestid.New())
	app.Use(recover.New()) // Esse corno não funcionou >:(
	app.Use(func(c *fiber.Ctx) error {
		logrus.WithFields(logrus.Fields{
			"Method": c.Method(),
			"rotue":  c.Route().Path,
		}).Info(c.Request())

		return c.Next()
	})
	app.Use(helmet.New())

	api := app.Group("/api")

	api.Get("/users/:id", s.Handler.GetUser)
	api.Post("/users", s.Handler.CreateUser)

	return app.Listen(port)
}
