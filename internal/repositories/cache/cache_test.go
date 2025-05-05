package cache_test

import (
	"encoding/json"
	"errors"
	"imobiliaria/internal/repositories/cache"
	"imobiliaria/internal/repositories/cache/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TODO: fazer os testes aqui
// Mais sobre go routines
// https://orbstack.dev/

type ModelTest struct {
	ID int `json:"id"`
}

func TestGet(t *testing.T) {
	mockedCache := new(mocks.Cache)

	type Expect struct {
		Data  *ModelTest
		Error error
	}

	tests := []struct {
		Desc   string
		Key    string
		Result Expect
		Mocks  func(*testing.T)
	}{
		{
			Desc: "Invalid Key",
			Key:  "",
			Result: Expect{
				Data:  nil,
				Error: cache.ErrInvalidKey,
			},
			Mocks: func(_ *testing.T) {},
		},
		{
			Desc: "Error Get",
			Key:  "ErrorGet",
			Result: Expect{
				Data:  nil,
				Error: cache.ErrCacheGet,
			},
			Mocks: func(_ *testing.T) {
				mockedCache.On("Get", "ErrorGet").Return(nil, errors.New("I'm a error :D")).Once()
			},
		},
		{
			Desc: "Empty Value",
			Key:  "EmptyValue",
			Result: Expect{
				Data:  nil,
				Error: cache.ErrCacheNotFound,
			},
			Mocks: func(_ *testing.T) {
				mockedCache.On("Get", "EmptyValue").Return([]byte{}, nil).Once()
			},
		},

		{
			Desc: "Error Unmarshall",
			Key:  "ErrorUnmarshall",
			Result: Expect{
				Data:  nil,
				Error: cache.ErrJSONUnmarshal,
			},
			Mocks: func(_ *testing.T) {
				mockedCache.On("Get", "ErrorUnmarshall").Return([]byte("mdc"), nil).Once()
			},
		},
		{
			Desc: "Success",
			Key:  "Success",
			Result: Expect{
				Data: &ModelTest{
					ID: 0,
				},
				Error: nil,
			},
			Mocks: func(t *testing.T) {
				userData, err := json.Marshal(
					&ModelTest{
						ID: 0,
					})
				assert.NoError(t, err, "error should be nil")

				mockedCache.On("Get", "Success").Return(userData, nil).Once()
			},
		},
	}

	for _, test := range tests {
		t.Run(test.Desc, func(t *testing.T) {
			test.Mocks(t)

			result, err := cache.Get[ModelTest](mockedCache, test.Key)

			t.Log(err)
			t.Log(result)

			assert.Equal(t, test.Result.Data, result, "data should be equal")
			assert.ErrorIs(t, err, test.Result.Error, "error should be equal")

			mockedCache.AssertExpectations(t)
		})
	}
}

func TestSet(t *testing.T) {
	mockedCache := new(mocks.Cache)

	type Input struct {
		Key   string
		Value *ModelTest
		Exp   cache.Expiration
	}

	type Expect struct {
		Error error
	}

	tests := []struct {
		Desc   string
		Input  Input
		Result Expect
		Mocks  func(*testing.T)
	}{
		{
			Desc: "Error Marshal",
			Input: Input{
				Key:   "ErrorMarshal",
				Value: nil,
				Exp:   cache.DefaultSetExpiration,
			},
			Result: Expect{
				Error: cache.ErrJSONMarshal,
			},
			Mocks: func(_ *testing.T) {},
		},
		{
			Desc: "Error Set",
			Input: Input{
				Key: "ErrorSet",
				Value: &ModelTest{
					ID: 1,
				},
				Exp: cache.DefaultSetExpiration,
			},
			Result: Expect{
				Error: cache.ErrCacheSet,
			},
			Mocks: func(t *testing.T) {
				value, err := json.Marshal(&ModelTest{ID: 1})
				assert.NoError(t, err, "error should be nil")

				mockedCache.On("Set", "ErrorSet", value, cache.DefaultSetExpiration).Return(errors.New("I'm an error :D")).Once()
			},
		},
		{
			Desc: "Success",
			Input: Input{
				Key: "Success",
				Value: &ModelTest{
					ID: 2,
				},
				Exp: cache.DefaultSetExpiration,
			},
			Result: Expect{
				Error: nil,
			},
			Mocks: func(t *testing.T) {
				value, err := json.Marshal(&ModelTest{ID: 2})
				assert.NoError(t, err, "error should be nil")

				mockedCache.On("Set", "Success", value, cache.DefaultSetExpiration).Return(nil).Once()
			},
		},
	}

	for _, test := range tests {
		t.Run(test.Desc, func(t *testing.T) {
			test.Mocks(t)

			err := cache.Set(mockedCache, test.Input.Key, test.Input.Value, test.Input.Exp)

			t.Log(err)

			assert.ErrorIs(t, err, test.Result.Error, "error should be equal")

			mockedCache.AssertExpectations(t)
		})
	}
}

/// FAzer um teste de usecases
