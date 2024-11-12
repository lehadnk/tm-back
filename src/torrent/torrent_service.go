package torrent

import "tm/src/torrent/dto"

type TorrentService struct {
}

func NewTorrentService() *TorrentService {
	var newTorrentService = TorrentService{}
	return &newTorrentService
}

func (torrentService *TorrentService) GetTorrentList(sort string, page int, pageSize int) dto.FinalTorrentsList {
	return torrentService.GetTorrentList(sort, page, pageSize)
}

func (torrentService *TorrentService) AddTorrent(torrentFilePath string, outputDirectory string) {
	torrentService.AddTorrent(torrentFilePath, outputDirectory)
}

func (torrentService *TorrentService) DeleteTorrent(torrentId int) {
	torrentService.DeleteTorrent(torrentId)
}
