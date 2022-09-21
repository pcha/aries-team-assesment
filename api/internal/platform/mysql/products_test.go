package mysql

import (
	"context"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/software-advice/aries-team-assessment/internal/products"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestProductRepository_Search(t *testing.T) {
	db, mockDB, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()
	repository := NewProductRepository(sqlx.NewDb(db, "mysql"))

	mockDB.
		ExpectQuery("SELECT *").
		WithArgs("%prod%", "%prod%").
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "name", "description", "created_at"}).
				AddRow(1, "product 1", "description 1", time.Now()).
				AddRow(2, "product 2", "description 2", time.Now()),
		)

	prods, err := repository.Search(context.Background(), "prod")
	require.NoError(t, err)
	require.Len(t, prods, 2)
	assert.Equal(t, int64(1), prods[0].ID().Int64())
	assert.Equal(t, "product 1", prods[0].Name().String())
	assert.Equal(t, "description 1", prods[0].Description().String())
	assert.Equal(t, int64(2), prods[1].ID().Int64())
	assert.Equal(t, "product 2", prods[1].Name().String())
	assert.Equal(t, "description 2", prods[1].Description().String())
	assert.NoError(t, mockDB.ExpectationsWereMet())
}

func TestProductRepository_Save(t *testing.T) {
	db, mockDB, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()
	repository := NewProductRepository(sqlx.NewDb(db, "mysql"))

	testName := "test"
	testDesc := "test product"
	prod, err := products.Make(testName, testDesc)
	require.NoError(t, err)

	lastInsertID := int64(3)
	mockDB.
		ExpectExec("INSERT INTO `products`").
		WithArgs(testName, testDesc, sqlmock.AnyArg()).
		WillReturnResult(sqlmock.NewResult(lastInsertID, 1))

	id, err := repository.Save(context.Background(), prod)
	require.NoError(t, err)
	assert.Equal(t, lastInsertID, id.Int64())
	assert.NoError(t, mockDB.ExpectationsWereMet())
}
