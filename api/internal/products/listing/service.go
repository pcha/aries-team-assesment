package listing

import (
	"context"
	"github.com/software-advice/aries-team-assessment/internal/products"
)

// Service is an application service to list products
type Service struct {
	repository products.Repository
}

// List executes the search for the given pattern
func (s Service) List(ctx context.Context, filter string) ([]products.Product, error) {
	return s.repository.Search(ctx, filter)
}

// BuildService returns a Service with the given repository
func BuildService(repository products.Repository) Service {
	return Service{
		repository: repository,
	}
}
