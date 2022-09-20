package routes

import (
	"encoding/json"
	"errors"
	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
	"github.com/software-advice/aries-team-assessment/internal/users"
	"github.com/software-advice/aries-team-assessment/internal/users/login"
	"github.com/software-advice/aries-team-assessment/internal/users/signup"
	"github.com/software-advice/aries-team-assessment/internal/users/tokenrenew"
	"net/http"
	"time"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type SignUpRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Claims struct {
	Username string `json:"username"`
}

type TokenResponse struct {
	Token  string `json:"token"`
	Claims Claims `json:"claims"`
}

func Login(service login.Service) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var req LoginRequest
		reqBody := ctx.Request().Body()
		err := json.Unmarshal(reqBody, &req)
		if err != nil {
			return ctx.
				Status(http.StatusBadRequest).
				JSON(ErrorResponse{Error: "invalid request body"})
		}
		tkn, err := service.Login(ctx.Context(), req.Username, []byte(req.Password))
		if err != nil {
			if errors.Is(err, users.ErrInvalidPassword) || errors.Is(err, users.ErrUserNotFound) {
				return ctx.
					Status(http.StatusUnauthorized).
					JSON(ErrorResponse{
						Error: "invalid username or password",
					})
			}
			return ctx.
				Status(http.StatusInternalServerError).
				JSON(ErrorResponse{
					Error: "internal error",
				})
		}
		return sendTokenResponse(ctx, tkn)
	}
}

func SignUp(service signup.Service) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var req SignUpRequest
		reqBody := ctx.Request().Body()
		err := json.Unmarshal(reqBody, &req)
		if err != nil {
			return ctx.
				Status(http.StatusBadRequest).
				JSON(ErrorResponse{Error: "invalid request body"})
		}
		err = service.SignUp(ctx.Context(), req.Username, []byte(req.Password))
		if err != nil {
			if errors.Is(err, signup.ErrMakingUser) {
				return ctx.
					Status(http.StatusBadRequest).
					JSON(ErrorResponse{
						Error: err.Error(),
					})
			}
			return ctx.Status(http.StatusInternalServerError).JSON(ErrorResponse{
				Error: "internal error",
			})
		}
		ctx.Status(http.StatusCreated)
		return nil
	}
}

func RenewToken(service tokenrenew.Service) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		claimsVal := ctx.UserContext().Value(ctxClaimsKey)
		claims, ok := claimsVal.(users.Claims)
		if !ok {
			log.Errorf("invalid claims value: %v", claimsVal)
			return sendInternalErrorResponse(ctx)
		}
		tkn, err := service.GetToken(claims)
		if err != nil {
			return sendInternalErrorResponse(ctx)
		}
		return sendTokenResponse(ctx, tkn)
	}
}

func sendTokenResponse(ctx *fiber.Ctx, tkn users.Token) error {
	ctx.Set(fiber.HeaderExpires, tkn.Claims().ExpiresAt().Format(time.RFC1123))
	return ctx.
		Status(http.StatusOK).
		JSON(TokenResponse{
			Token: tkn.TokenString().String(),
			Claims: Claims{
				Username: tkn.Claims().Username().String(),
			},
		})
}
