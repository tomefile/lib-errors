package liberrors_test

import (
	"os"
	"testing"

	liberrors "github.com/tomefile/lib-errors"
)

func TestErrorsFull(test *testing.T) {
	_, err := os.ReadFile("/unknown.txt")
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
		Context: liberrors.Context{
			FirstLine:   7,
			Buffer:      "this is some text\n\nand then I do ",
			Highlighted: "something that has an error\nand yet another line!\nhi\nhello\n\nworld!\n",
		},
	}
	derr.Print(test.Output())
}

func TestErrorsNoContext(test *testing.T) {
	_, err := os.ReadFile("/unknown.txt")
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
		Context: liberrors.Context{},
	}
	derr.Print(test.Output())
}

func TestErrorsShortContext(test *testing.T) {
	_, err := os.ReadFile("/unknown.txt")
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
		Context: liberrors.Context{
			FirstLine:   1,
			Buffer:      "this has no highlighted text",
			Highlighted: "",
		},
	}
	derr.Print(test.Output())
}

func TestErrorsOnlyHighlighted(test *testing.T) {
	_, err := os.ReadFile("/unknown.txt")
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
		Context: liberrors.Context{
			FirstLine:   12,
			Buffer:      "",
			Highlighted: "only highlighted text",
		},
	}
	derr.Print(test.Output())
}
