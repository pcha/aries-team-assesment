package searching

import (
	"context"
	"errors"
	"github.com/software-advice/aries-team-assessment/internal/products"
)

var ErrEmptyTerm = errors.New("search term can't be empty")

// Service allows to search products by a term
type Service struct {
	repository products.Repository
}

// Search executes the search for the given pattern
func (s Service) Search(ctx context.Context, term string) ([]products.Product, error) {
	if term == "" {
		return nil, ErrEmptyTerm
	}
	return s.repository.Search(ctx, term)
}

func BuildService(repository products.Repository) Service {
	return Service{
		repository: repository,
	}
}
