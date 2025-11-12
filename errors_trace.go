package liberrors

import (
	"fmt"
	"io"
)

type TraceItem struct {
	Name     string
	Col, Row uint
}

func (item TraceItem) PrintRoot(writer io.Writer) {
	fmt.Fprintf(
		writer,
		"    in %s:%d:%d\n",
		item.Name,
		item.Row,
		item.Col,
	)
}

func (item TraceItem) PrintNested(writer io.Writer) {
	fmt.Fprintf(
		writer,
		"    └─ from %s:%d:%d\n",
		item.Name,
		item.Row,
		item.Col,
	)
}
