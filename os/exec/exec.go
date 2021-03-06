package exec

import (
	`os`
	`os/exec`
)

func Command(name string, arg ...string) *exec.Cmd {
	cmd := exec.Command(name, arg...)
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr

	return cmd
}

func Exec(name string, arg ...string) (*exec.Cmd, error) {
	cmd := Command(name, arg...)
	return cmd, cmd.Run()
}

// vim:set ts=4 sw=4 noet:
