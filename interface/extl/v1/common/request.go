package common

import "time"

type (
	NeedBySystem struct {
		CreatedAt time.Time `json:"created_at,omitempty"`
		CreatedBy int       `json:"created_by,omitempty"`
		UpdatedAt time.Time `json:"updated_at,omitempty"`
		UpdatedBy int       `json:"updated_by,omitempty"`
	}
)
