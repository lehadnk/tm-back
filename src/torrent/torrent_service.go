package torrent

import (
	"tm/src/torrent/domain"
	"tm/src/torrent/dto"
)

type TorrentService struct {
	torrentManager *domain.TorrentManager
}

func NewTorrentService(torrentManager *domain.TorrentManager) *TorrentService {
	var newTorrentService = TorrentService{
		torrentManager,
	}
	return &newTorrentService
}

func (torrentService *TorrentService) GetTorrentsList(sort string, page int, pageSize int) dto.AggregatedTorrentDataList {
	return torrentService.torrentManager.GetTorrentsListFromDatabase(sort, page, pageSize)
}

func (torrentService *TorrentService) GetDownloadingTorrentsList() dto.AggregatedTorrentDataList {
	return torrentService.torrentManager.GetDownloadingTorrentsList()
}

func (torrentService *TorrentService) AddTorrent(file []byte) (*dto.TorrentEntity, error, error) {
	return torrentService.torrentManager.AddTorrent(file)
}

func (torrentService *TorrentService) DeleteTorrent(torrentId int) error {
	return torrentService.torrentManager.DeleteTorrent(torrentId)
}
