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

func (p *ProductRepository) Search(ctx context.Context, term string) ([]products.Product, error) {
	args := p.Called(ctx, term)
	return args.Get(0).([]products.Product), args.Error(1)
}

// Save mocks the saving of a product
func (p *ProductRepository) Save(ctx context.Context, product products.Product) (products.ID, error) {
	args := p.Called(ctx, product)
	return args.Get(0).(products.ID), args.Error(1)
}
