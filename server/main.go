package main

import (
	v0 "github.com/SemyonHoyrish/message-desk/server/api/v0"
	"github.com/SemyonHoyrish/message-desk/server/storage"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Warn("Couldn't loading .env file")
	}

	storage.DataFilename = os.Getenv("DATA_FILENAME")

	app := fiber.New()

	app.Use(cors.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("message desk api server")
	})

	api := app.Group("/api")
	apiv0 := api.Group("/v0")

	apiv0.Get("/messages", v0.GetMessagesHandler)
	apiv0.Post("/send", v0.SendMessageHandler)
	//apiv0.Get("/auth", v0.)
	//apiv0.Get("/remove", v0.)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	app.Listen(":" + port)
}
