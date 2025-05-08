package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"os"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
)

// GenerateRandomID generates a unique reference ID
func GenerateRandomID() string {
	randomBytes := make([]byte, 6) // 6 bytes for 12-character hex string
	if _, err := rand.Read(randomBytes); err != nil {
		panic("Failed to generate random ID")
	}
	return hex.EncodeToString(randomBytes)
}

func main() {
	app := fiber.New()

	// Route to handle storing uploaded files
	app.Post("/store", func(c *fiber.Ctx) error {
		// Get the file from the request
		file, err := c.FormFile("file")
		if err != nil {
			return c.Status(400).SendString("File not found")
		}

		// Generate a unique reference ID
		refID := GenerateRandomID()

		// Create the shop_uploads directory if it doesn't exist
		os.MkdirAll("./shop_uploads", os.ModePerm)

		// Create the file path using the reference ID and original file name
		savePath := filepath.Join("./shop_uploads", fmt.Sprintf("%s_%s", refID, file.Filename))

		// Save the file
		if err := c.SaveFile(file, savePath); err != nil {
			return c.Status(500).SendString("Failed to save file")
		}

		// Respond with the reference ID and file path
		return c.JSON(fiber.Map{
			"status":   "stored",
			"refID":    refID,
			"location": savePath,
		})
	})

	// Start the server
	app.Listen(":4000")
}
