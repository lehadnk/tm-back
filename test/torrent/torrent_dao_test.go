package torrent

import (
	_ "github.com/lib/pq"
	"reflect"
	"testing"
	"tm/src/torrent/dto"
	"tm/src/torrent/persistence"
)

func TestCreateTorrentInDb(t *testing.T) {
	torrentDao := persistence.NewTorrentDao()

	torrent := dto.NewTorrentEntity("Test torrent", "NEW", "http://test.com", "/tmp")
	torrentDao.SaveTorrent(torrent)

	readTorrent := torrentDao.GetTorrentById(torrent.Id)
	if reflect.DeepEqual(torrent, readTorrent) {
		t.Errorf("TorrentEntity isn't created")
	}
}

func TestGetListOfTorrents(t *testing.T) {
	torrentDao := persistence.NewTorrentDao()
	torrentDao.DeleteAllTorrents()

	for i := 0; i < 3; i++ {
		torrent := dto.NewTorrentEntity("Test torrent", "NEW", "http://test.com", "/tmp")
		torrentDao.SaveTorrent(torrent)
	}

	torrentsListResult := torrentDao.GetTorrentsList("id", 1, 3)

	if len(torrentsListResult) != 3 {
		t.Errorf("There are not all torrents")
	}

	torrentsListResult = torrentDao.GetTorrentsList("id", 2, 3)
	if len(torrentsListResult) != 0 {
		t.Errorf("TorrentEntity should be empty")
	}
}
