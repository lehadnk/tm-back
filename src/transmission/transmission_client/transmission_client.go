package transmission_client

import (
	"awesomeProject/src/transmission/cli_runner"
	"strconv"
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

// array of tranmission torrents
func (client *TransmissionClient) GetTorrentList() []string {
	var args = []string{"-l"}

	result, _ := client.cli.Run("transmission-remote", args)
	stringResult := separateToLines(result)
	separatedStringResult := parseLine(stringResult)
	return separatedStringResult
}

func (client *TransmissionClient) DeleteTorrent(transmissionTorrentId int) bool {
	var args = []string{"-t", strconv.Itoa(transmissionTorrentId), "--remove-and-delete\n"}

	_, stderr := client.cli.Run("transmission-remote", args)
	if stderr != nil {
		return false
	}
	return true
}

func (client *TransmissionClient) GetFileList(transmissionTorrentId int) bool {
	var args = []string{"-t", strconv.Itoa(transmissionTorrentId), "-f"}

	_, stderr := client.cli.Run("transmission-remote", args)
	if stderr != nil {
		return false
	}
	return true
}
