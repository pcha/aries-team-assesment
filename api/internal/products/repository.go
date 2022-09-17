package products

import "context"

// Repository provide persistence interaction with the struct Product.
type Repository interface {
	//GetAll returns all the persisted Products.
	GetAll(ctx context.Context) ([]Product, error)
	// Save persists the given Product and the assigned ID.
	Save(ctx context.Context, product Product) (ID, error)
}
