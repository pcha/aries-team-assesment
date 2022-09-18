package jwt

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/software-advice/aries-team-assessment/internal/users"
)

var ErrSigningToken = errors.New("token couldn't be signed")

type HS256Manager struct {
	key []byte
}

type Claims struct {
	jwt.RegisteredClaims
	Username string
}

func (m HS256Manager) Generate(claims users.Claims) (users.Token, error) {
	tkn := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(claims.ExpiresAt()),
		},
		Username: claims.Username().String(),
	})
	strTkn, err := tkn.SignedString(m.key)
	if err != nil {
		return users.Token{}, fmt.Errorf("%w: %v", ErrSigningToken, err)
	}
	authTnk := users.ParseToken(strTkn)
	return authTnk, nil
}

func (m HS256Manager) Validate(token users.Token) (users.Claims, error) {
	var claims = Claims{}
	_, err := jwt.ParseWithClaims(token.String(), &claims, func(_ *jwt.Token) (interface{}, error) {
		return m.key, nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Name}))
	if err != nil {
		return users.NotClaims, err
	}
	return users.BuildClaims(users.BuildFrom(claims.Username, nil)), nil
}

func BuildHS256Manager(key []byte) (HS256Manager, error) {
	if len(key) == 0 {
		return HS256Manager{}, errors.New("HS256 key can't be empty")
	}
	return HS256Manager{key: key}, nil
}
