package routes

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

type Product struct {
	Id          int       `json:"id" db:"id"`
	Name        string    `json:"name" db:"name"`
	Description string    `json:"description" db:"description"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
}

type NewProductRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func GetAllProducts(db *sqlx.DB) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var products []Product
		err := db.Select(&products, "SELECT * FROM `products` ORDER BY `name`")
		if err != nil {
			log.WithError(err).Error("Can't get products from DB :(")
		}
		return ctx.JSON(products)
	}
}

func CreateProduct(db *sqlx.DB) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		rawReq := ctx.Request()
		npReq := NewProductRequest{}
		err := json.Unmarshal(rawReq.Body(), &npReq)
		if err != nil {
			return ctx.Status(http.StatusBadRequest).JSON(map[string]string{
				"error": "invalid request body",
			})
		}
		prd := Product{
			Name:        npReq.Name,
			Description: npReq.Description,
			CreatedAt:   time.Now(),
		}
		_, err = db.NamedExec("INSERT INTO `products` (`name`, `description`, `created_at`) VALUES (:name, :description, :created_at)", prd)
		if err != nil {
			log.WithError(err).Error("Couldn't save product to DB")
			return ctx.Status(http.StatusInternalServerError).JSON(map[string]string{"error": "product not saved"})
		}
		ctx.Status(http.StatusCreated)
		return nil
	}
}
