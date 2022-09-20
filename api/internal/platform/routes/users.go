package routes

import (
	"encoding/json"
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/software-advice/aries-team-assessment/internal/users"
	"github.com/software-advice/aries-team-assessment/internal/users/login"
	"github.com/software-advice/aries-team-assessment/internal/users/signup"
	"net/http"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type SignUpRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
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
		return ctx.Status(http.StatusOK).JSON(LoginResponse{Token: tkn.String()})
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
