package liberrors_test

import (
	"os"
	"testing"

	liberrors "github.com/tomefile/lib-errors"
)

func TestErrorsFormat(test *testing.T) {
	_, err := os.ReadFile("this file will not open, so you can see how the error looks like!")
	derr := &liberrors.DetailedError{
		Name:    liberrors.ERROR_IO,
		Details: err.Error(),
		Trace: []liberrors.TraceItem{
			{
				Name: "example_file.txt",
				Col:  1,
				Row:  1,
			},
			{
				Name: "parent_file.md",
				Col:  1,
				Row:  9,
			},
		},
		Context: "",
	}
	derr.Print(test.Output())
}
