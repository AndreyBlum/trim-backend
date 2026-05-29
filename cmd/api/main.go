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

func vulnerableHandler(db *sql.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        userID := c.Query("id")
        query := "SELECT * FROM users WHERE id = " + userID // taint flow direto
        rows, err := db.Query(query)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        defer rows.Close()
        c.JSON(http.StatusOK, gin.H{"ok": true})
    }
}

func main() {

	runMigrations()

	config.ConnectDatabase()

	r := routes.SetupRoutes()
	r.Use(cors.Default())

	r.Run(":8080")
}
