package domain

type MockCliRunner struct {
}

func (runner MockCliRunner) Run(command string, args []string) (string, error) {
	var stdout string
	// runner code
	return stdout, nil
}

func (runner MockCliRunner) On(command string, stdout string, stderr string) {
	// mock setup code
}
