package repositories

import "context"

type Nothings interface {
	GetNothing(context.Context) (bool, error)
}
