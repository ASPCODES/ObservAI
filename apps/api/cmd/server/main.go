package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/health", func(c *fiber.Ctx) error {
        return c.JSON(fiber.Map{"status": "ok"})
    })

	log.Println("Server starting at port 8080")
	if err := app.Listen(":8080"); err != nil {
		log.Fatal(err)
	}

}