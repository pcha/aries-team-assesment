package mysql

import (
	"context"
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

// ProductRepository is a repository to interact with the table products
type ProductRepository struct {
	db *sqlx.DB
}

// GetAll returns all the products saved in the db
func (p ProductRepository) GetAll(ctx context.Context) ([]products.Product, error) {
	var result []Product
	err := p.db.SelectContext(ctx, &result, "SELECT * FROM `products` ORDER BY `name`")
	if err != nil {
		return nil, err
	}
	prods := make([]products.Product, len(result))
	for i, dto := range result {
		prod := products.BuildFrom(dto.ID, dto.Name, dto.Description, dto.CreatedAt)
		prods[i] = prod
	}
	return prods, nil
}

func (p ProductRepository) Search(ctx context.Context, term string) ([]products.Product, error) {
	var result []Product
	wildcardedTerm := "%" + term + "%"
	err := p.db.SelectContext(ctx, &result, "SELECT * FROM `products` WHERE `name` LIKE ? OR  `description` LIKE ? ORDER BY `name`", wildcardedTerm, wildcardedTerm)
	if err != nil {
		return nil, err
	}
	prods := make([]products.Product, len(result))
	for i, dto := range result {
		prod := products.BuildFrom(dto.ID, dto.Name, dto.Description, dto.CreatedAt)
		prods[i] = prod
	}
	return prods, nil
}

// Save writes the info in the table products
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
		return products.ID{}, err
	}
	id, _ := res.LastInsertId() // In some engines last insert id could not be supported and return error,
	// but it doesn't mean than the product wasn't saved, so a 0 id is better than an error
	return products.ParseID(id), nil
}

// NewProductRepository returns a new ProductRepository
func NewProductRepository(db *sqlx.DB) *ProductRepository {
	return &ProductRepository{db: db}
}
