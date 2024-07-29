package country_controller

import (
	"net/http"

	"github.com/dutchrican/beer_api/models"
	"github.com/dutchrican/beer_api/service"
	"github.com/gin-gonic/gin"
)

func IndexHandler(c *gin.Context, db service.DB) {
	country := make([]models.Country, 0)
	db.Db.Find(&country)
	c.IndentedJSON(http.StatusOK, country)
}

func PostHandler(c *gin.Context, db service.DB) {
	var m models.Country
	if err := c.BindJSON(&m); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if len(m.Name) == 0 {
		c.IndentedJSON(http.StatusUnprocessableEntity, gin.H{"error": "country name must be provided"})
		return
	}

	db.Db.Create(&m)
	c.IndentedJSON(http.StatusCreated, &m)
}
