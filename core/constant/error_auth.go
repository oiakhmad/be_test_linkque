package constant

import "errors"

var (
	InvalidCredentials   = "invalid credentials"
	WrongUsedCredentials = "wrong used credentials"
	BasicAuthNotProvided = "authorization Basic auth not provided"
	BasicAuthNotDecoded  = "basic auth not decoded. Please provide a valid token"

	ErrInvalidCredentials   = errors.New(InvalidCredentials)
	ErrWrongUsedCredentials = errors.New(WrongUsedCredentials)
	ErrBasicAuthNotProvided = errors.New(BasicAuthNotProvided)
	ErrBasicAuthNotDecoded  = errors.New(BasicAuthNotDecoded)
)
