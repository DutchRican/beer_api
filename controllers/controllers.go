package controllers

import (
	"encoding/json"
	"log"

	. "github.com/dutchrican/beer_api/models"
	"github.com/dutchrican/beer_api/service"
	"github.com/georgysavva/scany/sqlscan"
	"github.com/gofiber/fiber/v2"
	"github.com/lib/pq"
)

func IndexHandler(c *fiber.Ctx, db service.DB) error {
	var beers []Beer
	err := sqlscan.Select(c.Context(), db.Db, &beers, `SELECT * FROM public.beers`)

	if err != nil {
		log.Fatalln(err)
		c.JSON("Error occurred")
	}
	return c.JSON(beers)
}
func PostHandler(c *fiber.Ctx, db service.DB) error {
	var b Beer
	if err := c.BodyParser(&b); err != nil {
		return err
	}
	if len(b.Creator) == 0 || len(b.Beername) == 0 {
		c.Context().SetStatusCode(fiber.StatusUnprocessableEntity)
		return c.JSON(&fiber.Map{"error": "creator and beer_name must be provided"})
	}
	// err :=
	stmt := `INSERT INTO public.beers (beer_name, creator, origin_country, current_country, alcohol) values ($1, $2, $3, $4, $5)`
	_, err := db.Db.Exec(stmt, b.Beername, b.Creator, b.OriginCountry, b.CurrentCountry, b.Alcohol)
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
func PutHandler(c *fiber.Ctx, db service.DB) error {
	var b Beer
	var oldBeer []Beer
	if err := c.BodyParser(&b); err != nil {
		return err
	}
	err := sqlscan.Select(c.Context(), db.Db, &oldBeer, `SELECT * from public.beers WHERE id = $1;`, b.ID)
	if err != nil {
		c.Context().SetStatusCode(fiber.StatusNotFound)
		return c.JSON(&fiber.Map{"error": err.(*pq.Error).Message})
	}

	if len(oldBeer) == 0 {
		c.Context().SetStatusCode(fiber.StatusNotFound)
		return c.JSON(&fiber.Map{"error": "id not found"})
	}

	err = json.Unmarshal(c.Body(), &oldBeer[0])
	if err != nil {
		return err
	}
	stmt := `
	UPDATE public.beers 
	SET beer_name = $1, creator = $2, origin_country = $3, current_country = $4, alcohol = $5
	WHERE id = $6;
	`
	_, err = db.Db.Exec(stmt, oldBeer[0].Beername, oldBeer[0].Creator,
		oldBeer[0].OriginCountry, oldBeer[0].CurrentCountry, oldBeer[0].Alcohol, oldBeer[0].ID)
	if err != nil {
		return err
	}

	return c.JSON(&oldBeer[0])
}
func DeleteHandler(c *fiber.Ctx, db service.DB) error {
	var b Beer
	if err := c.BodyParser(&b); err != nil {
		return err
	}
	stmt := `DELETE from public.beers where id = $1;`
	item, err := db.Db.Exec(stmt, b.ID)
	if err != nil {
		c.Context().SetStatusCode(fiber.StatusNotFound)
		return c.JSON(&fiber.Map{"error": err.(*pq.Error).Message})
	}
	if count, _ := item.RowsAffected(); count == 0 {
		c.Context().SetStatusCode(fiber.StatusNotFound)
		return c.JSON(&fiber.Map{"error": "no such ID"})
	}
	return c.JSON(&fiber.Map{"deleted": b.ID})
}
