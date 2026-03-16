package main

import (
	"log"

	"trim/internal/config"
	"trim/internal/handler"

	"github.com/gin-gonic/gin"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func runMigrations() {

	m, err := migrate.New(
		"file://migrations",
		"postgres://postgres:postgres@localhost:5432/finance?sslmode=disable",
	)

	if err != nil {
		log.Fatal(err)
	}

	err = m.Up()

	if err != nil && err.Error() != "no change" {
		log.Fatal(err)
	}
}

func main() {

	runMigrations()

	config.ConnectDatabase()

	r := gin.Default()

	api := r.Group("/api")

	{
		api.GET("/transactions", handler.GetTransactions)
		api.POST("/transactions", handler.CreateTransaction)
	}

	r.Run(":8080")
}