package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/pdfcpu/pdfcpu/pkg/api"
)

func main() {
	app := fiber.New()

	app.Use(cors.New())

	app.Post("/upload", func(c *fiber.Ctx) error {
		// Get file
		file, err := c.FormFile("file")
		if err != nil {
			return c.Status(400).SendString("No file found in form-data")
		}

		// Get form values
		printType := c.FormValue("printType")
		bwPages := c.FormValue("bwPages")
		colorPages := c.FormValue("colorPages")
		pages := c.FormValue("pages")

		// Validate printType
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

		// Save file
		os.MkdirAll("./uploads", os.ModePerm)
		savePath := fmt.Sprintf("./uploads/%s_%s", refID, file.Filename)
		if err := c.SaveFile(file, savePath); err != nil {
			return c.Status(500).SendString("Error saving file: " + err.Error())
		}

		// Get PDF page count
		pageCount, err := api.PageCountFile(savePath)
		if err != nil {
			return c.Status(500).SendString("Failed to read PDF: " + err.Error())
		}

		// Validate page numbers
		var allPages []int
		if printType == "both" {
			allPages, err = mergePageRanges(bwPages + "," + colorPages)
		} else {
			allPages, err = mergePageRanges(pages)
		}
		if err != nil {
			return c.Status(400).SendString("Invalid page ranges: " + err.Error())
		}

		// Check if any page exceeds the total page count
		for _, page := range allPages {
			if page > pageCount || page < 1 {
				return c.Status(400).SendString(fmt.Sprintf("Page number %d is out of range (PDF has %d pages)", page, pageCount))
			}
		}

		// Success response
		return c.JSON(fiber.Map{
			"status":      "success",
			"referenceID": refID,
			"printType":   printType,
			"fileName":    file.Filename,
			"totalPages":  pageCount,
			"bwPages":     bwPages,
			"colorPages":  colorPages,
			"pages":       pages,
		})
	})

	app.Listen(":3000")
}

// Parses comma-separated ranges like "1,3-5,7" into []int{1,3,4,5,7}
func mergePageRanges(ranges string) ([]int, error) {
	var result []int
	parts := strings.Split(ranges, ",")
	for _, part := range parts {
		part = strings.TrimSpace(part)
		if part == "" {
			continue
		}
		if strings.Contains(part, "-") {
			bounds := strings.Split(part, "-")
			if len(bounds) != 2 {
				return nil, fmt.Errorf("invalid range: %s", part)
			}
			start, err1 := strconv.Atoi(bounds[0])
			end, err2 := strconv.Atoi(bounds[1])
			if err1 != nil || err2 != nil || start > end {
				return nil, fmt.Errorf("invalid range: %s", part)
			}
			for i := start; i <= end; i++ {
				result = append(result, i)
			}
		} else {
			num, err := strconv.Atoi(part)
			if err != nil {
				return nil, fmt.Errorf("invalid number: %s", part)
			}
			result = append(result, num)
		}
	}
	return result, nil
}
