package src

import (
	_ "fmt"
	"os/exec"
)

func Call() (string, error) {
	stdout, stderr := exec.Command("ls", "-l").Output()

	return string(stdout), stderr
}
