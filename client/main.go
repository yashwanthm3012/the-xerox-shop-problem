package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	//Initilize a new fiber app
	app := fiber.New()

	//Simple Route
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello This is yash")
	})

	app.Get("/:name", func(c *fiber.Ctx) error {
		name := c.Params("name")
		return c.SendString("Hello" + name)
	})

	//Start the server on port 3000
	log.Fatal(app.Listen(":3000"))

}
