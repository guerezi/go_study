package memory

import (
	"context"

	"imobiliaria/internal/models"
	"imobiliaria/internal/repositories"
)

type Memory struct {
	users map[int]*models.User
	// houses map[int]*models.House
}

func NewMemory() *Memory {
	return &Memory{
		users: make(map[int]*models.User),
		// houses: make(map[int]*models.House),
	}
}

// CreateUser implements repositories.Repositories.
func (m *Memory) CreateUser(ctx context.Context, user *models.User) (*models.User, error) {
	id := len(m.users) + 1
	user.ID = id
	m.users[id] = user

	return m.users[id], nil
}

// GetUser implements repositories.Repositories.
func (m *Memory) GetUser(ctx context.Context, id int) (*models.User, error) {
	return m.users[id], nil
}

// CreateHouse implements repositories.Repositories.
func (m *Memory) CreateHouse(ctx context.Context, house *models.House) (*models.House, error) {
	panic("unimplemented")
}

// DeleteHouse implements repositories.Repositories.
func (m *Memory) DeleteHouse(ctx context.Context, id int) error {
	panic("unimplemented")
}

// GetHouse implements repositories.Repositories.
func (m *Memory) GetHouse(ctx context.Context, id int) (*models.House, error) {
	panic("unimplemented")
}

// GetHouses implements repositories.Repositories.
func (m *Memory) GetHouses(ctx context.Context) ([]*models.House, error) {
	panic("unimplemented")
}

// GetHousesByUserID implements repositories.Repositories.
func (m *Memory) GetHousesByUserID(ctx context.Context, id int) ([]*models.House, error) {
	panic("unimplemented")
}

// UpdateHouse implements repositories.Repositories.
func (m *Memory) UpdateHouse(ctx context.Context, house *models.House) (*models.House, error) {
	panic("unimplemented")
}

var _ repositories.Repositories = &Memory{}
