package tests

import (
	"testing"

	"github.com/parambirje/encrypt.io/internal/validation"
)

var ValidateFileTestcases = []struct {
	name string
	path string
	want bool
}{
	{name: "file exists", path: "/home/thorlden/GoProjects/encrypt.io/README.md", want: true},
	{name: "file does not exist", path: "/home/thorlden/GoProjects/somefolder/filethatdoesntexist.md", want: false},
	{name: "folder path", path: "/home/thorlden/GoProjects/somefolder/", want: false},
	{name: "empty path", path: "", want: false},
}

func TestValidateFile(t *testing.T) {
	for _, testcase := range ValidateFileTestcases {
		t.Run(testcase.name, func(t *testing.T) {
			got := validation.ValidateFile(testcase.path)
			want := testcase.want

			if got != want {
				t.Errorf("Want %v, got %v", want, got)
			}
		})
	}
}
