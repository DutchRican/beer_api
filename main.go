package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

func indexHandler(c *fiber.Ctx) error {
	return c.SendString("Hi there!")
}
func postHandler(c *fiber.Ctx) error {
	return c.SendString("Hi there!")
}
func putHandler(c *fiber.Ctx) error {
	return c.SendString("Hi there!")
}
func deleteHandler(c *fiber.Ctx) error {
	return c.SendString("Hi there!")
}

func main() {
	app := fiber.New()

	app.Get("/", indexHandler)
	app.Post("/", postHandler)
	app.Put("/beer", putHandler)
	app.Delete("/beer", deleteHandler)

	port := os.Getenv("PORT")

	if port == "" {
		port = "3000"
	}

	log.Fatalln(app.Listen(fmt.Sprintf(":%v", port)))
}
