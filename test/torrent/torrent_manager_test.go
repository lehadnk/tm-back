package torrent

import (
	"strconv"
	"testing"
	cli_domain "tm/src/cli/domain"
	"tm/src/torrent/domain"
	"tm/src/torrent/dto"
	"tm/src/torrent/persistence"
	"tm/src/transmission"
	transmission_domain "tm/src/transmission/domain"
)

func TestGetTorrentList(t *testing.T) {
	mockCliRunner := cli_domain.MockCliRunner{}
	transmissionClient := transmission_domain.NewTransmissionClient(&mockCliRunner)
	transmissionService := transmission.NewTransmissionService(transmissionClient)
	torrentDao := persistence.NewTorrentDao()
	torrentDao.DeleteAllTorrents()
	torrentManager := domain.NewTorrentManager(torrentDao, transmissionService)

	mockCliRunner.On("transmission-remote", []string{"-a"}, "localhost:9091/transmission/rpc/ responded: success", nil)

	for i := 0; i < 2; i++ {
		torrent := dto.NewTorrent("Test torrent "+strconv.Itoa(i), "NEW", "/media/torrents/"+strconv.Itoa(i)+".torrent")
		torrentDao.CreateTorrent(torrent)
	}

	mockCliRunner.On("transmission-remote", []string{"-l"}, "    ID   Done       Have  ETA           Up    Down  Ratio  Status       Name\n     1*  0%    2.20 GB  10 hrs         0.0     0.0    1.4  Stopped      Test torrent 1\n     2*  0%    2.20 GB  10 hrs         0.0     0.0    1.4  Downloading      Test torrent 2\n17     0%    3.11 MB  10 hrs       0.0   644.0   0.00  Downloading  TNG_s1", nil)
	torrentManager.GetTorrentList("id", 1, 2)
}
