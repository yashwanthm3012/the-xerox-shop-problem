package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yashwanthm3012/client/db"
	"golang.org/x/crypto/bcrypt"
)

func SetupUserRoutes(app *fiber.App) {

	//Register Route
	app.Post("/register", func(c *fiber.Ctx) error {
		type Request struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}
		var body Request

		if err := c.BodyParser(&body); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid Input"})
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(body.Password), 14)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Password hasing failed"})
		}

		user := db.User{Username: body.Username, Password: string(hashedPassword)}
		if err := db.DB.Create(&user).Error; err != nil {
			return c.Status(400).JSON(fiber.Map{"message": "User already exists or DB error"})
		}

		return c.JSON(fiber.Map{"message": "User registered successfully"})
	})

	// Login Route
	app.Post("/login", func(c *fiber.Ctx) error {
		type Request struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}

		var body Request
		if err := c.BodyParser(&body); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
		}

		var user db.User
		if err := db.DB.Where("username= ?", body.Username).First(&user).Error; err != nil {
			return c.Status(404).JSON(fiber.Map{"error": "User not found"})
		}

		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password)); err != nil {
			return c.Status(401).JSON(fiber.Map{"error": "Incorrect Password"})
		}

		return c.JSON(fiber.Map{"message": "Login successful", "userId": user.ID})
	})
}
