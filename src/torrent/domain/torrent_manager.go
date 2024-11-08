package domain

import (
	"awesomeProject/src/torrent/persistence"
	"awesomeProject/src/transmission/transmission_client"
)

type TorrentManager struct {
	TorrentDao          *persistence.TorrentDao
	TransmissionClient  *transmission_client.TransmissionClient
	TransmissionService *transmission_client.TransmissionService
}

func NewTorrentManager(
	torrentDao *persistence.TorrentDao,
	transmissionClient *transmission_client.TransmissionClient,
	transmissionService *transmission_client.TransmissionService,
) *TorrentManager {
	var newTorrentManager = TorrentManager{
		torrentDao,
		transmissionClient,
		transmissionService,
	}
	return &newTorrentManager
}

func (torrentManager *TorrentManager) GetTorrentList(sort string, page int, pageSize int) TorrentsList {
	torrentsListFromDb := torrentManager.TorrentDao.GetListOfTorrents(sort, page, pageSize)
	torrentsCount := torrentManager.TorrentDao.GetCountOfTorrents()
	return TorrentsList{
		torrentsListFromDb,
		torrentsCount,
	}

	//torrentsListFromTransmission := torrentManager.TransmissionClient.GetTorrentList()

}

func (torrentManager *TorrentManager) AddNewTorrent(torrentFilePath string, outputDirectory string) {

	// запись файла на диск
	//вызов трансмишн сервиса для добавления загрузки в клиент
	//сохранение информации о загрузке в базу
}
