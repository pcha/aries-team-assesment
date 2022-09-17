package creation

import (
	"context"
	mocked "github.com/software-advice/aries-team-assessment/internal/platform/mockable"
	"github.com/software-advice/aries-team-assessment/internal/products"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

func TestProductCreationService_Create(t *testing.T) {
	repository := new(mocked.ProductRepository)
	service := Service{repository: repository}
	name := "test"
	description := "test product"

	repository.On("Save", mock.Anything, mock.MatchedBy(func(prod products.Product) bool {
		return prod.Name().String() == name &&
			prod.Description().String() == description &&
			prod.CreatedAt().Time().Before(time.Now()) &&
			timeIsAroundNow(prod.CreatedAt().Time())
	})).Return(products.ParseID(1), nil)

	id, err := service.Create(context.Background(), name, description)

	assert.Equal(t, int64(1), id.Int64())
	assert.Nil(t, err)
}

// timeIsAroundNow check that the given time os not too far from now.
//The purpose is to assert than the given time was generated with time.Now()
func timeIsAroundNow(t time.Time) bool {
	return t.Before(time.Now()) && t.Add(time.Hour).After(time.Now())
}
