package output

import (
	"fmt"
	"io"

	"go.mattglei.ch/timber"
)

func Line(w io.Writer, a ...any) {
	_, err := fmt.Fprintln(w, a...)
	if err != nil {
		timber.Error(err, "failed to write", a, "to output")
	}
}

func Linef(w io.Writer, format string, a ...any) {
	_, err := fmt.Fprintf(w, format, a...)
	if err != nil {
		timber.Error(err, "failed to write", a, "to output")
	}
}
