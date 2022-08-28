package controllers

import (
	"database/sql"
	"log"

	. "github.com/dutchrican/beer_api/models"
	"github.com/georgysavva/scany/sqlscan"
	"github.com/gofiber/fiber/v2"
	"github.com/lib/pq"
)

func IndexHandler(c *fiber.Ctx, db *sql.DB) error {
	var beers []Beer
	err := sqlscan.Select(c.Context(), db, &beers, `SELECT * FROM public.beers`)

	if err != nil {
		log.Fatalln(err)
		c.JSON("Error occurred")
	}
	return c.JSON(beers)
}
func PostHandler(c *fiber.Ctx, db *sql.DB) error {
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
func PutHandler(c *fiber.Ctx, db *sql.DB) error {
	return c.SendString("Hi there!")
}
func DeleteHandler(c *fiber.Ctx, db *sql.DB) error {
	return c.SendString("Hi there!")
}
