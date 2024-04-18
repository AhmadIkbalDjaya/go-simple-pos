package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Category struct {
	ID        uuid.UUID  	`gorm:"primary_key;column:id"`
	Name      string 			`gorm:"column:name"`
	CreatedAt time.Time		`gorm:"column:created_at;autoCreateTime;"`
	UpdatedAt time.Time		`gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
	Products	[]Product		`grom:"foreignKey:category_id;references:id"`
}

func (category *Category) TableName() string {
	return "categories"
}

func (category *Category) BeforeCreate(db *gorm.DB) error {
	category.ID = uuid.New()
	return nil
}