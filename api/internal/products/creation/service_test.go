package creation

import (
	"context"
	"errors"
	mocked "github.com/software-advice/aries-team-assessment/internal/platform/mockable"
	"github.com/software-advice/aries-team-assessment/internal/products"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

type args struct {
	name        string
	description string
}
type mockRepo struct {
	called bool
	id     int64
	err    error
}
type want struct {
	id  int64
	err error
}
type testCase struct {
	args args
	mock mockRepo
	want want
}

func TestProductCreationService_Create(t *testing.T) {
	cases := map[string]testCase{
		"ok": {
			args: args{
				name:        "prod",
				description: "test product",
			},
			mock: mockRepo{
				called: true,
				id:     1,
				err:    nil,
			},
			want: want{
				id:  1,
				err: nil,
			},
		},
		"no name": {
			args: args{
				name:        "",
				description: "test desc",
			},
			mock: mockRepo{
				called: false,
			},
			want: want{
				id:  0,
				err: ErrMakingProduct,
			},
		},
		"no description": {
			args: args{
				name:        "prod",
				description: "",
			},
			mock: mockRepo{
				called: false,
			},
			want: want{
				id:  0,
				err: ErrMakingProduct,
			},
		},
		"err on repo": {
			args: args{
				name:        "prod",
				description: "test prod",
			},
			mock: mockRepo{
				called: true,
				err:    errors.New("repo err"),
			},
			want: want{
				id:  0,
				err: ErrSavingProduct,
			},
		},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			repository := new(mocked.ProductRepository)
			service := Service{repository: repository}

			repository.On("Save", mock.Anything, mock.MatchedBy(func(prod products.Product) bool {
				return prod.Name().String() == tc.args.name &&
					prod.Description().String() == tc.args.description &&
					prod.CreatedAt().Time().Before(time.Now()) &&
					timeIsAroundNow(prod.CreatedAt().Time())
			})).Return(products.ParseID(tc.mock.id), tc.mock.err)

			id, err := service.Create(context.Background(), tc.args.name, tc.args.description)

			if tc.want.err != nil {
				assert.ErrorIs(t, err, tc.want.err)
			} else {
				assert.Nil(t, err)
				assert.Equal(t, tc.want.id, id.Int64())
			}
		})
	}

}

// timeIsAroundNow check that the given time os not too far from now.
//The purpose is to assert than the given time was generated with time.Now()
func timeIsAroundNow(t time.Time) bool {
	return t.Before(time.Now()) && t.Add(time.Second).After(time.Now())
}
