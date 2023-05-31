package products

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseName(t *testing.T) {
	type testCase struct {
		name    string
		wantErr error
	}
	cases := map[string]testCase{
		"valid name": {
			name: "some name",
		},
		"empty name": {
			name:    "",
			wantErr: NameCanNotBeEmptyErr,
		},
	}
	for caseName, tc := range cases {
		t.Run(caseName, func(t *testing.T) {
			desc, err := ParseName(tc.name)
			if tc.wantErr != nil {
				assert.ErrorIs(t, err, tc.wantErr)
				return
			}
			assert.Nil(t, err)
			assert.Equal(t, tc.name, desc.String())
		})
	}
}
