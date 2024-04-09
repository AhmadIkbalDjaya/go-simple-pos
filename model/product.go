package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Product struct {
	ID 						uuid.UUID `grom:"type:uuid;primaryKey;column:id"`
	Code 					string 		`grom:"unique"`
	Name 					string
	Unit 					string
	CategoryId 		uuid.UUID
	Stock 				int
	PurchasePrice int64
	SellingPrice 	int64
	Description 	string
	Photo 				string
	CreatedAt 		time.Time
	UpdatedAt 		time.Time
	Category 			Category	`grom:"foreignKey:category_id;references:id"`
}

func (product *Product) BeforeCreate(db *gorm.DB) error {
	product.ID = uuid.New()
	return nil
}