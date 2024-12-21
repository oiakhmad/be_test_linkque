package helper_test

import (
	"be_test_linkque/core/port/utils"
	"be_test_linkque/utils/helper"
	"bytes"
	"io"
	"os"
	"reflect"
	"strings"
	"testing"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Helper", func() {
	var helperUtil utils.Helper

	BeforeEach(func() {
		helperUtil = helper.GetInstance()
	})

	Describe("StructToMap", func() {
		type myStruct struct {
			Name string
		}

		It("should convert struct to map successfully", func() {
			dataStruct := myStruct{"Adnin"}
			dataMap, err := helperUtil.StructToMap(&dataStruct)
			Expect(err).To(BeNil())
			Expect(dataMap["Name"]).To(Equal(dataStruct.Name))
		})

		It("should failed convert struct to map", func() {
			errString := "json: cannot unmarshal string into Go value of type map[string]interface {}"
			dataStruct := "adnin"
			_, err := helperUtil.StructToMap(&dataStruct)
			Expect(err).ToNot(BeNil())
			Expect(err.Error()).To(Equal(errString))
		})

		It("should failed convert struct to map - invalid obj to Marshal", func() {
			errString := "json: unsupported type: chan int"
			_, err := helperUtil.StructToMap(make(chan int))
			Expect(err).ToNot(BeNil())
			Expect(err.Error()).To(Equal(errString))
		})
	})

	Describe("GenerateRandomNumber", func() {
		It("should generate random number successfully", func() {
			number := helperUtil.GenerateRandomNumber()
			Expect(number).ToNot(BeNil())
			Expect(number).ToNot(Equal(0))
		})
	})

	Describe("GetTime", func() {
		It("should get time with no ns & ms", func() {
			time := helperUtil.GetTime()
			Expect(time).ToNot(BeNil())
		})
	})

	Describe("ValidateTime", func() {
		It("ValidateTime - success", func() {
			delay, err := helperUtil.ValidateTime(time.Now().Add(time.Minute).Format(time.RFC3339))
			Expect(err).To(BeNil())
			Expect(delay != 0).To(BeTrue())
		})

		It("ValidateTime - error time is lower than now ", func() {
			delay, err := helperUtil.ValidateTime(time.Now().Format(time.RFC3339))
			Expect(err).ToNot(BeNil())
			Expect(strings.Contains(err.Error(), "time must greater")).To(BeTrue())
			Expect(delay).To(Equal(0))
		})

		It("ValidateTime - error cannot parse time, format is invalid ", func() {
			delay, err := helperUtil.ValidateTime(time.Now().Add(time.Minute).Format(time.RFC822))
			Expect(err).ToNot(BeNil())
			Expect(strings.Contains(err.Error(), "cannot parse")).To(BeTrue())
			Expect(delay).To(Equal(0))
		})
	})
})

func Test_helper_ConvertFileToBytes(t *testing.T) {
	btes, _ := os.ReadFile("helper.go")
	reader := bytes.NewReader(btes)

	type args struct {
		file io.Reader
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name:    "ReadFileBufferToBytes - success",
			args:    args{reader},
			want:    btes,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := helper.GetInstance()
			got, err := h.ReadFileBufferToBytes(tt.args.file)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadFileBufferToBytes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadFileBufferToBytes() got = %v, want %v", got, tt.want)
			}
		})
	}
}
