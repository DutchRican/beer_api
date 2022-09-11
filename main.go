package main

import (
	"log"
	"os"

	. "github.com/dutchrican/beer_api/controllers"
	"github.com/dutchrican/beer_api/service"
	"github.com/gin-gonic/gin"
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
	// port := getEnvVariables("PORT")
	// app := fiber.New()
	app := gin.Default()
	// app.Use(logger.New())

	db := service.DB{}
	if err := db.Open(dbOptions()); err != nil {
		log.Fatal(err)
	}

	app.GET("/", func(c *gin.Context) {
		IndexHandler(c, db)
	})
	app.POST("/", func(c *gin.Context) {
		PostHandler(c, db)
	})
	app.PUT("/beer", func(c *gin.Context) {
		PutHandler(c, db)
	})
	app.DELETE("/beer", func(c *gin.Context) {
		DeleteHandler(c, db)
	})

	app.Run()
}

func dbOptions() service.ConnectionOptions {
	username := getEnvVariables("USERNAME")
	password := getEnvVariables("PASSWORD")
	db_port := getEnvVariables("DB_PORT")
	db_ip := getEnvVariables("DB_IP")

	return service.ConnectionOptions{
		Username: username,
		Password: password,
		DB_port:  db_port,
		DP_Ip:    db_ip,
	}
}
