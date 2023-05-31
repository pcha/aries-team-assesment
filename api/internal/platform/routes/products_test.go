package routes

import (
	"bytes"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"github.com/software-advice/aries-team-assessment/internal/platform/mysql"
	"github.com/software-advice/aries-team-assessment/internal/products/creation"
	"github.com/software-advice/aries-team-assessment/internal/products/listing"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestGetProducts(t *testing.T) {
	now := time.Now()
	type mockDB struct {
		expectArg string
		returnErr bool
		mockRows  [][]driver.Value
	}
	type testCase struct {
		qryStr         string
		mockDB         mockDB
		wantStatusCode int
		wantOkBody     []Product
		wantErrBody    ErrorResponse
	}
	cases := map[string]testCase{
		"ok with term": {
			qryStr: "?q=a",
			mockDB: mockDB{
				expectArg: "%a%",
				returnErr: false,
				mockRows: [][]driver.Value{
					{1, "prod 1", "desc 1", now},
					{2, "prod 2", "desc 2", now},
				},
			},
			wantStatusCode: http.StatusOK,
			wantOkBody: []Product{
				{
					Id:          1,
					Name:        "prod 1",
					Description: "desc 1",
					CreatedAt:   now,
				},
				{
					Id:          2,
					Name:        "prod 2",
					Description: "desc 2",
					CreatedAt:   now,
				},
			},
		},
		"ok without term": {
			qryStr: "",
			mockDB: mockDB{
				expectArg: "",
				returnErr: false,
				mockRows: [][]driver.Value{
					{1, "prod 1", "desc 1", now},
					{2, "prod 2", "desc 2", now},
				},
			},
			wantStatusCode: http.StatusOK,
			wantOkBody: []Product{
				{
					Id:          1,
					Name:        "prod 1",
					Description: "desc 1",
					CreatedAt:   now,
				},
				{
					Id:          2,
					Name:        "prod 2",
					Description: "desc 2",
					CreatedAt:   now,
				},
			},
		},
		"error": {
			mockDB: mockDB{
				returnErr: true,
			},
			wantStatusCode: http.StatusInternalServerError,
			wantErrBody:    internalErrorResponse,
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			db, mockDb, err := sqlmock.New()
			require.NoError(t, err)
			defer db.Close()
			repository := mysql.NewProductRepository(sqlx.NewDb(db, "mysql"))
			service := listing.BuildService(repository)

			app := fiber.New()
			app.Get("/products", GetProducts(service))

			queryMock := mockDb.ExpectQuery("SELECT *")
			if tc.mockDB.expectArg != "" {
				queryMock.WithArgs(tc.mockDB.expectArg, tc.mockDB.expectArg)
			}
			if tc.mockDB.returnErr {
				queryMock.WillReturnError(errors.New("mocked error"))
			} else {
				rows := sqlmock.NewRows([]string{"id", "name", "description", "created_at"})
				for _, row := range tc.mockDB.mockRows {
					rows.AddRow(row...)
				}
				queryMock.WillReturnRows(rows)
			}

			req := httptest.NewRequest("GET", "/products"+tc.qryStr, nil)
			res, _ := app.Test(req, int(time.Hour))
			resBody, _ := io.ReadAll(res.Body)

			assert.Equal(t, tc.wantStatusCode, res.StatusCode)
			if tc.wantOkBody != nil {
				asserEqualProductsResponse(t, tc.wantOkBody, resBody)
			} else {
				var parsedBody ErrorResponse
				err = json.Unmarshal(resBody, &parsedBody)
				require.Equal(t, tc.wantErrBody, parsedBody)
			}
		})
	}
}

func asserEqualProductsResponse(t *testing.T, expected []Product, given []byte) {
	var resJson []Product
	err := json.Unmarshal(given, &resJson)
	require.NoError(t, err)
	require.Equal(t, len(expected), len(resJson))
	for i, product := range resJson {
		assert.Equal(t, expected[i].Id, product.Id)
		assert.Equal(t, expected[i].Name, product.Name)
		assert.Equal(t, expected[i].Description, product.Description)
	}
}

func TestCreateProduct(t *testing.T) {
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
