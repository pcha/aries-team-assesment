package users

import (
	"errors"
	"fmt"
	"time"
)

var ErrGeneratingToken = errors.New("error generating token")

type TokenGenerationService struct {
	generator TokenGenerator
}

func (s TokenGenerationService) GetNewToken(claims Claims) (Token, error) {
	claims.expiresAt = time.Now().Add(15 * time.Minute)
	token, err := s.generator.Generate(claims)
	if err != nil {
		return Token{}, fmt.Errorf("%w: %v", ErrGeneratingToken, err)
	}
	return Token{
		claims: claims,
		str:    token,
	}, nil
}

func BuildTokenGenerationService(generator TokenGenerator) TokenGenerationService {
	return TokenGenerationService{generator: generator}
}
