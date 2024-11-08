package transmission_client

type TransmissionService struct {
	transmissionClient TransmissionClient
}

func NewTransmissionService(transmissionClient TransmissionClient) *TransmissionService {
	var newTransmissionService = TransmissionService{
		transmissionClient,
	}
	return &newTransmissionService
}

func (transmissionService *TransmissionService) AddNewTransmissionTorrent(torrentFilePath string, outputDirectory string) {
	transmissionService.transmissionClient.AddTorrentFile(torrentFilePath, outputDirectory)
}

func (transmissionService *TransmissionService) GetTransmissionTorrentInfo() {

}

func (transmissionService *TransmissionService) DeleteTransmissionTorrent() {

}
