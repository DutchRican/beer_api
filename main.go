package main

import (
	"log"
	"os"

	beers "github.com/dutchrican/beer_api/controllers/beer_controller"
	country "github.com/dutchrican/beer_api/controllers/country_controller"
	"github.com/dutchrican/beer_api/service"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func getEnvVariables(key string) string {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return os.Getenv(key)
}

func main() {
	// port := getEnvVariables("PORT")
	// app := fiber.New()
	app := gin.Default()

	db := service.DB{}
	if err := db.Open(dbOptions()); err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	group := app.Group("/api/v1")
	group.GET("/beers", func(c *gin.Context) {
		beers.IndexHandler(c, db)
	})
	group.POST("/beer", func(c *gin.Context) {
		beers.PostHandler(c, db)
	})
	group.PUT("/beer", func(c *gin.Context) {
		beers.PutHandler(c, db)
	})
	group.DELETE("/beer", func(c *gin.Context) {
		beers.DeleteHandler(c, db)
	})

	group.GET("/countries", func(c *gin.Context) {
		country.IndexHandler(c, db)
	})
	group.POST("/country", func(c *gin.Context) {
		country.PostHandler(c, db)
	})

	app.Run()
}

func dbOptions() service.ConnectionOptions {
	username := getEnvVariables("USERNAME")
	password := getEnvVariables("PASSWORD")
	db_port := getEnvVariables("DB_PORT")

	return service.ConnectionOptions{
		User:     username,
		Password: password,
		Port:     db_port,
		Host:     getEnvVariables("HOST"),
		Dbname:   getEnvVariables("DB_NAME"),
	}
}
