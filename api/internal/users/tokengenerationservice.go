package users

import (
	"errors"
	"fmt"
	"time"
)

var ErrGeneratingToken = errors.New("error generating token")

type TokenGenerationService struct {
	generator      TokenGenerator
	tokensDuration time.Duration
}

func (s TokenGenerationService) GetNewToken(claims Claims) (Token, error) {
	claims.expiresAt = time.Now().Add(s.tokensDuration)
	token, err := s.generator.Generate(claims)
	if err != nil {
		return Token{}, fmt.Errorf("%w: %v", ErrGeneratingToken, err)
	}
	return Token{
		claims: claims,
		str:    token,
	}, nil
}

func BuildTokenGenerationService(generator TokenGenerator, tokensDuration time.Duration) TokenGenerationService {
	return TokenGenerationService{generator: generator, tokensDuration: tokensDuration}
}
