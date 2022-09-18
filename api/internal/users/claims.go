package users

import "time"

type Claims struct {
	username  Username
	expiresAt time.Time
}

var NotClaims = Claims{}

func (c Claims) Username() Username {
	return c.username
}

func (c Claims) ExpiresAt() time.Time {
	return c.expiresAt
}

func BuildClaims(user User) Claims {
	return Claims{
		username:  user.Username(),
		expiresAt: time.Now().Add(time.Hour),
	}
}
