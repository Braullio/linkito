package links

import (
	"fmt"
	"github.com/Braullio/linkito/internal/database"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Link struct {
	gorm.Model `json:"-"`
	ID         uuid.UUID  `gorm:"type:uuid" swaggertype:"string" example:"650a234e-cf54-4e62-97b3-5105dba6ae05"`
	Link       string     `json:"links" swaggertype:"string" example:"https://www.example.com"`
	CreatedAt  time.Time  `json:"created_at" swaggertype:"string" format:"date-time" example:"2024-07-05T21:39:45.141302071-03:00"`
	UpdatedAt  *time.Time `json:"updated_at" swaggertype:"string" format:"date-time" example:"2024-07-05T21:39:45.141302071-03:00"`
	DeletedAt  *time.Time `json:"-"`
}

func (link *Link) Build(dto Request) {
	link.ID = uuid.New()
	link.Link = dto.Link
	link.CreatedAt = dto.TimeNow
}

func (link *Link) Create() error {
	result := database.GetDB().Model(Link{}).Create(link)
	if result.Error != nil {
		return fmt.Errorf("failed to create link: %w", result.Error)
	}
	return nil
}
