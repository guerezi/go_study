// ler : https://docs.gofiber.io/
// TODO(guerezi): trocar mux pra fiber
// 		Entender como mudar o w e r por um context de fiber
// TODO(guerezi): separar em packages

package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

type Response struct {
	Body string `json:"body"`
}

/// air

func main() {
	var u Usecases = &AllUsecases{}
	l := log.New()
	l.SetFormatter(&log.JSONFormatter{PrettyPrint: true})

	var h IHandler = &Handler{
		Usecases: u,
		Logger:   l,
	}

	// TODO(guerezi): Fazer -> https://docs.gofiber.io/guide/error-handling#custom-error-handler
	app := fiber.New()
	// Logging Request ID
	app.Use(requestid.New())
	app.Use(func(c *fiber.Ctx) error {
		l.WithFields(log.Fields{
			"req":    "aaaa",
			"Method": c.Method(),
			"rotue":  c.Route().Path,
		}).Info(c.Request())
		return nil
	})
	app.Use(helmet.New())
	app.Use(recover.New())

	api := app.Group("/api") // func(c *fiber.Ctx) error {
	// 	log.Println("Handler do group")
	// 	return nil
	// })

	api.Get("/", h.GetNothing)

	// http.ListenAndServe(":3000", app)
	app.Listen(":3000")
}

// Arquivo HTTP

// Implements implicito
type Handler struct {
	Usecases Usecases
	Logger   *log.Logger
}

type IHandler interface {
	GetNothing(c *fiber.Ctx) error
}

type ErrorHandler struct {
	Message string
	Status  int
}

// Error implements error.
func (e ErrorHandler) Error() string {
	return fmt.Sprintf("Error %d: %s ", e.Status, e.Message)
}

var _ error = new(ErrorHandler)

var _ IHandler = new(Handler)

func (h *Handler) GetNothing(c *fiber.Ctx) error {
	c.Response().Header.Add("Content-Type", "enconding/json")
	entry := h.Logger.WithFields(log.Fields{
		"reqId": c.GetReqHeaders()["X-Request-Id"],
	})

	entry.Infoln("REQUEST ID:", c.Get(fiber.HeaderXRequestedWith))
	entry.WithContext(c.Context()).Info("ctx")

	// ctx := context.Background()
	// ctx, cancel := context.WithCancel(context.Background())
	// ctx, _ := context.WithDeadline(context.Background(), time.Now().Add(5*time.Second))
	// ctx := context.WithValue(context.Background(), "chaves", "valor")
	// defer cancel()

	go func() {
		for {
			log.Println("Deus é top")
			time.Sleep(1 * time.Second)

			/// Aqui é um await
			// _, b := <-ctx.Done()
			// log.Println("Ctx was called")

			// value := ctx.Value("chaves")

			return
		}
	}()

	if _, err := h.Usecases.GetNothing(c.Context()); err != nil {
		// json.NewEncoder(c.Response().BodyWriter()).Encode({err})
		// return errors.Join(fiber.ErrBadGateway, err)
	}

	// err := ErrorHandler{
	// 	Message: "Deu ruim",
	// 	Status:  http.StatusEarlyHints,
	// }

	// e := errors.Join(fiber.ErrBadGateway, fiber.ErrConflict, fiber.ErrBadRequest)
	// anotherE := errors.Unwrap(e)
	aaaanother := fmt.Errorf("errei_mds %w", fiber.ErrForbidden)
	anotherE := errors.Unwrap(aaaanother)

	// errors.Is()

	// errors.As()

	h.Logger.WithError(aaaanother).WithFields(log.Fields{
		"chaves":  "chaves",
		"todos":   "atentos",
		"olhando": "pra tv",
	}).Errorln("Amém")

	return anotherE

	// return fiber.ErrBadRequest
	// return c.Status(201).SendString("Hello, World!")

	// json.NewEncoder(c.Response().BodyWriter()).Encode(err)
	// return err
}

// Pacote usecases

// arquivo
type NothingUsecases struct{}

// GetNothing implements IUsecases.
func (u *NothingUsecases) GetNothing(c context.Context) (bool, error) {
	log.Println("Got nothing")

	return true, errors.New("Errei, fui garoto")
}

// arquivo
type UserUsecases struct{}

var ErrUserNotFound error = errors.New("No User was founded")
var ErrUsersNotFound error = errors.New("No Users were founded")

func (u *UserUsecases) GetUsers() (any, error) {
	return nil, ErrUsersNotFound
}

type Usecases interface {
	GetNothing(context.Context) (bool, error)
	GetUsers() (any, error)
}

type AllUsecases struct {
	NothingUsecases
	UserUsecases
}

var _ Usecases = new(AllUsecases)
