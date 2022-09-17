package mockable

import (
	"context"
	"github.com/software-advice/aries-team-assessment/internal/products"
	"github.com/stretchr/testify/mock"
)

// ProductRepository is a mockable implementing products.Repository
type ProductRepository struct {
	mock.Mock
}

// GetAll mocks the query of all products
func (p *ProductRepository) GetAll(ctx context.Context) ([]products.Product, error) {
	args := p.Called(ctx)
	return args.Get(0).([]products.Product), args.Error(1)
}

// Save mocks the saving of a product
func (p *ProductRepository) Save(ctx context.Context, product products.Product) (products.ID, error) {
	args := p.Called(ctx, product)
	return args.Get(0).(products.ID), args.Error(1)
}
