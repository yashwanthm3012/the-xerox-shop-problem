package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	//Initilize a new fiber app
	app := fiber.New()

	// Simple Route
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello This is yash")
	})

	// Upload Route
	app.Post("/upload", func(c *fiber.Ctx) error {

		// Black and white Pages
		bwPages := c.FormValue("bwPages")

		// Color Pages
		colorPages := c.FormValue("colorPages")

		// Number of Copies
		numCopies := c.FormValue("numCopies")

		return c.JSON(fiber.Map{
			"message":    "upload successfull",
			"bwPages":    bwPages,
			"colorPages": colorPages,
			"copies":     numCopies,
		})
	})

	// Start the server on port 3000
	log.Fatal(app.Listen(":3000"))

}
