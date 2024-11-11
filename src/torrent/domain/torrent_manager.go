package domain

import (
	"awesomeProject/src/torrent/dto"
	"awesomeProject/src/torrent/persistence"
	"awesomeProject/src/transmission"
)

type TorrentManager struct {
	TorrentDao          *persistence.TorrentDao
	TransmissionService *transmission.TransmissionService
}

func NewTorrentManager(
	torrentDao *persistence.TorrentDao,
	transmissionService *transmission.TransmissionService,
) *TorrentManager {
	var newTorrentManager = TorrentManager{
		torrentDao,
		transmissionService,
	}
	return &newTorrentManager
}

func (torrentManager *TorrentManager) GetTorrentList(sort string, page int, pageSize int) dto.FinalTorrentsList {
	torrentsListFromDB := torrentManager.TorrentDao.GetTorrentsList(sort, page, pageSize)
	torrentsCount := torrentManager.TorrentDao.GetCountOfTorrents()
	torrentsListFromTransmission := torrentManager.TransmissionService.GetTransmissionTorrentList()

	var finalTorrents []*dto.FinalTorrent
	for i := 0; i < len(torrentsListFromDB); i++ {
		for j := 0; j < len(torrentsListFromTransmission); i++ {
			if torrentsListFromDB[i].Name != torrentsListFromTransmission[i].Name {
				continue
			}

			finalTorrent := dto.FinalTorrent{
				Torrent:             torrentsListFromDB[i],
				TransmissionTorrent: torrentsListFromTransmission[j],
			}
			finalTorrents = append(finalTorrents, &finalTorrent)
		}
	}
	finalTorrentsList := dto.FinalTorrentsList{
		FinalTorrentArray: finalTorrents,
		FinalTorrentCount: torrentsCount,
	}
	return finalTorrentsList
}

func (torrentManager *TorrentManager) AddTorrent(torrentFilePath string, outputDirectory string) {
	
	// запись файла на диск
	//вызов трансмишн сервиса для добавления загрузки в клиент
	//сохранение информации о загрузке в базу
}
