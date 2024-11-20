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

func (transmissionService *TransmissionService) AddTransmissionTorrentFile(torrentFilePath string, outputDirectory string) {
	transmissionService.transmissionClient.AddTransmissionTorrentFile(torrentFilePath, outputDirectory)
}

func (transmissionService *TransmissionService) GetTransmissionTorrentList() []*dto.TransmissionTorrent {
	return transmissionService.transmissionClient.GetTransmissionTorrentList()
}

func (transmissionService *TransmissionService) DeleteTransmissionTorrent(transmissionTorrentId int) error {
	return transmissionService.transmissionClient.DeleteTransmissionTorrent(transmissionTorrentId)
}

func (transmissionService *TransmissionService) GetTransmissionTorrentByName(name string) *dto.TransmissionTorrent {
	return transmissionService.transmissionClient.GetTransmissionTorrentByName(name)
}
