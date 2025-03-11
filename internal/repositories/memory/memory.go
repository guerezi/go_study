package memory

import (
	"context"
	"imobiliaria/internal/repositories"
)

type Memory struct {
	users map[string]any
}

func NewMemory() *Memory {
	return &Memory{
		users: make(map[string]any),
	}
}

// CreateUser implements repositories.Repositories.
func (m *Memory) CreateUser(ctx context.Context, user any) (any, error) {
	m.users["1"] = user
	return user, nil
}

// GetNothing implements repositories.Repositories.
func (m *Memory) GetNothing(context.Context) (bool, error) {
	panic("unimplemented")
}

// GetUser implements repositories.Repositories.
func (m *Memory) GetUser(context.Context, string) (any, error) {
	return m.users["1"], nil
}

var _ repositories.Repositories = &Memory{}
