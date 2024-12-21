package validator

import (
	"reflect"
	"testing"
)

func TestGetValidatorMessage(t *testing.T) {
	type args struct {
		err  error
		data interface{}
	}
	tests := []struct {
		name         string
		args         args
		wantMessages []string
	}{
		{
			name: "GetValidatorMessage - sucess",
			args: args{
				data: struct {
					Name string `json:"name" validate:"required"`
					Age  string `json:"age" validate:"numeric"`
				}{
					Age: "age",
				},
			},
			wantMessages: []string{
				"Name is a required field",
				"Age must be a valid numeric value",
			},
		},
	}
	for _, test := range tests {
		validator := GetValidator()

		t.Run(test.name, func(t *testing.T) {
			test.args.err = validator.Struct(test.args.data)
			if gotMessages := GetValidatorMessage(test.args.err); !reflect.DeepEqual(gotMessages, test.wantMessages) {
				t.Errorf("GetValidatorMessage() = %v, want %v", gotMessages, test.wantMessages)
			}
		})
	}
}
