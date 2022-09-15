package routes

import (
	"encoding/json"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
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

func TestCreateProduct_OK(t *testing.T) {
	db, mockDb, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Couldn't start mock db. %s", err)
	}
	defer db.Close()
	app := fiber.New()
	app.Post("/products", CreateProduct(sqlx.NewDb(db, "mysql")))

	mockDb.ExpectExec("INSERT INTO `products`").WithArgs("test product", "It is a test product", sqlmock.AnyArg()).WillReturnResult(sqlmock.NewResult(1, 1))
	req := httptest.NewRequest("POST", "/products", strings.NewReader(`{
	"name": "test product",
	"description": "It is a test product"
}`))
	res, err := app.Test(req)
	if err != nil {
		t.Fatalf("Error on faking request. %s", err)
	}
	if res.StatusCode != http.StatusCreated {
		t.Errorf("Status code isn't %d, got %d instead", http.StatusCreated, res.StatusCode)
	}
}

func TestCreateProduct_InvalidRequest(t *testing.T) {
	db, _, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Couldn't start mock db. %s", err)
	}
	defer db.Close()
	app := fiber.New()
	app.Post("/products", CreateProduct(sqlx.NewDb(db, "mysql")))

	req := httptest.NewRequest("POST", "/products", strings.NewReader(`{
	"name": "test product",
"description": 123
}`))
	res, err := app.Test(req)
	if err != nil {
		t.Fatalf("Error on faking request. %s", err)
	}
	if res.StatusCode != http.StatusBadRequest {
		t.Errorf("Status code isn't %d, got %d instead", http.StatusBadRequest, res.StatusCode)
	}
}

func TestCreateProduct_DBError(t *testing.T) {
	db, mockDb, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Couldn't start mock db. %s", err)
	}
	defer db.Close()
	app := fiber.New()
	app.Post("/products", CreateProduct(sqlx.NewDb(db, "mysql")))

	mockDb.ExpectExec("INSERT INTO `products`").WithArgs("test product", "It is a test product", sqlmock.AnyArg()).WillReturnError(errors.New("db error"))
	req := httptest.NewRequest("POST", "/products", strings.NewReader(`{
	"name": "test product"
}`))
	res, err := app.Test(req)
	if err != nil {
		t.Fatalf("Error on faking request. %s", err)
	}
	if res.StatusCode != http.StatusInternalServerError {
		t.Errorf("Status code isn't %d, got %d instead", http.StatusInternalServerError, res.StatusCode)
	}
}
