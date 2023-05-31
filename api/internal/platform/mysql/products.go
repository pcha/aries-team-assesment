package mysql

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/software-advice/aries-team-assessment/internal/products"
	"time"
)

// Product is a DTO for represent the Product in the database
type Product struct {
	ID          int64     `db:"id"`
	Name        string    `db:"name"`
	Description string    `db:"description"`
	CreatedAt   time.Time `db:"created_at"`
}

// ProductRepository is a products.Repository to manage sql persistence
type ProductRepository struct {
	db *sqlx.DB
}

// Search returns all the products matching the given term.
func (p ProductRepository) Search(ctx context.Context, term string) ([]products.Product, error) {
	var result []Product
	selectPart := "SELECT * FROM `products` "
	wherePart := "WHERE `name` LIKE ? OR  `description` LIKE ? "
	orderPart := "ORDER BY `name`"
	query := selectPart
	var args []interface{}
	if term != "" {
		query += wherePart
		wildcardedTerm := "%" + term + "%"
		args = append(args, wildcardedTerm, wildcardedTerm)
	}
	query += orderPart
	err := p.db.SelectContext(ctx, &result, query, args...)
	if err != nil {
		return nil, fmt.Errorf("mysql: error selecting products matching term %q - %v", term, err)
	}
	prods := make([]products.Product, len(result))
	for i, dto := range result {
		prod := products.BuildFrom(dto.ID, dto.Name, dto.Description, dto.CreatedAt)
		prods[i] = prod
	}
	return prods, nil
}

// Save writes the info in the table products.
// If the given engine supports it, the first parameter is the last inserted ID.
func (p ProductRepository) Save(ctx context.Context, product products.Product) (products.ID, error) {
	//TODO implement me
	res, err := p.db.ExecContext(
		ctx,
		"INSERT INTO `products` (`name`, `description`, `created_at`) VALUES (?, ?, ?)",
		product.Name().String(),
		product.Description().String(),
		product.CreatedAt().Time(),
	)
	if err != nil {
		return products.ID{}, fmt.Errorf("mysql: error saving product %v - %v", product, err)
	}
	id, _ := res.LastInsertId() // In some engines last insert id could not be supported and return error,
	// but it doesn't mean than the product wasn't saved, so a 0 id is better than an error
	return products.ParseID(id), nil
}

// NewProductRepository returns a new *ProductRepository for the given db object.
func NewProductRepository(db *sqlx.DB) *ProductRepository {
	return &ProductRepository{db: db}
}
