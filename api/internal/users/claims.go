package users

import "time"

// Claims is a struct with the data claimed by a Token.
type Claims struct {
	username  Username
	expiresAt time.Time
}

// NotClaims is an empty Claims.
var NotClaims = Claims{}

// Username returns the claimed Username.
func (c Claims) Username() Username {
	return c.username
}

// ExpiresAt returns the claimed expiration time.
func (c Claims) ExpiresAt() time.Time {
	return c.expiresAt
}

// BuildClaims return a Claims with the given parameters.
func BuildClaims(username Username, expirationTime time.Time) Claims {
	return Claims{
		username:  username,
		expiresAt: expirationTime,
	}
}
