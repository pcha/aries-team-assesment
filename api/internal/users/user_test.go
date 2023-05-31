package users

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUser_IsNotUser(t *testing.T) {
	type testCase struct {
		user   User
		assert bool
	}
	cases := map[string]testCase{
		"not user": {
			user:   NotUsr,
			assert: true,
		},
		"a user": {
			user: BuildUnsafe("something", []byte("")),
		},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.assert, tc.user.IsNotUser())
		})
	}
}
