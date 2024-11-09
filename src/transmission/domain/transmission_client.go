package domain

import (
	"awesomeProject/src/cli/domain"
	"awesomeProject/src/transmission/dto"
	"strconv"
)

type TransmissionClient struct {
	cli domain.CliRunnerInterface
}

func NewTransmissionClient(runner domain.CliRunnerInterface) *TransmissionClient {
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

func (client *TransmissionClient) GetTorrentsList() []*dto.TransmissionTorrent {
	var args = []string{"-l"}
	parser := TransmissionParser{}

	result, _ := client.cli.Run("transmission-remote", args)
	return parser.SeparateToLines(result)
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
