package mysql

import (
	"context"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestProductRepository_GetAll(t *testing.T) {
	db, mockDB, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()
	repository := NewProductRepository(sqlx.NewDb(db, "mysql"))

	mockDB.
		ExpectQuery("SELECT *").
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "name", "description", "created_at"}).
				AddRow(1, "product 1", "description 1", time.Now()).
				AddRow(2, "product 2", "description 2", time.Now()),
		)

	prods, err := repository.GetAll(context.Background())
	require.NoError(t, err)
	require.Len(t, prods, 2)
	assert.Equal(t, int64(1), prods[0].ID().Int64())
	assert.Equal(t, "product 1", prods[0].Name().String())
	assert.Equal(t, "description 1", prods[0].Description().String())
	assert.Equal(t, int64(2), prods[1].ID().Int64())
	assert.Equal(t, "product 2", prods[1].Name().String())
	assert.Equal(t, "description 2", prods[1].Description().String())

}

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

}
