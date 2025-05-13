package domain

import (
	"errors"
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

func (client *TransmissionClient) AddTorrentFileToTransmission(torrentFilePath string, outputDirectory string) (bool, error) {
	var args = []string{"-a", torrentFilePath, "-w", outputDirectory}

	stdout, stderr := client.cli.Run("transmission-remote", args)
	if stderr != nil {
		return false, stderr
	}

	if !strings.Contains(stdout, "success") {
		return false, errors.New("transmission-remote returned with no success response: " + stdout)
	}

	return true, nil
}

func (client *TransmissionClient) GetTransmissionTorrentList() []*dto.TransmissionTorrentData {
	var args = []string{"-l"}
	parser := TransmissionParser{}

	result, _ := client.cli.Run("transmission-remote", args)
	return parser.SeparateToLines(result)
}

func (client *TransmissionClient) DeleteTransmissionTorrent(transmissionTorrentId int) error {
	var args = []string{"-t", strconv.Itoa(transmissionTorrentId), "--remove-and-delete"}

	response, err := client.cli.Run("transmission-remote", args)
	if err != nil {
		return err
	}

	if !strings.Contains(response, "success") {
		return errors.New("Transmission returned incorrect response to delete request: " + response)
	}

	return nil
}

func (client *TransmissionClient) GetTransmissionTorrentByName(name string) *dto.TransmissionTorrentData {
	list := client.GetTransmissionTorrentList()
	for i := 0; i < len(list); i++ {
		if list[i].Name == name {
			return list[i]
		}
	}

	return nil
}
