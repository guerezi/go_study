package handlers

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func (h *Handler) GetNothing(c *fiber.Ctx) error {
	c.Response().Header.Add("Content-Type", "enconding/json")
	entry := h.Logger.WithFields(logrus.Fields{
		"reqId": c.GetReqHeaders()["X-Request-Id"],
	})

	entry.Infoln("REQUEST ID:", c.Get(fiber.HeaderXRequestedWith))
	entry.WithContext(c.Context()).Info("ctx")

	// ctx := context.Background()
	ctx, cancel := context.WithCancel(context.Background())
	// ctx, _ := context.WithDeadline(context.Background(), time.Now().Add(5*time.Second))
	// ctx := context.WithValue(context.Background(), "chaves", "valor")
	defer cancel()

	go func() {
		for {
			// log.Println("Deus é top")

			select {
			case <-ctx.Done():
				log.Println("Ctx was called")
				return
			default:
				log.Println("Default")
			}

			// time.Sleep(10 * time.Millisecond)

			/// Aqui é um await
			// _, b := <-ctx.Done()
			// log.Println("Ctx was called")

			// value := ctx.Value("chaves")

			// return
		}
	}()

	time.Sleep(1 * time.Second)

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

	h.Logger.WithError(aaaanother).WithFields(logrus.Fields{
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
