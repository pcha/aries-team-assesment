package tokenvalidation

import "github.com/software-advice/aries-team-assessment/internal/users"

type Service struct {
	validator users.TokenValidator
}

func (s Service) Validate(token string) (users.Claims, error) {
	tkn := users.ParseTokenString(token)
	claims, err := s.validator.Validate(tkn)
	if err != nil {
		return users.NotClaims, err
	}
	return claims, nil
}

func BuildService(validator users.TokenValidator) Service {
	return Service{validator: validator}
}
