package domain

import "log"

type MockCliRunner struct {
	mocks       []*CliCommandMock
	RunCommands []string
}

type CliCommandMock struct {
	command string
	args    []string
	stdout  string
	stderr  error
}

func (runner *MockCliRunner) Run(command string, args []string) (string, error) {
	var mock *CliCommandMock
	for i := 0; i < len(runner.mocks); i++ {
		if runner.mocks[i].command != command {
			continue
		}

		if len(runner.mocks[i].args) == 0 {
			mock = runner.mocks[i]
			break
		}

		for j := 0; j < len(runner.mocks[i].args); j++ {
			if j > len(args) {
				break
			}

			if runner.mocks[i].args[j] != args[j] {
				break
			}

			if j == len(runner.mocks[i].args)-1 {
				mock = runner.mocks[i]
			}
		}
	}

	for i := 0; i < len(args); i++ {
		command += " " + args[i]
	}

	if mock == nil {
		log.Fatalln("Not mocked: " + command)
	} else {
		log.Println("Mocked: " + command)
	}

	runner.RunCommands = append(runner.RunCommands, command)

	return mock.stdout, mock.stderr
}

func (runner *MockCliRunner) WasCommandRun(command string) bool {
	for i := 0; i < len(runner.RunCommands); i++ {
		if runner.RunCommands[i] == command {
			return true
		}
	}

	return false
}

func (runner *MockCliRunner) On(command string, args []string, stdout string, stderr error) {
	runner.mocks = append(runner.mocks, &CliCommandMock{command: command, args: args, stdout: stdout, stderr: stderr})
}
