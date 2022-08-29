package main

import (
	"fmt"
	"log"
	"os"

	. "github.com/dutchrican/beer_api/controllers"
	"github.com/dutchrican/beer_api/service"
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
	port := getEnvVariables("PORT")
	app := fiber.New()
	app.Use(logger.New())

	db := service.DB{}
	if err := db.Open(dbOptions()); err != nil {
		log.Fatal(err)
	}

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
