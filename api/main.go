package main

import (
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"github.com/software-advice/aries-team-assessment/internal/routes"
)

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

func main() {
	loadEnvVars()
	db := connectDatabase()

	app := fiber.New()
	app.Get("/ping", func(ctx *fiber.Ctx) error {
		return ctx.JSON(fiber.Map{"ping": "pong"})
	})
	app.Get("/products", routes.GetAllProducts(db))
	app.Post("/products", routes.CreateProduct(db))
	err := app.Listen(":3000")
	if err != nil {
		log.WithError(err).Fatal("Something went wrong starting server in port 3000")
	}
}
