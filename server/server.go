package server

import (
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
	// TODO(guerezi): Fazer -> https://docs.gofiber.io/guide/error-handling#custom-error-handler

	app := fiber.New()
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
