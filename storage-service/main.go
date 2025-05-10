package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Post("/store", func(c *fiber.Ctx) error {
		file, err := c.FormFile("file")
		if err != nil {
			return c.Status(400).SendString("File not found")
		}

		refID := c.FormValue("refID")
		if refID == "" {
			return c.Status(400).SendString("Missing refID")
		}

		os.MkdirAll("./shop_uploads", os.ModePerm)
		savePath := filepath.Join("./shop_uploads", fmt.Sprintf("%s_%s", refID, file.Filename))

		if err := c.SaveFile(file, savePath); err != nil {
			return c.Status(500).SendString("Failed to save file")
		}

		return c.JSON(fiber.Map{
			"status":   "stored",
			"refID":    refID,
			"location": savePath,
		})
	})

	app.Listen(":4000")
}
