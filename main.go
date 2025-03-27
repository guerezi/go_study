package main

import (
	"context"
	"imobiliaria/internal/repositories/mysql"
	"imobiliaria/internal/usecases"
	"imobiliaria/server"
	"imobiliaria/server/handlers"

	"github.com/go-playground/validator/v10"
	"github.com/sethvargo/go-envconfig"
	"github.com/sirupsen/logrus"
)

// Deus está triste com tanto copilot

/// air

const DefaultPort = ":3000"

type Config struct {
	Host     string `env:"DB_HOST" default:"localhost"`
	Port     string `env:"DB_PORT" default:"3306"`
	User     string `env:"DB_USER" default:"root"`
	Password string `env:"DB_PASSWORD" default:"password"`
	Database string `env:"DB_NAME" default:"database"`
}

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{PrettyPrint: true})

	ctx := context.Background()

	// Configurações do banco de dados
	var c Config
	if err := envconfig.Process(ctx, &c); err != nil {
		panic(err)
	}

	// obg chat gpt
	// Preicsa ser ponteiro pq to recendo interface (??)
	m, err := mysql.NewRepository(&mysql.Config{
		Host:     c.Host,
		Port:     c.Port,
		User:     c.User,
		Password: c.Password,
		Database: c.Database,
	})

	if err != nil {
		logrus.WithError(err).Fatal("Error creating repository")
	}

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
