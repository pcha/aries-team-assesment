package routes

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
)

// ErrorResponse represents the response to return if something fails.
type ErrorResponse struct {
	Error string `json:"error"`
}

var internalErrorResponse = ErrorResponse{
	Error: "internal error",
}

func sendInternalErrorResponse(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusInternalServerError).JSON(internalErrorResponse)
}
