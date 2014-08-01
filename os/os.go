package os

import (
	`github.com/Hellyna/go-util/fmt`
	`os`
)

func Die(exit_status int, a ...interface{}) {
	fmt.Printerrln(a...)
	os.Exit(exit_status)
}

func Dief(exit_status int, format string, a ...interface{}) {
	fmt.Printerrf(format, a...)
	os.Exit(exit_status)
}
/*
func DieFIfErr(err error, exit_status int, format string, a ...interface{}) {
	if err != nil {
		Dief(exit_status, format, a...)
	}
}

func DieIfErr(err error, exit_status int, a...interface{}) {
	if err != nil {
		Die(exit_status, a...)
	}
}
*/

// vim:set ts=4 sw=4 noet:
