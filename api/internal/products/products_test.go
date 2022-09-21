package products

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMake(t *testing.T) {
	type testCase struct {
		name        string
		description string
		wantErr     error
	}
	cases := map[string]testCase{
		"valid": {
			name:        "some name",
			description: "some description",
		},
		"empty name": {
			name:        "",
			description: "some description",
			wantErr:     DescriptionCanNotBeEmptyErr,
		},
		"empty description": {
			name:        "some name",
			description: "",
			wantErr:     NameCanNotBeEmptyErr,
		},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			prod, err := Make(tc.name, tc.description)
			if tc.wantErr != nil {
				assert.Equal(t, err, tc.wantErr)
				return
			}
			assert.Nil(t, err)
			assert.Equal(t, tc.name, prod.Name().String())
			assert.Equal(t, tc.description, prod.Description().String())
			assert.NotEmpty(t, prod.CreatedAt().Time())
		})
	}
}
