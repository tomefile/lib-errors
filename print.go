package liberrors

import (
	"io"
	"strings"
)

type CanPrint interface {
	Print(writer io.Writer)
}

func GetPrinted(printable CanPrint) string {
	var builder strings.Builder
	printable.Print(&builder)
	return builder.String()
}
