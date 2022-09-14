package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
	"time"
)

type Product struct {
	Id          int       `json:"id" db:"id"`
	Name        string    `json:"name" db:"name"`
	Description string    `json:"description" db:"description"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
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
