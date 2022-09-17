package routes

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"github.com/software-advice/aries-team-assessment/internal/platform/mysql"
	"github.com/software-advice/aries-team-assessment/internal/products/creation"
	"github.com/software-advice/aries-team-assessment/internal/products/listing"
	"io"
	"net/http"
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
	repository := mysql.NewProductRepository(sqlx.NewDb(db, "mysql"))
	service := listing.BuildService(repository)

	app := fiber.New()
	app.Get("/products", GetAllProducts(service))

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

func TestCreateProductK(t *testing.T) {
	type mockDB struct {
		useMock           bool
		expectName        string
		expectDescription string
		mockLastInsertID  int64
		mockRowsAffected  int64
		mockErr           error
	}
	type testCase struct {
		reqBody        []byte
		mockDB         mockDB
		wantStatusCode int
	}
	cases := map[string]testCase{
		"success": {
			reqBody: []byte(`{
	"name": "test product",
	"description": "It is a test product"
}`),
			mockDB: mockDB{
				useMock:           true,
				expectName:        "test product",
				expectDescription: "It is a test product",
				mockLastInsertID:  2,
				mockRowsAffected:  1,
				mockErr:           nil,
			},
			wantStatusCode: http.StatusCreated,
		},
		"invalid body": {
			reqBody:        []byte(""),
			mockDB:         mockDB{},
			wantStatusCode: http.StatusBadRequest,
		},
		"empty name": {
			reqBody: []byte(`{
	"name": "",
	"description": "It is a test product"
}`),
			mockDB:         mockDB{},
			wantStatusCode: http.StatusBadRequest,
		},
		"empty description": {
			reqBody: []byte(`{
	"name": "test product",
	"description": ""
}`),
			mockDB:         mockDB{},
			wantStatusCode: http.StatusBadRequest,
		},
		"DB error": {
			reqBody: []byte(`{
	"name": "test product",
	"description": "It is a test product"
}`),
			mockDB: mockDB{
				useMock:           true,
				expectName:        "test product",
				expectDescription: "It is a test product",
				mockLastInsertID:  0,
				mockRowsAffected:  0,
				mockErr:           errors.New("db err"),
			},
			wantStatusCode: http.StatusInternalServerError,
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			db, mockDb, err := sqlmock.New()
			if err != nil {
				t.Fatalf("Couldn't start mock db. %s", err)
			}
			defer db.Close()
			repository := mysql.NewProductRepository(sqlx.NewDb(db, "mysql"))
			service := creation.BuildService(repository)
			app := fiber.New()
			app.Post("/products", CreateProduct(service))

			if tc.mockDB.useMock {
				mockDb.ExpectExec("INSERT INTO `products`").
					WithArgs(tc.mockDB.expectName, tc.mockDB.expectDescription, sqlmock.AnyArg()).
					WillReturnResult(sqlmock.NewResult(tc.mockDB.mockLastInsertID, tc.mockDB.mockRowsAffected)).
					WillReturnError(tc.mockDB.mockErr)
			}
			req := httptest.NewRequest("POST", "/products", bytes.NewReader(tc.reqBody))
			res, err := app.Test(req)
			if err != nil {
				t.Fatalf("Error on faking request. %s", err)
			}
			if res.StatusCode != tc.wantStatusCode {
				t.Errorf("Status code isn't %d, got %d instead", tc.wantStatusCode, res.StatusCode)
			}
		})
	}
}
