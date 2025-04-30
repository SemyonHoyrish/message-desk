package v0

import (
	"github.com/SemyonHoyrish/message-desk/server/storage"
	"github.com/gofiber/fiber/v2"
	"time"
)

// GetMessagesHandler handles /api/v0/messages
func GetMessagesHandler(c *fiber.Ctx) error {
	msgs := storage.ReadMessages()
	return c.Status(200).JSON(msgs)
}

type sendMessageRequest struct {
	Content string `json:"content"`
}

// SendMessageHandler handles /api/v0/send
func SendMessageHandler(c *fiber.Ctx) error {
	req := sendMessageRequest{}
	err := c.BodyParser(&req)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "bad json"})
	}
	storage.WriteMessage(storage.Message{
		Content: req.Content,
		Time:    time.Now(),
	})
	return c.Status(200).JSON(fiber.Map{"message": "ok"})
}
