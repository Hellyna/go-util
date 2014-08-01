package fmt

import (
	`fmt`
	`os`
)

func Printerrf(format string, a ...interface{}) (n int, err error) {
	return fmt.Fprintf(os.Stderr, format, a...)
}

func Printerrln(a ...interface{}) (n int, err error) {
	return fmt.Fprintln(os.Stderr, a...)
}

// vim:set ts=4 sw=4 noet:
