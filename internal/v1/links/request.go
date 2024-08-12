package links

import (
	"encoding/json"
	"errors"
	"github.com/gofiber/fiber/v2"
	"time"
)

type Request struct {
	Link    string    `json:"links" swaggertype:"string" example:"https://www.example.com" binding:"required"`
	TimeNow time.Time `swaggerignore:"true"`
}

func biuldDTO(c *fiber.Ctx) (Request, error) {
	var dto Request

	err := json.Unmarshal(c.Body(), &dto)
	if err != nil || dto.Link == "" {
		return dto, errors.New("links não informado")
	}

	dto.TimeNow = time.Now()

	return dto, nil
}
