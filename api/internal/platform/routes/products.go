package routes

import (
	"encoding/json"
	"errors"
	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
	"github.com/software-advice/aries-team-assessment/internal/products/creation"
	"github.com/software-advice/aries-team-assessment/internal/products/listing"
	"net/http"
	"time"
)

// Product represents how a product is show to the frontend.
type Product struct {
	Id          int64     `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}

// CreateProductRequest represents the expected request to create a Product.
type CreateProductRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// ErrorResponse represents the response to return if something fails.
type ErrorResponse struct {
	Error string `json:"error"`
}

// ProductCreatedResponse represents the response to return if a product is created successfully.
type ProductCreatedResponse struct {
	ID int64 `json:"id"`
}

// GetAllProducts is the handler to return all the products.
func GetAllProducts(service listing.Service) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		prods, err := service.List(ctx.Context())
		if err != nil {
			log.WithError(err).Error("Can't get products from DB :(")
		}
		res := make([]Product, len(prods))
		for i, prod := range prods {
			res[i] = Product{
				Id:          prod.ID().Int64(),
				Name:        prod.Name().String(),
				Description: prod.Description().String(),
				CreatedAt:   prod.CreatedAt().Time(),
			}
		}
		return ctx.JSON(res)
	}
}

// CreateProduct is the handler to create a new product
func CreateProduct(service creation.Service) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		rawReq := ctx.Request()
		npReq := CreateProductRequest{}
		err := json.Unmarshal(rawReq.Body(), &npReq)
		if err != nil {
			return ctx.Status(http.StatusBadRequest).JSON(ErrorResponse{Error: "Invalid request body - " + err.Error()})
		}

		id, err := service.Create(ctx.Context(), npReq.Name, npReq.Description)

		if err != nil {
			log.WithError(err).Error("Couldn't create product.")
			if errors.Is(err, creation.ErrMakingProduct) {
				return ctx.
					Status(http.StatusBadRequest).
					JSON(ErrorResponse{
						Error: err.Error(),
					})
			}
			return ctx.
				Status(http.StatusInternalServerError).
				JSON(ErrorResponse{
					Error: "internal error",
				})
		}
		return ctx.
			Status(http.StatusCreated).
			JSON(ProductCreatedResponse{
				ID: id.Int64(),
			})
	}
}
