package main

import (
	"log"
	"os"

	"github.com/ASPCODES/ObservAI/tree/main/apps/api/internal/db"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	// here .env is loading

	if err := godotenv.Load(); err != nil {
		log.Println(".env file not found")
	}

	// connecting postgress 

	pg := db.NewPostgres(os.Getenv("DATABASE_URL"))
	_ = pg


	//Starting the server

	app := fiber.New()

	app.Get("/health", func (c *fiber .Ctx) error {
		return c.JSON(fiber.Map{"status": "Okay"})
	})

	port := os.Getenv("PORT")
	log.Printf("server starting on port %s", port)
	if err := app.Listen(":" + port); err != nil {
		log.Fatal(err)
	}
}