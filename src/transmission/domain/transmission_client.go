package domain

import (
	"strconv"
	"strings"
	"tm/src/cli/domain"
	"tm/src/transmission/dto"
)

type TransmissionClient struct {
	cli domain.CliRunnerInterface
}

func NewTransmissionClient(cli domain.CliRunnerInterface) *TransmissionClient {
	return &TransmissionClient{cli: cli}
}

func (client *TransmissionClient) AddTransmissionTorrentFile(torrentFilePath string, outputDirectory string) bool {
	var args = []string{"-a", torrentFilePath, "-w", outputDirectory}

	stdout, stderr := client.cli.Run("transmission-remote", args)
	if stderr != nil {
		return false
	}

	if !strings.Contains(stdout, "success") {
		return false
	}

	return true
}

func (client *TransmissionClient) GetTransmissionTorrentList() []*dto.TransmissionTorrent {
	var args = []string{"-l"}
	parser := TransmissionParser{}

	result, _ := client.cli.Run("transmission-remote", args)
	return parser.SeparateToLines(result)
}

func (client *TransmissionClient) DeleteTransmissionTorrent(transmissionTorrentId int) bool {
	var args = []string{"-t", strconv.Itoa(transmissionTorrentId), "--remove-and-delete\n"}

	_, stderr := client.cli.Run("transmission-remote", args)
	if stderr != nil {
		return false
	}
	return true
}
