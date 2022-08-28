package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/georgysavva/scany/sqlscan"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
)

type Beer struct {
	ID             int     `json:"id"`
	Beername       string  `db:"beer_name" json:"beer_name"`
	Creator        string  `db:"creator" json:"creator"`
	OriginCountry  string  `db:"origin_country" json:"origin_country"`
	CurrentCountry string  `db:"current_country" json:"current_country"`
	Alcohol        float32 `db:"alcohol" json:"alcohol"`
}

func indexHandler(c *fiber.Ctx, db *sql.DB) error {
	var beers []Beer
	err := sqlscan.Select(c.Context(), db, &beers, `SELECT * FROM public.beers`)

	if err != nil {
		log.Fatalln(err)
		c.JSON("Error occurred")
	}
	return c.JSON(beers)
}
func postHandler(c *fiber.Ctx, db *sql.DB) error {
	var b Beer
	// var existingB Beer
	if err := c.BodyParser(&b); err != nil {
		return err
	}
	if len(b.Creator) == 0 || len(b.Beername) == 0 {
		c.Context().SetStatusCode(fiber.StatusUnprocessableEntity)
		return c.JSON(&fiber.Map{"error": "creator and beer_name must be provided"})
	}
	// err :=
	stmt := `INSERT INTO public.beers (beer_name, creator, origin_country, current_country, alcohol) values ($1, $2, $3, $4, $5)`
	_, err := db.Exec(stmt, b.Beername, b.Creator, b.OriginCountry, b.CurrentCountry, b.Alcohol)
	if err != nil {
		switch err.(*pq.Error).Code {
		case "23502":
			c.Context().SetStatusCode(fiber.StatusUnprocessableEntity)
		default:
			c.Context().SetStatusCode(fiber.StatusConflict)
		}
		return c.JSON(&fiber.Map{"error": err.(*pq.Error).Message})
	}
	c.Context().SetStatusCode(fiber.StatusCreated)
	return c.JSON(b)
}
func putHandler(c *fiber.Ctx, db *sql.DB) error {
	return c.SendString("Hi there!")
}
func deleteHandler(c *fiber.Ctx, db *sql.DB) error {
	return c.SendString("Hi there!")
}

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
		return indexHandler(c, db)
	})
	app.Post("/", func(c *fiber.Ctx) error {
		return postHandler(c, db)
	})
	app.Put("/beer", func(c *fiber.Ctx) error {
		return putHandler(c, db)
	})
	app.Delete("/beer", func(c *fiber.Ctx) error {
		return deleteHandler(c, db)
	})

	log.Fatalln(app.Listen(fmt.Sprintf(":%v", port)))
}
