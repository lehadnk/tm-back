package domain

type CliRunnerInterface interface {
	Run(command string, args []string) (string, error)
}
