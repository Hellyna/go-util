package exec

import (
	`os`
	`os/exec`
)

func Command(name string, arg ...string) {
	cmd := exec.Command(name, arg...)
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
}

// vim:set ts=4 sw=4 noet: