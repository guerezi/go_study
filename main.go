package main

import (
	"context"
	"strconv"

	"imobiliaria/internal/repositories/cache/redis"
	"imobiliaria/internal/repositories/database/mysql"
	"imobiliaria/internal/usecases"
	"imobiliaria/internal/validator"
	"imobiliaria/server"
	"imobiliaria/server/handlers"

	"github.com/sethvargo/go-envconfig"
	"github.com/sirupsen/logrus"
)

// Deus está triste com tanto copilot

/// air

const DefaultPort = ":3000"

type Config struct {
	Host          string `env:"DB_HOST" default:"localhost"`
	Port          string `env:"DB_PORT" default:"3306"`
	User          string `env:"DB_USER" default:"root"`
	Password      string `env:"DB_PASSWORD" default:"password"`
	Database      string `env:"DB_NAME" default:"database"`
	RedisHost     string `env:"REDIS_HOST" default:"localhost"`
	RedisPort     string `env:"REDIS_PORT" default:"6379"`
	RedisPassword string `env:"REDIS_PASSWORD" default:"password"`
	RedisDatabase string `env:"REDIS_DB" default:"0"`
}

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{PrettyPrint: true})
	logrus.SetLevel(logrus.TraceLevel)

	ctx := context.Background()

	// Configurações do banco de dados
	var c Config
	if err := envconfig.Process(ctx, &c); err != nil {
		panic(err)
	}

	redisPort := func() int {
		port, err := strconv.Atoi(c.RedisPort)
		if err != nil {
			logrus.WithError(err).Fatal("Invalid Redis port")
		}
		return port
	}()

	// TODO: Deveria estar dentro do newRepository?
	r, err := redis.NewCache(&redis.Config{
		Host:     c.RedisHost,
		Port:     redisPort,
		Password: c.RedisPassword,
		Database: 0, // redisDatabase,
	})
	if err != nil {
		logrus.WithError(err).Fatal("Error creating redis repository")
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

	v := validator.NewValidator()

	u := usecases.NewUsecases(m, v, r)
	h := handlers.NewHandler(u, v)

	/// H com & quer dizer que não é uma copia
	s := &server.Server{
		Handler: h,
		Cache:   r,
	}

	if err := s.Listen(DefaultPort); err != nil {
		logrus.WithError(err).Fatal("Server error")
	}
}
