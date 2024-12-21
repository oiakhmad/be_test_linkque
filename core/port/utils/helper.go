package utils

import (
	"io"
	"time"
)

type Helper interface {
	StructToMap(obj interface{}) (newMap map[string]interface{}, err error)
	GenerateRandomNumber() uint
	GetTime() time.Time
	ValidateTime(datetime string) (int, error)
	ReadFileBufferToBytes(file io.Reader) ([]byte, error)
	RemoveDuplicateStr(strSlice []string) []string
	RemoveSlice(slice, elements []string) []string
	IsValidUUID(uuid string) bool
	GenerateRandomString(n int) string

	//function ini untuk memvalidase	 value apa aja yang diperbolehkan untuk di proses
	ValidateArray(listValue []string, value string) bool
	ConvertDateLocalToIso8601(date time.Time) time.Time
	//auth
	HashPassword(password string) (string, error)
	CheckPasswordHash(password, hash string) bool
}
