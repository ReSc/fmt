package fmt

import (
	"fmt"
	"io"
)

// String wraps fmt.Sprintf()
func String(format string, a ...interface{}) string {
	return fmt.Sprintf(format, a...)
}

// Error wraps fmt.Errorf()
func Error(format string, a ...interface{}) error {
	return fmt.Errorf(format, a...)
}

// Error wraps fmt.Errorf()
func Panic(format string, a ...interface{}) {
	panic(Error(format, a...))
}

// Write wraps fmt.Fprintf()
func Write(w io.Writer, format string, a ...interface{}) (int, error) {
	return fmt.Fprintf(w, format, a...)
}

// Writefln wraps fmt.Fprintf(), and adds a newline
func Writeln(w io.Writer, format string, a ...interface{}) (int, error) {
	count, err := fmt.Fprintf(w, format, a...)
	if err != nil {
		return count, err
	}
	count2, err := fmt.Fprintln(w)
	return count + count2, err
}

type Writer struct {
	w io.Writer
}

func NewWriter(w io.Writer) *Writer {
	return &Writer{w}
}

func (w *Writer) Write(p []byte) (int, error) {
	return w.w.Write(p)
}

// Writefmt wraps fmt.Fprintf()
func (w *Writer) Writefmt(format string, a ...interface{}) (int, error) {
	return Write(w, format, a...)
}

// Writeln wraps fmt.Fprintf(), and adds a newline
func (w *Writer) Writeln(format string, a ...interface{}) (int, error) {
	return Writeln(w, format, a...)
}
