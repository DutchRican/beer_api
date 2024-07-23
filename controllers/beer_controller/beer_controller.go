package beer_controller

import (
	"net/http"

	"github.com/dutchrican/beer_api/models"
	"github.com/dutchrican/beer_api/service"
	"github.com/georgysavva/scany/sqlscan"
	"github.com/gin-gonic/gin"
	"github.com/imdario/mergo"
	"github.com/lib/pq"
)

func IndexHandler(c *gin.Context, db service.DB) {
	var beers []models.Beer
	err := sqlscan.Select(c, db.Db, &beers, `SELECT
    beers.id,
    beers.beer_name,
    beers.creator,
    origin_countries.country_name AS origin_country,
    current_countries.country_name AS current_country,
    beers.alcohol
	FROM
		beers
	JOIN
    	countries AS origin_countries ON beers.origin_country_id = origin_countries.id
	JOIN
    	countries AS current_countries ON beers.current_country_id = current_countries.id;`)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err.(*pq.Error).Message)
		return
	}
	c.IndentedJSON(http.StatusOK, beers)
}

func PostHandler(c *gin.Context, db service.DB) {
	var b models.BeerDTO
	if err := c.BindJSON(&b); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.(*pq.Error).Message})
		return
	}
	if len(b.Creator) == 0 || len(b.Beername) == 0 {
		c.IndentedJSON(http.StatusUnprocessableEntity, gin.H{"error": "creator and beer_name must be provided"})
		return
	}
	// err :=
	stmt := `INSERT INTO public.beers (beer_name, creator, origin_country_id, current_country_id, alcohol) values ($1, $2, $3, $4, $5)`
	_, err := db.Db.Exec(stmt, b.Beername, b.Creator, b.OriginCountryId, b.CurrentCountryId, b.Alcohol)
	if err != nil {
		switch err.(*pq.Error).Code {
		case "23502", "23503":
			c.IndentedJSON(http.StatusUnprocessableEntity, gin.H{"error": "Cannot process entity"})
			return
		case "23505":
			c.IndentedJSON(http.StatusConflict, gin.H{"error": "duplicate entry"})
			return
		default:
			c.IndentedJSON(http.StatusConflict, gin.H{"error": err.(*pq.Error).Code})
			return
		}
	}
	c.IndentedJSON(http.StatusCreated, &b)
}

func PutHandler(c *gin.Context, db service.DB) {
	var b models.BeerDTO
	var oldBeer []models.Beer
	if err := c.BindJSON(&b); err != nil {
		return
	}
	err := sqlscan.Select(c, db.Db, &oldBeer, `SELECT * from public.beers WHERE id = $1;`, b.ID)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.(*pq.Error).Message})
		return
	}

	if len(oldBeer) == 0 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "id not found"})
		return
	}
	err = mergo.Merge(&b, oldBeer[0], mergo.WithOverrideEmptySlice)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.(*pq.Error).Message})
		return
	}

	stmt := `
	UPDATE public.beers 
	SET beer_name = $1, creator = $2, origin_country_id = $3, current_country_id = $4, alcohol = $5
	WHERE id = $6;
	`
	_, err = db.Db.Exec(stmt, b.Beername, b.Creator,
		b.OriginCountryId, b.CurrentCountryId, b.Alcohol, b.ID)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}

	c.IndentedJSON(http.StatusOK, &b)
}
func DeleteHandler(c *gin.Context, db service.DB) {
	var b models.Beer
	if err := c.BindJSON(&b); err != nil {
		return
	}
	stmt := `DELETE from public.beers where id = $1;`
	item, err := db.Db.Exec(stmt, b.ID)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, err)
		return
	}
	if count, _ := item.RowsAffected(); count == 0 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "no such ID"})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"deleted": b.ID})
}
