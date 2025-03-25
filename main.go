package main

import (
	"imobiliaria/internal/repositories"
	"imobiliaria/internal/repositories/mysql"
	"imobiliaria/internal/usecases"
	"imobiliaria/server"
	"imobiliaria/server/handlers"

	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
)

// Deus está triste com tanto copilot

/// air

const DefaultPort = ":3000"

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{PrettyPrint: true})

	// Preicsa ser ponteiro pq to recendo interface (??)
	var m repositories.Repositories = mysql.NewRepository()
	u := usecases.NewUsecases(m)
	v := validator.New()

	h := handlers.Handler{
		Usecases:  u,
		Validator: v,
	}

	/// H com & quer dizer que não é uma copia
	s := &server.Server{
		Handler: &h,
	}

	if err := s.Listen(DefaultPort); err != nil {
		logrus.WithError(err).Fatal("Server error")
	}
}
