package response

import (
	"github.com/gofiber/fiber/v2"
	"time"
)

type Error struct {
	Message   string    `json:"error" swaggertype:"string" example:"Error message"`
	TimeStamp time.Time `json:"time_stamp" swaggertype:"string" example:"2024-07-05T21:39:45.141302071-03:00"`
}

func ReturnError(c *fiber.Ctx, status int, e string) error {
	var res Error

	res.Message = e
	res.TimeStamp = time.Now()

	return c.Status(status).JSON(res)
}
