package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Product struct {
	ID 						uuid.UUID `gorm:"primary_key;column:id"`
	Code 					string 		`gorm:"column:code"`
	Name 					string		`gorm:"column:name"`
	Unit 					string		`gorm:"column:unit"`
	CategoryId 		uuid.UUID	`gorm:"column:category_id"`
	Stock 				int				`gorm:"column:stock"`
	PurchasePrice int64			`gorm:"column:purchase_price"`
	SellingPrice 	int64			`gorm:"column:selling_price"`
	Description 	string		`gorm:"column:description"`
	Photo 				string		`gorm:"column:photo"`
	CreatedAt 		time.Time	`gorm:"column:created_at;autoCreateTime"`
	UpdatedAt 		time.Time	`gorm:"column:updated_at;autoCreateTime;utoUpdateTime"`
	Category 			Category	`gorm:"foreignKey:category_id;references:id"`
}

func (product *Product) TableName() string {
	return "products"
}

func (product *Product) BeforeCreate(db *gorm.DB) error {
	product.ID = uuid.New()
	return nil
}