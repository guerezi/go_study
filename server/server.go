package server

import (
	"errors"

	errorsUsecase "imobiliaria/internal/usecases/errors"
	"imobiliaria/server/handlers"
	errorsHandler "imobiliaria/server/handlers/errors"

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
			var h *errorsHandler.Error
			if errors.As(err, &e) {
				switch e.Code {
				case errorsUsecase.ErrorCodeInvalid:
					return ctx.Status(fiber.StatusBadRequest).SendString(e.Message)
				case errorsUsecase.ErrorCodeNotFound:
					return ctx.Status(fiber.StatusNotFound).SendString(e.Message)
				default:
					// sentry.add(err)
					return ctx.Status(fiber.StatusNotImplemented).SendString("Unknow Error")
				}
			} else if errors.As(err, &h) {
				switch h.Status {
				case fiber.StatusBadRequest:
					return ctx.Status(fiber.StatusBadRequest).SendString(h.Message)
				case fiber.StatusNotFound:
					return ctx.Status(fiber.StatusNotFound).SendString(h.Message)
				case fiber.StatusInternalServerError:
					return ctx.Status(fiber.StatusInternalServerError).SendString(h.Message)
				default:
					// sentry.add(err)
					return ctx.Status(fiber.StatusNotImplemented).SendString("Unknow Error")
				}
			}

			return ctx.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")

		},
	})
	// TODO: add a max body size, maybe 1mb
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

	users := api.Group("/users")

	users.Get("/:id", s.Handler.GetUser)
	users.Post("/", s.Handler.CreateUser)

	houses := api.Group("/houses")

	houses.Get("/", s.Handler.GetHouses)
	houses.Get("/:id", s.Handler.GetHouse)
	houses.Post("/", s.Handler.CreateHouse)
	houses.Put("/:id", s.Handler.UpdateHouse)
	houses.Delete("/:id", s.Handler.DeleteHouse)
	houses.Get("/user/:id", s.Handler.GetHousesByUserID)

	return app.Listen(port)
}
