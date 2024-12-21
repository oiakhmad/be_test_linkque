package email

import "time"

type (
	ResponseUser struct {
		Name      string    `json:"name" example:"zul" validate:"omitempty"`
		Age       int       `json:"age" example:"username"`
		City      string    `json:"city" example:"password"`
		Create_at time.Time `json:"created_at" example:"email" validate:"omitempty"`
	}
)
