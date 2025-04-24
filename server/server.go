package server

import (
	"errors"
	"time"

	// "imobiliaria/internal/repositories/cache/redis"
	errorsUsecase "imobiliaria/internal/usecases/errors"
	"imobiliaria/server/handlers"
	errorsHandler "imobiliaria/server/handlers/errors"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/gofiber/fiber/v2/middleware/session"

	"github.com/sirupsen/logrus"
)

type Server struct {
	Handler *handlers.Handler
	// Redis   *redis.Redis
}

func (s *Server) Listen(port string) error {

	sessionStorage := session.New(session.Config{
		Expiration: 24 * time.Hour,
		// Storage:    s.Redis.Storage,
	})

	app := fiber.New(fiber.Config{
		BodyLimit: 1 * 1024 * 1024, // 1MB
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			// NOTE: this mf should be a trace or a debug
			logrus.WithError(err).Infoln("Got an exception")

			// Retrieve the custom status code if it's a *fiber.Error
			var e *errorsUsecase.Error
			var h *errorsHandler.Error
			if errors.As(err, &e) {
				logrus.Info("ERROR: ", e)

				switch e.Code {
				case errorsUsecase.ErrorCodeInvalid:
					return ctx.Status(fiber.StatusBadRequest).SendString(e.Message)
				case errorsUsecase.ErrorCodeNotFound:
					return ctx.Status(fiber.StatusNotFound).SendString(e.Message)
				case errorsUsecase.ErrorDataBase:
					return ctx.Status(fiber.StatusTeapot).SendString(e.Message)
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
	app.Use(requestid.New())
	app.Use(recover.New()) // Esse corno não funcionou >:(
	app.Use(func(c *fiber.Ctx) error {
		logrus.WithFields(logrus.Fields{
			"Method": c.Method(),
			"rotue":  c.Route().Path,
		}).Info(c.Request())

		return c.Next()
	})
	app.Use(compress.New())
	app.Use(limiter.New(limiter.Config{
		Expiration: 10 * time.Second,
	}))
	app.Use(helmet.New())

	// absurdos
	app.Post("/api/login", func(c *fiber.Ctx) error {
		c.Locals("sessionStorage", sessionStorage)

		return s.Handler.Login(c)
	})

	app.Post("/api/register", s.Handler.CreateUser)

	api := app.Group("/api", func(c *fiber.Ctx) error {
		logrus.Infoln("Middleware de segurança :)")

		sess, err := sessionStorage.Get(c)
		if err != nil {
			logrus.WithError(err).Error("Error getting session")

			return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
		}

		username := sess.Get("user") // Email?
		logrus.Info("Username: ", username)

		if username == nil || username == "" {
			return c.Status(fiber.StatusUnauthorized).SendString("Unauthorized")
		}

		return c.Next()
	})

	users := api.Group("/users")

	users.Get("/:id", s.Handler.GetUser)

	houses := api.Group("/houses")

	houses.Get("/", s.Handler.GetHouses)
	houses.Get("/:id", s.Handler.GetHouse)
	houses.Post("/", s.Handler.CreateHouse)
	houses.Put("/:id", s.Handler.UpdateHouse)
	houses.Delete("/:id", s.Handler.DeleteHouse)
	houses.Get("/user/:id", s.Handler.GetHousesByUserID)

	return app.Listen(port)
}
