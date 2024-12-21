package entity

// Client domain model / entity
type (
	User struct {
		ID   string
		Name string
		Age  int
		City string

		NeedBySystem
	}
)
