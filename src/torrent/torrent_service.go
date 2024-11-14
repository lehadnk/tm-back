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

func (torrentService *TorrentService) GetTorrentList(sort string, page int, pageSize int) dto.FinalTorrentsList {
	return torrentService.torrentManager.GetTorrentList(sort, page, pageSize)
}

func (torrentService *TorrentService) AddTorrent(file []byte) (*dto.Torrent, error) {
	return torrentService.torrentManager.AddTorrent(file)
}

func (torrentService *TorrentService) DeleteTorrent(torrentId int) {
	torrentService.torrentManager.DeleteTorrent(torrentId)
}
