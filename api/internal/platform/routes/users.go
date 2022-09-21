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

// LoginRequest id the DTO used to parse the login request body
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// SignUpRequest is the DTO used to parse the signup request
type SignUpRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Claims has the token claims
type Claims struct {
	Username string `json:"username"`
}

// TokenResponse is the response to return when the client expects a token
type TokenResponse struct {
	Token  string `json:"token"`
	Claims Claims `json:"claims"`
}

// Login returns a handler that executes the login
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
			if errors.Is(err, users.ErrInvalidPassword) || errors.Is(err, login.ErrUserNotFound) {
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

// SignUp returns a handler that create a new users.User
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
			if errors.Is(err, signup.ErrAlreadyTakenUsername) {
				return ctx.
					Status(http.StatusConflict).
					JSON(ErrorResponse{Error: "username already taken"})
			}
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

// RenewToken returns a handler that responds with a new token for the claims set in the ctx.UserContext.
// This handler assumes that it's preceded by VerifyToken middleware who sets the claims in th ctx.UserContext
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

// send the TokenResponse with the expiration in the headers
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
