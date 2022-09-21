package routes

import (
	"context"
	"encoding/json"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"github.com/software-advice/aries-team-assessment/internal/platform/jwt"
	mocked "github.com/software-advice/aries-team-assessment/internal/platform/mockable"
	"github.com/software-advice/aries-team-assessment/internal/platform/mysql"
	"github.com/software-advice/aries-team-assessment/internal/users"
	"github.com/software-advice/aries-team-assessment/internal/users/login"
	"github.com/software-advice/aries-team-assessment/internal/users/signup"
	"github.com/software-advice/aries-team-assessment/internal/users/tokenrenew"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestLogin(t *testing.T) {
	db, mockDb, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()
	repo := mysql.NewUsersRepository(sqlx.NewDb(db, "mysql"))
	testKey := []byte("testKey")
	tknGen, err := jwt.BuildHS256Manager(testKey)
	require.NoError(t, err)
	tknGenSrvc := users.BuildTokenGenerationService(tknGen, time.Second)
	service := login.BuildService(repo, tknGenSrvc)

	app := fiber.New()
	app.Post("/users/login", Login(service))

	username := "test"
	hash := "$2a$10$Vq8Tx8eLAFevAULXWtfJXOFFh6eMAMgJ4rQwPett62hO6.6zCJ9eW"
	mockDb.ExpectQuery("SELECT *").
		WithArgs(username).
		WillReturnRows(sqlmock.
			NewRows([]string{"id", "username", "password_hash", "created_at"}).
			AddRow(1, username, hash, time.Now()))

	req := httptest.NewRequest(http.MethodPost, "/users/login", strings.NewReader(`{
	"username": "test",
	"password": "asd123"
}`))
	res, err := app.Test(req)
	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	rawResBody, _ := io.ReadAll(res.Body)
	var parsedRes TokenResponse
	err = json.Unmarshal(rawResBody, &parsedRes)
	require.NoError(t, err)
	tkn := parsedRes.Token
	assert.NotEmpty(t, tkn)
	assert.NoError(t, mockDb.ExpectationsWereMet())
}

func TestRenewToken(t *testing.T) {
	mockTokenGen := new(mocked.TokenGenerator)
	service := tokenrenew.BuildService(users.BuildTokenGenerationService(mockTokenGen, time.Second))

	app := fiber.New()
	testUsername := "testUsername"
	testClaims := users.BuildClaims(users.ParseUnsafeUsername(testUsername), time.Now())
	var fakeVerifiedTkn fiber.Handler = func(ctx *fiber.Ctx) error {
		ctx.SetUserContext(context.WithValue(ctx.UserContext(), ctxClaimsKey, testClaims))
		return ctx.Next()
	}
	app.Post("/users/token/renew", fakeVerifiedTkn, RenewToken(service))
	req := httptest.NewRequest(http.MethodPost, "/users/token/renew", nil)

	testTkn := "test-token"
	mockTokenGen.On("Generate", mock.MatchedBy(func(claims users.Claims) bool {
		return claims.Username().String() == testUsername &&
			claims.ExpiresAt().After(time.Now())
	})).Return(users.ParseTokenString(testTkn), nil)

	res, err := app.Test(req)
	require.NoError(t, err)
	var parsedBody TokenResponse
	bodyBytes, err := io.ReadAll(res.Body)
	require.NoError(t, err)
	err = json.Unmarshal(bodyBytes, &parsedBody)
	require.NoError(t, err)
	assert.Equal(t, testTkn, parsedBody.Token)
}

func TestSignUp(t *testing.T) {
	type testCase struct {
		wantStatus int
		userExists bool
		canSave    bool
	}
	cases := map[string]testCase{
		"user create": {
			wantStatus: http.StatusCreated,
			canSave:    true,
		},
		"user already exists": {
			wantStatus: http.StatusConflict,
			userExists: true,
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			db, mockDb, err := sqlmock.New()
			require.NoError(t, err)
			defer db.Close()
			repo := mysql.NewUsersRepository(sqlx.NewDb(db, "mysql"))
			require.NoError(t, err)
			service := signup.BuildService(repo)

			app := fiber.New()
			app.Post("/users/signup", SignUp(service))

			username := "test"
			rows := sqlmock.
				NewRows([]string{"id", "username", "password_hash", "created_at"})
			if tc.userExists {
				rows.AddRow(1, username, []byte("$some.hash"), time.Now())
			}
			mockDb.ExpectQuery("SELECT *").
				WithArgs(username).
				WillReturnRows(rows)
			if tc.canSave {

				mockDb.ExpectExec("INSERT INTO `users`").
					WithArgs(username, sqlmock.AnyArg()).
					WillReturnResult(sqlmock.NewResult(1, 1))
			}

			req := httptest.NewRequest(http.MethodPost, "/users/signup", strings.NewReader(`{
	"username": "test",
	"password": "asd123"
}`))
			res, err := app.Test(req, int(time.Hour))
			require.NoError(t, err)
			assert.Equal(t, tc.wantStatus, res.StatusCode)
			assert.NoError(t, mockDb.ExpectationsWereMet())
		})
	}

}
