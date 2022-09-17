package creation

import (
	"context"
	"errors"
	"fmt"
	"github.com/software-advice/aries-team-assessment/internal/products"
)

var ErrMakingProduct = errors.New("it couldn't create the product")
var ErrSavingProduct = errors.New("it couldn't save the product")

// Service is the service for create products
type Service struct {
	repository products.Repository
}

// Create creates a new Product and save it in the database
func (s Service) Create(ctx context.Context, name string, description string) (products.ID, error) {
	product, err := products.Make(name, description)
	if err != nil {
		return products.ID{}, fmt.Errorf("%w - %v", ErrMakingProduct, err)
	}
	id, err := s.repository.Save(ctx, product)
	if err != nil {
		return products.ID{}, fmt.Errorf("%w - %v", ErrSavingProduct, err)
	}
	return id, nil
}

// BuildService returns a Service
func BuildService(repository products.Repository) Service {
	return Service{
		repository: repository,
	}
}
