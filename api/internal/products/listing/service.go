package listing

import (
	"context"
	"github.com/software-advice/aries-team-assessment/internal/products"
)

// Service for listing all the products.
type Service struct {
	repository products.Repository
}

// List returns all the products or an error if something fail.
func (s Service) List(ctx context.Context) ([]products.Product, error) {
	prods, err := s.repository.GetAll(ctx)
	return prods, err
}

// BuildService returns a Service
func BuildService(repository products.Repository) Service {
	return Service{
		repository: repository,
	}
}
