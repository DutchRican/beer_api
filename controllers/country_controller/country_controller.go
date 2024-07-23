package country_controller

import (
	"net/http"

	"github.com/dutchrican/beer_api/models"
	"github.com/dutchrican/beer_api/service"
	"github.com/georgysavva/scany/sqlscan"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

func IndexHandler(c *gin.Context, db service.DB) {
	var country []models.Country
	err := sqlscan.Select(c, db.Db, &country, `SELECT * FROM public.countries`)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}
	c.IndentedJSON(http.StatusOK, country)
}

func PostHandler(c *gin.Context, db service.DB) {
	var m models.Country
	if err := c.BindJSON(&m); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.(*pq.Error).Message})
		return
	}
	if len(m.Name) == 0 {
		c.IndentedJSON(http.StatusUnprocessableEntity, gin.H{"error": "country name must be provided"})
		return
	}

	stmt := `INSERT INTO public.countries (country_name) values ($1)`
	_, err := db.Db.Exec(stmt, m.Name)
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
	c.IndentedJSON(http.StatusCreated, &m)
}
