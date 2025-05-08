package usecases

import (
	"context"
	"encoding/json"
	e "errors"
	"testing"

	"imobiliaria/internal/models"
	cc "imobiliaria/internal/repositories/cache"
	ch "imobiliaria/internal/repositories/cache/mocks"
	db "imobiliaria/internal/repositories/database/mocks"
	"imobiliaria/internal/usecases/errors"
	"imobiliaria/internal/validator"

	"github.com/stretchr/testify/assert"
)

func TestGetHouse(t *testing.T) {
	repo := &db.Repositories{}
	cache := &ch.Cache{}

	usecases := NewUsecases(repo, nil, cache)

	ctx := context.Background()

	type Expect struct {
		House *models.House
		Error error
	}

	tests := []struct {
		Name   string
		Result Expect
		Params uint
		Mocks  func()
	}{
		{
			Name: "Cache",
			Result: Expect{
				House: &models.House{
					ID:      1,
					Street:  "Street",
					Number:  "123",
					City:    "City",
					State:   "SC",
					ZipCode: "12345-678",
					Price:   100000,
				},
				Error: nil,
			},
			Params: 1,
			Mocks: func() {
				houseData, err := json.Marshal(&models.House{
					ID:      1,
					Street:  "Street",
					Number:  "123",
					City:    "City",
					State:   "SC",
					ZipCode: "12345-678",
					Price:   100000,
				})

				assert.NoError(t, err, "error should be nil")

				cache.On("Get", "house:1").Return(houseData, nil).Once()
			},
		},
		{
			Name: "DatabaseError",
			Result: Expect{
				House: nil,
				Error: errors.NewError("Error getting house", errors.ErrorDataBase, assert.AnError),
			},
			Params: 4,
			Mocks: func() {
				cache.On("Get", "house:4").Return(nil, nil).Once()
				repo.On("GetHouse", ctx, uint(4)).Return(nil, assert.AnError).Once()
			},
		},

		{
			Name: "NotFound",
			Result: Expect{
				House: nil,
				Error: errors.NewError("House not found", errors.ErrorCodeNotFound, nil),
			},
			Params: 3,
			Mocks: func() {
				cache.On("Get", "house:3").Return(nil, nil).Once()
				repo.On("GetHouse", ctx, uint(3)).Return(nil, nil).Once()
			},
		},
		{
			Name: "Database",
			Result: Expect{
				House: &models.House{
					ID:      2,
					Street:  "Street",
					Number:  "123",
					City:    "City",
					State:   "SC",
					ZipCode: "12345-678",
					Price:   100000,
				},
				Error: nil,
			},
			Params: 2,
			Mocks: func() {
				cache.On("Get", "house:2").Return(nil, e.New("Some cache error")).Once()

				repo.On("GetHouse", ctx, uint(2)).Return(&models.House{
					ID:      2,
					Street:  "Street",
					Number:  "123",
					City:    "City",
					State:   "SC",
					ZipCode: "12345-678",
					Price:   100000,
				}, nil).Once()

				cahceData, err := json.Marshal(&models.House{
					ID:      2,
					Street:  "Street",
					Number:  "123",
					City:    "City",
					State:   "SC",
					ZipCode: "12345-678",
					Price:   100000,
				})

				assert.NoError(t, err, "error should be nil")

				cache.On("Set", "house:2", cahceData, cc.DefaultSetExpiration).Return(e.New("Some other cache error that should not show")).Once()
			},
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			test.Mocks()
			result, err := usecases.GetHouse(ctx, test.Params)

			t.Log(result)

			assert.Equal(t, test.Result.House, result, "house should be equal")
			assert.Equal(t, test.Result.Error, err, "error should be equal")
		})
	}
}

func TestCreateHouse(t *testing.T) {
	repo := new(db.Repositories)
	cache := new(ch.Cache)
	v := validator.NewValidator()

	usecases := NewUsecases(repo, v, cache)

	ctx := context.Background()

	type Expect struct {
		House *models.House
		Error error
	}

	tests := []struct {
		Name   string
		Result Expect
		Params *models.House
		Mocks  func()
	}{
		{
			Name: "InvalidHouseData",
			Result: Expect{
				House: nil,
				Error: errors.NewError("house should be defined", errors.ErrorCodeInvalid, nil),
			},
			Params: nil,
			Mocks:  func() {},
		},
		{
			Name: "ValidationErr",
			Result: Expect{
				House: nil,
				Error: errors.NewError("Error validating house model", errors.ErrorCodeInvalid, validator.ErrValidation),
			},
			Params: &models.House{
				Street:  "Street",
				Number:  "123",
				City:    "City",
				State:   "SC",
				ZipCode: "wrong-zip",
				Price:   100000,
			},
			Mocks: func() {},
		},
		{
			Name: "CreationErr",
			Result: Expect{
				House: nil,
				Error: errors.NewError("Error creating house", errors.ErrorDataBase, e.New("some database err")),
			},
			Params: &models.House{
				ID:      1,
				Street:  "Street",
				Number:  "123",
				City:    "City",
				State:   "SC",
				ZipCode: "12345-678",
				Price:   100000,
			},
			Mocks: func() {
				repo.On("CreateHouse", ctx, &models.House{
					ID:      1,
					Street:  "Street",
					Number:  "123",
					City:    "City",
					State:   "SC",
					ZipCode: "12345-678",
					Price:   100000,
				}).Return(nil, e.New("some database err")).Once()
			},
		},
		{
			Name: "Success",
			Result: Expect{
				House: &models.House{
					ID:      1,
					Street:  "Street",
					Number:  "123",
					City:    "City",
					State:   "SC",
					ZipCode: "12345-678",
					Price:   100000,
				},
				Error: nil,
			},
			Params: &models.House{
				Street:  "Street",
				Number:  "123",
				City:    "City",
				State:   "SC",
				ZipCode: "12345-678",
				Price:   100000,
			},
			Mocks: func() {
				repo.On("CreateHouse", ctx, &models.House{
					Street:  "Street",
					Number:  "123",
					City:    "City",
					State:   "SC",
					ZipCode: "12345-678",
					Price:   100000,
				}).Return(&models.House{
					ID:      1,
					Street:  "Street",
					Number:  "123",
					City:    "City",
					State:   "SC",
					ZipCode: "12345-678",
					Price:   100000,
				}, nil).Once()
			},
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			test.Mocks()
			result, err := usecases.CreateHouse(ctx, test.Params)

			t.Log(result)
			t.Log(err)

			assert.Equal(t, test.Result.House, result, "house should be equal")
			// assert.ErrorIs(t, err, test.Result.Error, "error should be equal")
			assert.Equal(t, err, test.Result.Error, "error should be equal")
		})
	}

	repo.AssertExpectations(t)
	cache.AssertExpectations(t)
}
