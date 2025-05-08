package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"os"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/yashwanthm3012/xerox/client/utils"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())

	app.Post("/upload", func(c *fiber.Ctx) error {
		fileHeader, err := c.FormFile("file")
		if err != nil {
			return c.Status(400).SendString("No file found in form-data")
		}

		printType := c.FormValue("printType")
		bwPages := c.FormValue("bwPages")
		colorPages := c.FormValue("colorPages")
		pages := c.FormValue("pages")

		if printType == "" {
			return c.Status(400).SendString("Missing printType")
		}
		if printType == "both" && (bwPages == "" || colorPages == "") {
			return c.Status(400).SendString("Missing bwPages or colorPages for 'both'")
		}
		if (printType == "bw" || printType == "color") && pages == "" {
			return c.Status(400).SendString("Missing pages for 'bw' or 'color'")
		}

		// Generate reference ID
		randomBytes := make([]byte, 2)
		if _, err := rand.Read(randomBytes); err != nil {
			return c.Status(500).SendString("Error generating reference ID")
		}
		refID := hex.EncodeToString(randomBytes)

		// Save file temporarily
		tempDir := "./tmp"
		os.MkdirAll(tempDir, os.ModePerm)
		tempPath := filepath.Join(tempDir, refID+"_"+fileHeader.Filename)
		if err := c.SaveFile(fileHeader, tempPath); err != nil {
			return c.Status(500).SendString("Error saving temp file: " + err.Error())
		}

		pageCount, err := api.PageCountFile(tempPath)
		if err != nil {
			return c.Status(500).SendString("Failed to read PDF: " + err.Error())
		}

		var allPages []int
		if printType == "both" {
			allPages, err = utils.MergePageRanges(bwPages + "," + colorPages)
		} else {
			allPages, err = utils.MergePageRanges(pages)
		}
		if err != nil {
			return c.Status(400).SendString("Invalid page ranges: " + err.Error())
		}
		for _, page := range allPages {
			if page > pageCount || page < 1 {
				return c.Status(400).SendString(fmt.Sprintf("Page %d out of range (PDF has %d pages)", page, pageCount))
			}
		}

		err = utils.ForwardToShopStorage(tempPath, fileHeader.Filename)
		if err != nil {
			return c.Status(500).SendString("Failed to forward to storage: " + err.Error())
		}
		os.Remove(tempPath)

		return c.JSON(fiber.Map{
			"status":      "success",
			"referenceID": refID,
			"printType":   printType,
			"fileName":    fileHeader.Filename,
			"totalPages":  pageCount,
			"bwPages":     bwPages,
			"colorPages":  colorPages,
			"pages":       pages,
		})
	})

	app.Listen(":3000")
}
