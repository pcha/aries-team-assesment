package routes

import (
	"encoding/json"
	"errors"
	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
	"github.com/software-advice/aries-team-assessment/internal/products"
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

// ProductCreatedResponse represents the response to return if a product is created successfully.
type ProductCreatedResponse struct {
	ID int64 `json:"id"`
}

// CreateProduct returns a http handler to create a new product.
func CreateProduct(service creation.Service) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		// request parsing
		rawReq := ctx.Request()
		npReq := CreateProductRequest{}
		err := json.Unmarshal(rawReq.Body(), &npReq)
		if err != nil {
			return ctx.Status(http.StatusBadRequest).JSON(ErrorResponse{Error: "Invalid request body - " + err.Error()})
		}

		// service call
		id, err := service.Create(ctx.Context(), npReq.Name, npReq.Description)

		// response parsing
		if err != nil {
			if errors.Is(err, creation.ErrMakingProduct) {
				return ctx.
					Status(http.StatusBadRequest).
					JSON(ErrorResponse{
						Error: err.Error(),
					})
			}
			log.WithError(err).Error("Couldn't create product.")
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

// GetProducts returns a http handler to return the products. I
// f a query param "q" is passed it will filter the products using the given values
func GetProducts(service listing.Service) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		// parse request
		term := ctx.Query("q")

		// service call
		prods, err := service.List(ctx.Context(), term)

		// parse response
		if err != nil {
			log.WithError(err).Error("Can't get products from DB :(")
			return ctx.
				Status(http.StatusInternalServerError).
				JSON(internalErrorResponse)
		}
		res := parseProductsList(prods)
		return ctx.
			Status(http.StatusOK).
			JSON(res)
	}

}

func parseProductsList(prods []products.Product) []Product {
	res := make([]Product, len(prods))
	for i, prod := range prods {
		res[i] = Product{
			Id:          prod.ID().Int64(),
			Name:        prod.Name().String(),
			Description: prod.Description().String(),
			CreatedAt:   prod.CreatedAt().Time(),
		}
	}
	return res
}
