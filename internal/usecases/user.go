package usecases

import "context"

type Users interface {
	CreateUser(context.Context, any) (any, error)
	GetUser(context.Context, string) (any, error)
}

// CreateUser implements Usecases.
func (u *usecases) CreateUser(ctx context.Context, user any) (any, error) {
	return u.repo.CreateUser(ctx, user)
}

// GetUser implements Usecases.
func (u *usecases) GetUser(ctx context.Context, id string) (any, error) {
	return u.repo.GetUser(ctx, id)
}
