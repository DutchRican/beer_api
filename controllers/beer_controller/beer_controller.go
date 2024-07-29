package beer_controller

import (
	"net/http"

	"github.com/dutchrican/beer_api/models"
	"github.com/dutchrican/beer_api/service"
	"github.com/gin-gonic/gin"
)

func IndexHandler(c *gin.Context, db service.DB) {
	beers := make([]models.Beer, 0)
	db.Db.Preload("Country").Find(&beers)

	c.IndentedJSON(http.StatusOK, beers)
}

func PostHandler(c *gin.Context, db service.DB) {
	var b models.Beer
	if err := c.BindJSON(&b); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if len(b.Company) == 0 || len(b.Beername) == 0 {
		c.IndentedJSON(http.StatusUnprocessableEntity, gin.H{"error": "creator and beer_name must be provided"})
		return
	}
	db.Db.Create(&b)
	var createdBeer models.Beer
	db.Db.Preload("Country").First(&createdBeer, b.ID)
	c.IndentedJSON(http.StatusCreated, &createdBeer)
}

func PutHandler(c *gin.Context, db service.DB) {
	var b models.Beer
	if err := c.BindJSON(&b); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// get old beer
	var oldBeer models.Beer
	if err := db.Db.First(&oldBeer, b.ID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid beer ID"})
		return
	}

	// update beer with old an new values
	db.Db.Model(&oldBeer).Updates(&b)

	var updatedBeer models.Beer
	db.Db.Preload("Country").First(&updatedBeer, b.ID)
	c.IndentedJSON(http.StatusOK, &updatedBeer)
}

func DeleteHandler(c *gin.Context, db service.DB) {
	var b models.Beer
	if err := c.BindJSON(&b); err != nil {
		return
	}
	db.Db.Delete(&b)
	c.IndentedJSON(http.StatusOK, gin.H{"deleted": b.ID})
}
