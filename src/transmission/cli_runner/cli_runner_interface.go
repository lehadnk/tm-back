package cli_runner

type CliRunnerInterface interface {
	Run(command string, args []string) (string, error)
}
