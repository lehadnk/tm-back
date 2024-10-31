package transmission_client

import (
	"awesomeProject/src/transmission/cli_runner"
)

type TransmissionClient struct {
	cli cli_runner.CliRunnerInterface
}

func NewTransmissionClient(runner cli_runner.CliRunnerInterface) *TransmissionClient {
	return &TransmissionClient{cli: runner}
}

func (client *TransmissionClient) AddTorrentFile(torrentFilePath string, outputDirectory string) bool {
	var args = []string{"-a", torrentFilePath, "-w", outputDirectory}

	_, stderr := client.cli.Run("transmission-remote", args)
	if stderr != nil {
		return false
	}
	return true
}

func (client *TransmissionClient) GetTorrentList() {

}

func (client *TransmissionClient) DeleteTorrent(transmissionTorrentId int) {

}

func (client *TransmissionClient) GetFileList(transmissionTorrentId int) {

}
