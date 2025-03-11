package repositories

import "context"

type Users interface {
	CreateUser(context.Context, any) (any, error)
	GetUser(context.Context, string) (any, error)
}
