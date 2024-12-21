package models

import "time"

type (
	NeedBySystem struct {
		CreatedAt time.Time `gorm:"column:created_at"`
		CreatedBy int       `gorm:"column:created_by"`
		UpdatedAt time.Time `gorm:"column:updated_at"`
		UpdatedBy int       `gorm:"column:updated_by"`
	}

	User struct {
		ID   int    `gorm:"column:id"`
		Name string `gorm:"column:name"`
		Age  int    `gorm:"column:age"`
		City string `gorm:"column:city"`
		NeedBySystem
	}
)
