package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/software-advice/aries-team-assessment/internal/platform/jwt"
	"github.com/software-advice/aries-team-assessment/internal/platform/mysql"
	"github.com/software-advice/aries-team-assessment/internal/platform/routes"
	"github.com/software-advice/aries-team-assessment/internal/products/creation"
	"github.com/software-advice/aries-team-assessment/internal/products/listing"
	"github.com/software-advice/aries-team-assessment/internal/users"
	"github.com/software-advice/aries-team-assessment/internal/users/login"
	"github.com/software-advice/aries-team-assessment/internal/users/signup"
	"github.com/software-advice/aries-team-assessment/internal/users/tokenrenew"
	"github.com/software-advice/aries-team-assessment/internal/users/tokenvalidation"
	"os"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

const envKeyMinutesToTokensExpire = "MINUTES_TO_TOKENS_EXPIRE"

func connectDatabase() *sqlx.DB {
	connString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?multiStatements=true&parseTime=true",
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		"localhost",
		"3306",
		os.Getenv("MYSQL_DATABASE"))

	db, err := sqlx.Connect("mysql", connString)
	if err != nil {
		log.WithError(err).Fatal("Can't connect to db, did you start your local database? Check the README! :D")
	}
	return db
}

func loadEnvVars() {
	err := godotenv.Load()
	if err != nil {
		log.WithError(err).Fatal("Something went wrong loading .env file, did you copy the sample.env file? Check the README!!")
	}
}

func setupTokenManager() jwt.HS256Manager {
	key := []byte(os.Getenv("HS265_KEY"))
	tokenManager, err := jwt.BuildHS256Manager(key)
	if err != nil {
		log.WithError(err).Fatal("Can't build the token manager")
	}
	return tokenManager
}

func getMinutesToTokensExpire() time.Duration {
	minsStr := os.Getenv(envKeyMinutesToTokensExpire)
	minsInt, err := strconv.Atoi(minsStr)
	if err != nil {
		log.WithError(err).Warnf("invalid env var value in key %q. Default used", envKeyMinutesToTokensExpire)
		minsInt = 15
	}
	return time.Duration(minsInt) * time.Minute
}

func main() {
	loadEnvVars()
	db := connectDatabase()
	tokenManager := setupTokenManager()
	usersRepository := mysql.NewUsersRepository(db)
	productsRepository := mysql.NewProductRepository(db)
	userSignUpService := signup.BuildService(usersRepository)
	tokenGenerationService := users.BuildTokenGenerationService(tokenManager, getMinutesToTokensExpire())
	usersLoginService := login.BuildService(usersRepository, tokenGenerationService)
	tokenValidationService := tokenvalidation.BuildService(tokenManager)
	tokenRenewService := tokenrenew.BuildService(tokenGenerationService)
	productCreationService := creation.BuildService(productsRepository)
	productsListingService := listing.BuildService(productsRepository)

	verifyTokenMiddleware := routes.VerifyToken(tokenValidationService)

	app := fiber.New()
	app.Use(cors.New()) //TODO: explicit?
	app.Get("/ping", func(ctx *fiber.Ctx) error {
		return ctx.JSON(fiber.Map{"ping": "pong"})
	})
	app.Post("/users", routes.SignUp(userSignUpService))
	app.Post("/users/login", routes.Login(usersLoginService))
	app.Post("/users/token/renew", verifyTokenMiddleware, routes.RenewToken(tokenRenewService))
	products := app.Group("/products", verifyTokenMiddleware)
	products.Post("/", routes.CreateProduct(productCreationService))
	products.Get("/", routes.GetProducts(productsListingService))
	err := app.Listen(":3000")
	if err != nil {
		log.WithError(err).Fatal("Something went wrong starting server in port 3000")
	}
}
