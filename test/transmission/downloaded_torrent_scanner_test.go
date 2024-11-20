package transmission

import (
	"testing"
	cli_domain "tm/src/cli/domain"
	"tm/src/common"
	"tm/src/filesystem"
	filesystem_domain "tm/src/filesystem/domain"
	"tm/src/torrent"
	torrent_domain "tm/src/torrent/domain"
	"tm/src/torrent/dto"
	"tm/src/torrent/persistence"
	"tm/src/transmission"
	"tm/src/transmission/communication"
	transmission_domain "tm/src/transmission/domain"
)

func TestMoveTorrentFileAfterDownload(t *testing.T) {
	mockCliRunner := cli_domain.MockCliRunner{}
	torrentParser := torrent_domain.NewTorrentParser()
	torrentDao := persistence.NewTorrentDao()
	transmissionClient := transmission_domain.NewTransmissionClient(&mockCliRunner)
	transmissionService := transmission.NewTransmissionService(transmissionClient)
	filesystemService := filesystem.NewFilesystemService(filesystem_domain.NewFilesystemManager("/tmp", "/tmp", "/tmp"))
	torrentManager := torrent_domain.NewTorrentManager(torrentDao, torrentParser, transmissionService, filesystemService)
	torrentService := torrent.NewTorrentService(torrentManager)
	downloadedTorrentsScanner := communication.NewDownloadedTorrentsScanner(torrentService, filesystemService, &mockCliRunner)

	mockCliRunner.On("transmission-remote", []string{"-l"}, "    ID   Done       Have  ETA           Up    Down  Ratio  Status       Name\n     1   100%    2.20 GB  Done         0.0     0.0   0.00  Idle         Ochen.strashnoe.kino.2000.RUS.BDRip.XviD.AC3.-HQCLUB\nSum:             2.20 GB               0.0     0.0", nil)

	filename := common.StringWithCharset(24, "abcdefghijklmnopqrstuvwxyz")
	torrentTest := dto.NewTorrent("Ochen.strashnoe.kino.2000.RUS.BDRip.XviD.AC3.-HQCLUB", "DOWNLOADING", "/tmp/", "/tmp/"+filename)
	torrentDao.SaveTorrent(torrentTest)

	downloadedTorrentsScanner.Scan()

}
