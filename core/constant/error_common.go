package constant

import "errors"

var (
	ErrFailedToConvertStructToBytesF = "failed to convert struct to bytes : %v"
	ErrFailedToConvertBytesToStructF = "failed to convert bytes to struct : %v"
	FailedToConvertBytesToStruct     = "failed to convert bytes to struct"
	ErrTimeMustGreaterThanF          = "time must greater than %v"
	DataNotFound                     = "data not found"

	MessageDeclinedAlreadyExists = "declined: message with this request_id already exists"

	ErrError        = errors.New("error")
	ErrDataNotFound = errors.New(DataNotFound)

	ErrFailedToConvertStructToBytes = errors.New("failed to convert struct to bytes")
	ErrFailedToConvertBytesToStruct = errors.New(FailedToConvertBytesToStruct)

	ErrMessageDeclinedAlreadyExists = errors.New(MessageDeclinedAlreadyExists)
)
