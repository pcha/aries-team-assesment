package routes

import (
	"encoding/json"
	"github.com/DATA-DOG/go-sqlmock"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"io"
	"net/http/httptest"
	"testing"
	"time"
)

func TestGetAllProducts(t *testing.T) {

	db, mockDb, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Couldn't start mock db. %s", err)
	}
	defer db.Close()

	app := fiber.New()
	app.Get("/products", GetAllProducts(sqlx.NewDb(db, "mysql")))

	mockDb.ExpectQuery("SELECT *").WillReturnRows(
		sqlmock.NewRows([]string{"id", "name", "description", "created_at"}).
			AddRow("1111", "Some Name", "Some description", time.Now()))

	req := httptest.NewRequest("GET", "/products", nil)
	res, _ := app.Test(req)
	var resJson []Product
	reader, _ := io.ReadAll(res.Body)
	err = json.Unmarshal(reader, &resJson)
	if err != nil {
		t.Error("Couldn't parse response into JSON")
	}

	if res.StatusCode != 200 {
		t.Errorf("Status code isn't 200, got %d instead", res.StatusCode)
	}

	if len(resJson) != 1 {
		t.Errorf("Response didn't have 1 product, got %d instead", len(resJson))
	}
}
