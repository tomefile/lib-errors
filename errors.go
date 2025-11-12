package liberrors

import (
	"fmt"
	"io"
	"strings"

	libescapes "github.com/bbfh-dev/lib-ansi-escapes"
)

// Common error names
const (
	ERROR_INTERNAL = "Internal Error"
	ERROR_ASSERT   = "Assertion Error"

	ERROR_IO         = "I/O Error"
	ERROR_READING    = "Reading Error"
	ERROR_WRITING    = "Writing Error"
	ERROR_SYNTAX     = "Syntax Error"
	ERROR_FORMATTING = "Formatting Error"
)

type DetailedError struct {
	Name    string
	Details string

	Trace   []TraceItem
	Context string
}

// It is recommended to use [.Print()] and [.GetPrinted()] for printing.
func (err *DetailedError) Error() string {
	return fmt.Sprintf("(%s) %s", err.Name, err.Details)
}

func (err *DetailedError) Print(writer io.Writer) {
	fmt.Fprintf(writer, "[!] %s\n", err.Name)

	if len(err.Trace) > 0 {
		err.Trace[0].PrintRoot(writer)
		for _, item := range err.Trace[1:] {
			item.PrintNested(writer)
		}
	}

	ctx := strings.TrimSpace(err.Context)
	if len(ctx) != 0 {
		fmt.Fprintf(writer, "\n    %s\n", ctx)
	}

	fmt.Fprintf(
		writer,
		"\n[?] Details\n    %s%s%s\n",
		libescapes.TextColorBrightRed,
		err.Details,
		libescapes.ColorReset,
	)
}

func (err *DetailedError) GetPrinted() string {
	var builder strings.Builder
	err.Print(&builder)
	return builder.String()
}

func (err *DetailedError) AddTraceItem(item TraceItem) *DetailedError {
	err.Trace = append(err.Trace, item)
	return err
}
