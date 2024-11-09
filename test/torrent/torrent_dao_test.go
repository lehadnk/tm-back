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

	torrent := dto.NewTorrent("Test torrent", "NEW", "http://test.com")
	torrentDao.CreateTorrent(torrent)

	readTorrent := torrentDao.GetTorrentById(torrent.Id)
	if reflect.DeepEqual(torrent, readTorrent) {
		t.Errorf("Torrent isn't created")
	}
}

func TestGetListOfTorrents(t *testing.T) {
	torrentDao := persistence.NewTorrentDao()
	torrentDao.DeleteAllTorrents()

	for i := 0; i < 3; i++ {
		torrent := dto.NewTorrent("Test torrent", "NEW", "http://test.com")
		torrentDao.CreateTorrent(torrent)
	}

	torrentsListResult := torrentDao.GetListOfTorrents("id", 1, 3)

	if len(torrentsListResult) != 3 {
		t.Errorf("There are not all torrents")
	}

	torrentsListResult = torrentDao.GetListOfTorrents("id", 2, 3)
	if len(torrentsListResult) != 0 {
		t.Errorf("Torrent should be empty")
	}
}
