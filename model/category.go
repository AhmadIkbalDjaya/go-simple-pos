package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Category struct {
	ID        uuid.UUID  	`gorm:"type:uuid;primary_key;column:id"`
	Name      string 			`gorm:"column:name"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Products	[]Product		`grom:"foreignKey:category_id;references:id"`
}

func (category *Category) BeforeCreate(db *gorm.DB) error {
	category.ID = uuid.New()
	return nil
}