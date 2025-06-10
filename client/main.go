package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/yashwanthm3012/client/db"
	"github.com/yashwanthm3012/client/routes"
	"github.com/yashwanthm3012/client/utils"
)

func main() {
	db.InitDB() // Connect to MySQL
	//Initilize a new fiber app
	app := fiber.New()

	routes.SetupUserRoutes(app) //routes

	// Simple Route
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello This is yash")
	})

	// Upload Route
	app.Post("/upload", func(c *fiber.Ctx) error {

		// Get the file
		file, err := c.FormFile("document")
		if err != nil {
			return err
		}

		// Make sure that upload directory exists
		uploadPath := "./uploads/" + file.Filename
		err = c.SaveFile(file, uploadPath)
		if err != nil {
			return err
		}

		// Set output path for extracted PDF
		outputPath := "./extracted/extracted_" + file.Filename

		// Get the print type bw, color or both
		printType := c.FormValue("printType")

		// Black and white Pages
		bwPages := c.FormValue("bwPages")

		// Color Pages
		colorPages := c.FormValue("colorPages")

		// Number of Copies
		numCopies := c.FormValue("numCopies")

		//Validate based on printType
		var (
			valid    = true
			errorMsg string
		)

		switch printType {
		case "bw":
			if bwPages == "" {
				valid = false
				errorMsg = "bwPages are required for black and white print"
			} else {
				err := utils.ExtractPages(uploadPath, outputPath, bwPages)
				if err != nil {
					return c.Status(500).SendString("Failed to extract pages" + err.Error())
				}
			}

		case "color":
			if colorPages == "" {
				valid = false
				errorMsg = "colorPages are required for color print"
			} else {
				err := utils.ExtractPages(uploadPath, outputPath, colorPages)
				if err != nil {
					return c.Status(500).SendString("Cannot extract pages" + err.Error())
				}
			}

		case "both":
			if colorPages == "" || bwPages == "" {
				valid = false
				errorMsg = "page ranges required for both color and black and white print"
			}

		default:
			valid = false
			errorMsg = "Print type must be bw, color or both"
		}

		if !valid {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": errorMsg,
			})
		}

		// Number of Pages
		numPages, err := utils.CountPages(uploadPath)
		if err != nil {
			return err
		}

		// Random 4-digit hexadecimal number
		hexNumber, err := utils.Generate4DigitHex()
		if err != nil {
			log.Fatal(err)
		}

		// To extract the specified page numbers
		err = utils.ExtractPages(uploadPath, outputPath, bwPages)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}

		return c.JSON(fiber.Map{
			"message":    "upload successfull",
			"file":       file.Filename,
			"bwPages":    bwPages,
			"colorPages": colorPages,
			"copies":     numCopies,
			"printType":  printType,
			"numPages":   numPages,
			"ID":         hexNumber,
		})
	})

	// Start the server on port 3000
	log.Fatal(app.Listen(":3000"))

}
