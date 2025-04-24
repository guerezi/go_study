package redis

import (
	repositories "imobiliaria/internal/repositories/cache"

	"github.com/gofiber/storage/redis"
	"github.com/sirupsen/logrus"
)

// TODO: Não deveria estar aqui como eu pensei :(
type Redis struct {
	Storage *redis.Storage
}

type Config struct {
	Host     string
	Port     int
	Password string
	Database int
}

func NewCache(config *Config) (*Redis, error) {
	logrus.Trace("Creating Redis repository")

	storeage := redis.New(redis.Config{
		Host:     config.Host,
		Port:     config.Port,
		Password: config.Password,
		Database: config.Database,
	})

	logrus.Infoln("Created redis storage")

	Redis := &Redis{
		Storage: storeage,
	}

	return Redis, nil
}

var _ repositories.Repositories = new(Redis)
