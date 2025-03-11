package main

import (
	"imobiliaria/internal/handlers"
	"imobiliaria/internal/repositories"
	"imobiliaria/internal/repositories/memory"
	"imobiliaria/internal/usecases"
	"imobiliaria/pkg/server"

	"github.com/sirupsen/logrus"
)

// Deus está triste com tanto copilot

/// air

const DefaultPort = ":3000"

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{PrettyPrint: true})

	// Preicsa ser ponteiro pq to recendo interface (??)
	var m repositories.Repositories = memory.NewMemory()
	var u usecases.Usecases = usecases.NewUsecases(m)

	var h handlers.Handler = handlers.Handler{
		Usecases: u,
	}

	/// H com & quer dizer que não é uma copia
	s := &server.Server{
		Handler: &h,
	}

	if err := s.Listen(DefaultPort); err != nil {
		logrus.WithError(err).Fatal("Server error")
	}
}
