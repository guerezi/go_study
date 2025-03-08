// ler : https://docs.gofiber.io/
// TODO(guerezi): trocar mux pra fiber
// 		Entender como mudar o w e r por um context de fiber
// TODO(guerezi): separar em packages

package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
)

type Response struct {
	Body string `json:"body"`
}

/// air

func main() {
	var u Usecases = &AllUsecases{}
	var h IHandler = &Handler{
		Usecases: u,
	}

	app := fiber.New()

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
	log.Println("Get nothing")

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

	if ok := h.Usecases.GetNothing(); ok {
		// json.NewEncoder(c.Response().BodyWriter()).Encode({err})
	}

	err := ErrorHandler{
		Message: "Deu ruim",
		Status:  http.StatusEarlyHints,
	}

	log.Println(err)

	c.SendStatus(418)
	return c.SendString("Hello, World!")

	// json.NewEncoder(c.Response().BodyWriter()).Encode(err)
	// return err
}

// Pacote usecases

// arquivo
type NothingUsecases struct{}

// GetNothing implements IUsecases.
func (u *NothingUsecases) GetNothing() bool {
	log.Println("Got nothing")
	return true
}

// arquivo
type UserUsecases struct{}

var ErrUserNotFound error = errors.New("No User was founded")
var ErrUsersNotFound error = errors.New("No Users were founded")

func (u *UserUsecases) GetUsers() (any, error) {
	return nil, ErrUsersNotFound
}

type Usecases interface {
	GetNothing() bool
	GetUsers() (any, error)
}

type AllUsecases struct {
	NothingUsecases
	UserUsecases
}

var _ Usecases = new(AllUsecases)
