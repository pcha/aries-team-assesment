package tokenrenew

import "github.com/software-advice/aries-team-assessment/internal/users"

// Service is an application service that handle the Token renovation
type Service struct {
	tokenGenerationService users.TokenGenerationService
}

// GetToken receives the current claims and return a new Token with renovated claims.
func (s Service) GetToken(claims users.Claims) (users.Token, error) {
	return s.tokenGenerationService.GetNewToken(claims)
}

// BuildService returns a Service with the given users.TokenGenerationService
func BuildService(tokenGenerationService users.TokenGenerationService) Service {
	return Service{tokenGenerationService: tokenGenerationService}
}
