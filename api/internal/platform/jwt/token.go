package jwt

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/software-advice/aries-team-assessment/internal/users"
)

// ErrSigningToken is an error to return if the HS256Manager fails signing the token
var ErrSigningToken = errors.New("jwt: token couldn't be signed")

// ErrValidatingToken is an error to return if the token validation fails
var ErrValidatingToken = errors.New("jwt: invalid token")

var ErrEmptyKey = errors.New("HS256 key can't be empty")

// HS256Manager is a users.TokenGenerator and users.TokenValidator using the HS256 method
type HS256Manager struct {
	key []byte
}

// The Claims to include in the Token
type Claims struct {
	jwt.RegisteredClaims
	Username string
}

// Generate signs a new users.TokenString for the given claims.
// If it can't sign a new token it'll return ErrSigningToken
func (m HS256Manager) Generate(claims users.Claims) (users.TokenString, error) {
	tkn := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(claims.ExpiresAt()),
		},
		Username: claims.Username().String(),
	})
	strTkn, err := tkn.SignedString(m.key)
	if err != nil {
		return users.TokenString{}, fmt.Errorf("%w - %v", ErrSigningToken, err)
	}
	authTnk := users.ParseTokenString(strTkn)
	return authTnk, nil
}

// Validate verity that the given token is valid and return the corresponding users.Claims.
// If the token is invalid it'll return ErrValidatingToken
func (m HS256Manager) Validate(token users.TokenString) (users.Claims, error) {
	var claims = Claims{}
	_, err := jwt.ParseWithClaims(token.String(), &claims, func(_ *jwt.Token) (interface{}, error) {
		return m.key, nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Name}))
	if err != nil {
		return users.NotClaims, fmt.Errorf("%w - %v", ErrValidatingToken, err)
	}
	return users.BuildClaims(users.ParseUnsafeUsername(claims.Username), claims.ExpiresAt.Time), nil
}

// BuildHS256Manager returns a HS256Manager with the given signing&verifying key.
// If the key is empty it'll return a ErrEmptyKey.
func BuildHS256Manager(key []byte) (HS256Manager, error) {
	if len(key) == 0 {
		return HS256Manager{}, ErrEmptyKey
	}
	return HS256Manager{key: key}, nil
}
