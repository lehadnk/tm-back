package transmission

import (
	"tm/src/transmission/domain"
	"tm/src/transmission/dto"
)

type TransmissionService struct {
	transmissionClient *domain.TransmissionClient
}

func NewTransmissionService(transmissionClient *domain.TransmissionClient) *TransmissionService {
	var newTransmissionService = TransmissionService{
		transmissionClient,
	}
	return &newTransmissionService
}

func (transmissionService *TransmissionService) AddTorrentFileToTransmission(torrentFilePath string, outputDirectory string) (bool, error) {
	return transmissionService.transmissionClient.AddTorrentFileToTransmission(torrentFilePath, outputDirectory)
}

func (transmissionService *TransmissionService) GetTransmissionTorrentList() []*dto.TransmissionTorrentData {
	return transmissionService.transmissionClient.GetTransmissionTorrentList()
}

func (transmissionService *TransmissionService) DeleteTransmissionTorrent(transmissionTorrentId int) error {
	return transmissionService.transmissionClient.DeleteTransmissionTorrent(transmissionTorrentId)
}

func (transmissionService *TransmissionService) GetTransmissionTorrentByName(name string) *dto.TransmissionTorrentData {
	return transmissionService.transmissionClient.GetTransmissionTorrentByName(name)
}
