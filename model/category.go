package model

import "time"

type Category struct {
	ID        int64  `gorm:"primary_key;column:id;autoIncrement"`
	Name      string `gorm:"column:name"`
	CreatedAt time.Time
	UpdatedAt time.Time
}