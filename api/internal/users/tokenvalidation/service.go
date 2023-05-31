package tokenvalidation

import "github.com/software-advice/aries-team-assessment/internal/users"

// Service is an application service to handle the validation of tokens
type Service struct {
	validator users.TokenValidator
}

// Validate receives a token string and, if it is valid, return the corresponding users.Claims.
// If the token is invalid it returns an error
func (s Service) Validate(token string) (users.Claims, error) {
	tkn := users.ParseTokenString(token)
	claims, err := s.validator.Validate(tkn)
	if err != nil {
		return users.NotClaims, err
	}
	return claims, nil
}

// BuildService returns a Service with the given users.TokenValidator
func BuildService(validator users.TokenValidator) Service {
	return Service{validator: validator}
}
