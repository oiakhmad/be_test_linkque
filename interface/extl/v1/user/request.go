package email

type (
	RequestUserLogin struct {
		Username string `json:"username" validate:"required"`
		Password string `json:"password" validate:"required"`
	}

	Indentify struct {
		InputFormat string `json:"input_format" validate:"required" `
	}
)
