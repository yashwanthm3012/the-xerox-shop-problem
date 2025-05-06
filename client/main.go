package main

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Post("/upload", func(c *fiber.Ctx) error {
		// Get file from form-data key "file"
		file, err := c.FormFile("file")
		if err != nil {
			return c.Status(400).SendString("No file found in form-data")
		}

		// Create uploads directory if it doesn't exist
		os.MkdirAll("./uploads", os.ModePerm)

		// Save file
		savePath := fmt.Sprintf("./uploads/%s", file.Filename)
		if err := c.SaveFile(file, savePath); err != nil {
			return c.Status(500).SendString("Error saving file: " + err.Error())
		}

		return c.SendString("File uploaded successfully")
	})

	app.Listen(":3000")
}
