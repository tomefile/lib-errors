package liberrors

import (
	"fmt"
	"io"
	"strings"

	libescapes "github.com/bbfh-dev/lib-ansi-escapes"
)

type Context struct {
	FirstLine   uint
	Buffer      string
	Highlighted string
}

func (ctx Context) IsEmpty() bool {
	return len(ctx.Buffer) == 0 && len(ctx.Highlighted) == 0
}

func (ctx Context) Print(writer io.Writer) {
	writer.Write([]byte("\n" + libescapes.TextColorWhite))
	var line_index uint

	for line := range strings.SplitSeq(ctx.Buffer, "\n") {
		if line_index != 0 {
			writer.Write([]byte{'\n'})
		}
		fmt.Fprintf(
			writer,
			"%5d |  %s",
			ctx.FirstLine+line_index,
			line,
		)
		line_index++
	}

	writer.Write([]byte(libescapes.TextColorBrightRed))

	if len(ctx.Highlighted) == 0 {
		writer.Write([]byte("←—"))
	} else {
		i := uint(0)
		buffer := strings.TrimSuffix(ctx.Highlighted, "\n")
		for line := range strings.SplitSeq(buffer, "\n") {
			if i == 0 {
				writer.Write([]byte(line))
			} else {
				fmt.Fprintf(
					writer,
					"\n%5d |  %s",
					ctx.FirstLine+line_index+i-1,
					line,
				)
			}
			i++
		}

	}

	writer.Write([]byte(libescapes.ColorReset + "\n"))
}
