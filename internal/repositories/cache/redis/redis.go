package redis

import (
	"time"

	"imobiliaria/internal/repositories/cache"

	"github.com/gofiber/storage/redis"
	"github.com/sirupsen/logrus"
)

type Redis struct {
	Storage *redis.Storage
}

type Config struct {
	Host     string
	Port     int
	Password string
	Database int
}

func NewCache(config *Config) (cache.Cache, error) {
	logrus.Trace("Creating Redis repository")

	storage := redis.New(redis.Config{
		Host:     config.Host,
		Port:     config.Port,
		Password: config.Password,
		Database: config.Database,
	})

	logrus.Trace("Created redis storage")

	Redis := &Redis{
		Storage: storage,
	}

	return Redis, nil
}

// Delete implements repositories.Cache.
func (r *Redis) Delete(key string) error {
	return r.Storage.Delete(key)
}

// Get implements repositories.Cache.
func (r *Redis) Get(key string) ([]byte, error) {
	return r.Storage.Get(key)
}

// Set implements repositories.Cache.
func (r *Redis) Set(key string, value []byte, exp cache.Expiration) error {
	return r.Storage.Set(key, value, time.Duration(exp))
}

// GetStorage implements cache.Cache.
func (r *Redis) GetStorage() interface{} {
	return r.Storage
}

var _ cache.Cache = (*Redis)(nil)
