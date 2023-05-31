package products

import "context"

// Repository provide persistence interaction with the struct Product.
type Repository interface {
	// Search products, if the term is not empty it is used to filter the results
	Search(ctx context.Context, term string) ([]Product, error)
	// Save persists the given Product and the assigned ID.
	Save(ctx context.Context, product Product) (ID, error)
}
