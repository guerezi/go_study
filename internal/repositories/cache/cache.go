package cache

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
)

//go:generate mockery --name=Cache --output=./mocks --filename=cache.go
type Cache interface {
	Get(key string) ([]byte, error)
	Set(key string, value []byte, exp Expiration) error
	Delete(key string) error
}

type Expiration time.Duration

const DefaultSetExpiration Expiration = Expiration(1 * time.Minute)

var (
	ErrJSONMarshal   = fmt.Errorf("error marshalling value to Cache")
	ErrJSONUnmarshal = fmt.Errorf("error unmarshalling value from Cache")
	ErrCacheSet      = fmt.Errorf("error setting key in Cache")
	ErrCacheGet      = fmt.Errorf("error getting key from Cache")
	ErrInvalidKey    = fmt.Errorf("error Key is not valid")
	ErrCacheNotFound = fmt.Errorf("error Value not found in Cache")
)

func BuildKey(prefix string, model any) string {
	return fmt.Sprintf("%s:%s", prefix, fmt.Sprintf("%v", model))
}

func Get[T any](r Cache, key string) (*T, error) {
	if key == "" {
		return nil, ErrInvalidKey
	}

	v := new(T)

	logrus.Tracef("Getting key %s from redis", key)

	value, err := r.Get(key)
	if err != nil {
		return nil, errors.Join(ErrCacheGet, err)
	}

	if len(value) == 0 {
		logrus.Tracef("Key %s not found in redis", key)

		return nil, ErrCacheNotFound
	}

	if err := json.Unmarshal(value, v); err != nil {
		logrus.WithError(err).Tracef("Error unmarshalling value from cache: %s", string(value))

		return nil, errors.Join(ErrJSONUnmarshal, err)
	}

	logrus.Tracef("Got key %s from redis", key)

	return v, nil
}

func Set[T any](r Cache, key string, value *T, exp Expiration) error {
	logrus.Trace("Setting key in redis")

	v, err := json.Marshal(value)
	if err != nil {
		return errors.Join(ErrJSONMarshal, err)
	}

	if err := r.Set(key, v, exp); err != nil {
		return errors.Join(ErrCacheSet, err)
	}

	logrus.Tracef("Key %s set in redis", key)

	return nil
}

func Delete(r Cache, key string) error {
	logrus.Tracef("Deleting key %s from redis", key)

	if err := r.Delete(key); err != nil {
		return errors.Join(ErrCacheGet, err)
	}

	logrus.Tracef("Deleted key %s from redis", key)

	return nil
}
