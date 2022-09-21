package routes

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
)

// ErrorResponse represents the response to return if something fails.
type ErrorResponse struct {
	Error string `json:"error"`
}

// internalErrorResponse is a generic internal error response
var internalErrorResponse = ErrorResponse{
	Error: "internal error",
}

// sendInternalErrorResponse set the status 500 and send internalErrorResponse
func sendInternalErrorResponse(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusInternalServerError).JSON(internalErrorResponse)
}
