package helper

import (
	"be_test_linkque/core/constant"
	"be_test_linkque/core/port/utils"
	"bytes"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"io"
	math "math/rand"
	"regexp"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type helper struct{}

var instance utils.Helper

func GetInstance() utils.Helper {
	if instance != nil {
		return instance
	}

	instance = &helper{}
	return instance
}

func (h *helper) StructToMap(obj interface{}) (newMap map[string]interface{}, err error) {
	data, err := json.Marshal(obj)
	if err != nil {
		return
	}

	err = json.Unmarshal(data, &newMap)
	return
}

func (h *helper) GenerateRandomNumber() uint {
	p, _ := rand.Prime(rand.Reader, 32)
	return uint(p.Uint64())
}

// GetTime godoc
// Get time.now() but with no nanosecond & milisecond
func (h *helper) GetTime() time.Time {
	format := "2006-01-02 15:04:05"
	now := time.Now().String()
	s := strings.Split(now, ".")
	date, _ := time.Parse(format, s[0])
	return date
}

func (h *helper) ValidateTime(datetime string) (int, error) {
	now := time.Now()

	parsed, err := time.Parse(time.RFC3339, datetime)
	if err != nil {
		return 0, err
	}

	delay := int(parsed.UTC().Sub(now.UTC()).Milliseconds())
	if delay < 0 {
		return 0, fmt.Errorf(constant.ErrTimeMustGreaterThanF, now.Format(time.RFC3339))
	}

	return delay, nil
}

func (h *helper) ReadFileBufferToBytes(file io.Reader) ([]byte, error) {
	//Create Empty Buffer
	buffer := bytes.NewBuffer(nil)

	// Copy File Stream to Buffer
	io.Copy(buffer, file)

	return buffer.Bytes(), nil
}

func (h *helper) RemoveDuplicateStr(strSlice []string) []string {
	allKeys := make(map[string]bool)
	list := []string{}
	for _, item := range strSlice {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}

// remove array (slice) element
func (h *helper) RemoveSlice(slice, elements []string) []string {
	out := []string{}
	bucket := map[string]bool{}

	for _, element := range slice {
		if !inSlice(elements, element) && !bucket[element] {
			out = append(out, element)
			bucket[element] = true
		}
	}

	return out
}

func inSlice(slice []string, elem string) bool {
	for _, i := range slice {
		if i == elem {
			return true
		}
	}

	return false
}

func (h *helper) IsValidUUID(uuid string) bool {
	r := regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$")
	return r.MatchString(uuid)
}

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

func (h *helper) GenerateRandomString(n int) string {
	math.Seed(time.Now().UnixNano())

	sb := strings.Builder{}
	sb.Grow(n)
	for i := 0; i < n; i++ {
		sb.WriteByte(charset[math.Intn(len(charset))])
	}
	return sb.String()
}

func (h *helper) ValidateArray(listValue []string, value string) bool {
	for _, v := range listValue {
		if v == value {
			return true
		}
	}
	return false
}

func (h *helper) ConvertDateLocalToIso8601(date time.Time) time.Time {
	inputTime, _ := time.Parse("2006-01-02 15:04:05 -0700 MST", date.String())
	iso8601Time := inputTime.UTC().Add(7 * time.Hour).Format(time.RFC3339)
	layout := "2006-01-02T15:04:05Z"
	resultDate, _ := time.Parse(layout, iso8601Time)
	return resultDate
}

func (h *helper) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func (h *helper) CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
