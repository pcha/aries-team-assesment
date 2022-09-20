package tokenrenew

import "github.com/software-advice/aries-team-assessment/internal/users"

type Service struct {
	tokenGenerationService users.TokenGenerationService
}

func (s Service) GetToken(claims users.Claims) (users.Token, error) {
	return s.tokenGenerationService.GetNewToken(claims)
}

func BuildService(tokenGenerationService users.TokenGenerationService) Service {
	return Service{tokenGenerationService: tokenGenerationService}
}
