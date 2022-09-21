package products

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseDescription(t *testing.T) {
	type testCase struct {
		description string
		wantErr     error
	}
	cases := map[string]testCase{
		"valid description": {
			description: "some name",
		},
		"empty description": {
			description: "",
			wantErr:     DescriptionCanNotBeEmptyErr,
		},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			desc, err := ParseDescription(tc.description)
			if tc.wantErr != nil {
				assert.ErrorIs(t, err, tc.wantErr)
				return
			}
			assert.Nil(t, err)
			assert.Equal(t, tc.description, desc.String())
		})
	}
}
