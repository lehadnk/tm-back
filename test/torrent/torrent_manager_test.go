package torrent

import (
	"awesomeProject/src/torrent/domain"
	"awesomeProject/src/torrent/dto"
	"awesomeProject/src/torrent/persistence"
	"testing"
)

func TestGetTorrentList(t *testing.T) {
	torrentDao := persistence.NewTorrentDao()
	torrentDao.DeleteAllTorrents()
	torrentManager := domain.NewTorrentManager(torrentDao, nil)

	for i := 0; i < 2; i++ {
		torrent := dto.NewTorrent("Test torrent", "NEW", "http://test.com")
		torrentDao.CreateTorrent(torrent)
	}
	torrentManager.GetTorrentList("id", 1, 2)

}
