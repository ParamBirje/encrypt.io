package tests

import (
	"testing"

	"github.com/parambirje/encrypt.io/internal/validation"
)

var testcases = []struct {
	name      string
	password  string
	password2 string
	want      bool
}{
	{name: "passwords do not match", password: "test", password2: "test1234", want: false},
	{name: "passwords match", password: "test", password2: "test", want: true},
	{name: "passwords empty", password: "", password2: "", want: true},
}

func TestValidatePassword(t *testing.T) {
	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			got := validation.ValidatePassword([]byte(testcase.password), []byte(testcase.password2))
			want := testcase.want

			if got != want {
				t.Errorf("Want %v, got %v", want, got)
			}
		})
	}
}
