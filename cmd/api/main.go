package main

import (
	"log"

	"trim/internal/config"
	"trim/internal/routes"

	"github.com/gin-contrib/cors"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func runMigrations() {

	m, err := migrate.New(
		"file://migrations",
		"postgres://postgres:postgres@localhost:5432/trim?sslmode=disable",
	)

	if err != nil {
		log.Fatal(err)
	}

	err = m.Up()

	if err != nil && err.Error() != "no change" {
		log.Fatal(err)
	}

	log.Default().Println("Migrations applied successfully")
}

func main() {

	runMigrations()

	config.ConnectDatabase()

	r := routes.SetupRoutes()
	r.Use(cors.Default())

	r.Run(":8080")
}
