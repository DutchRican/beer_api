package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	. "github.com/dutchrican/beer_api/controllers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func getEnvVariables(key string) string {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return os.Getenv(key)
}

func main() {
	app := fiber.New()
	app.Use(logger.New())
	port := getEnvVariables("PORT")
	username := getEnvVariables("USERNAME")
	password := getEnvVariables("PASSWORD")
	db_port := getEnvVariables("DB_PORT")
	db_ip := getEnvVariables("DB_IP")
	connStr := fmt.Sprintf("postgresql://%s:%s@%s:%s/beers?sslmode=disable", username, password, db_ip, db_port)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	res, err := db.Exec(`CREATE TABLE IF NOT EXISTS beers (
		id SERIAL PRIMARY KEY ,
		beer_name TEXT NOT NULL UNIQUE,
		creator TEXT NOT NULL,
		origin_country TEXT,
		current_country TEXT,
		alcohol NUMERIC(4, 2)
	 );`)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(res)

	app.Get("/", func(c *fiber.Ctx) error {
		return IndexHandler(c, db)
	})
	app.Post("/", func(c *fiber.Ctx) error {
		return PostHandler(c, db)
	})
	app.Put("/beer", func(c *fiber.Ctx) error {
		return PutHandler(c, db)
	})
	app.Delete("/beer", func(c *fiber.Ctx) error {
		return DeleteHandler(c, db)
	})

	log.Fatalln(app.Listen(fmt.Sprintf(":%v", port)))
}
