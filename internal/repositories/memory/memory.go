package memory

import (
	"context"
	"imobiliaria/internal/models"
	"imobiliaria/internal/repositories"
)

type Memory struct {
	users map[int]*models.User
}

func NewMemory() *Memory {
	return &Memory{
		users: make(map[int]*models.User),
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

var _ repositories.Repositories = &Memory{}
