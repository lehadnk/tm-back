package domain

import "os/exec"

type CliRunner struct {
}

func (runner CliRunner) Run(command string, args []string) (string, error) {
	stdout, stderr := exec.Command(command, args...).Output()
	return string(stdout), stderr
}
