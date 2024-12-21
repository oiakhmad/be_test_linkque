package entity

import "time"

type NeedBySystem struct {
	CreatedAt time.Time
	CreatedBy int
	UpdatedAt time.Time
	UpdatedBy int
}
