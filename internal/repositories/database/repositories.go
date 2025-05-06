package repositories

//go:generate mockery --name=Repositories --output=./mocks --filename=repositories.go
type Repositories interface {
	Users
	Houses
	// Banks
	// Transactions
}

// two repositories
// database one
// cache one
