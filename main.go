package main

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
)

var db *sql.DB

func main() {
	// Bind handlers to endpoints
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	app.Get("/add_content", addContentHandler)
	app.Get("/add_user", addUserHandler)
	app.Get("/add_edge", addEdgeHandler)
	app.Get("/recomment_content", recommendContentHandler)

	// Create database connection
	var err error
	db, err = sql.Open("postgres", "host=localhost port=5432 user=postgres password=meric61724 dbname=postgres sslmode=disable")
	if err != nil {
		log("Cannot open database connection: " + err.Error())
	}

	app.Listen(":8080")
}
