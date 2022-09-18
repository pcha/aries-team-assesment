package routes

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/software-advice/aries-team-assessment/internal/users/tokenvalidation"
	"net/http"
	"strings"
)

func VerifyToken(service tokenvalidation.Service) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		auth := ctx.Get(fiber.HeaderAuthorization)
		tkn, err := extractBearer(auth)
		if err != nil {
			return ctx.Status(http.StatusUnauthorized).JSON(ErrorResponse{Error: err.Error()})
		}
		_, err = service.Validate(tkn)
		if err != nil {
			return ctx.Status(http.StatusUnauthorized).JSON(ErrorResponse{Error: "invalid token"})
		}
		return ctx.Next()
	}
}

func extractBearer(auth string) (string, error) {
	bearerPrefix := "Bearer "
	if !strings.HasPrefix(auth, bearerPrefix) {
		return "", errors.New("invalid authorization format")
	}
	tkn := auth[len(bearerPrefix):]
	return tkn, nil
}
