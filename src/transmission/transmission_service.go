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

func (transmissionService *TransmissionService) AddNewTransmissionTorrent(torrentFilePath string, outputDirectory string) {
	transmissionService.transmissionClient.AddTorrentFile(torrentFilePath, outputDirectory)
}

func (transmissionService *TransmissionService) GetTransmissionTorrentList() []*dto.TransmissionTorrent {
	transmissionTorrentList := transmissionService.transmissionClient.GetTorrentsList()
	return transmissionTorrentList
}

func (transmissionService *TransmissionService) GetTransmissionTorrentInfo() {

}

func (transmissionService *TransmissionService) DeleteTransmissionTorrent() {

}
