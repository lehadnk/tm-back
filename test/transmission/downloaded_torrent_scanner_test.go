package transmission

import (
	"log"
	"os"
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
	transmission_dto "tm/src/transmission/dto"
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

	transmissionOutputDirectoryName := common.StringWithCharset(24, "abcdefghijklmnopqrstuvwxyz")
	err := os.Mkdir("/tmp/"+transmissionOutputDirectoryName, 0777)
	if err != nil {
		log.Fatalln("Transmission output directory was not created: " + err.Error())
	}

	torrentName := common.StringWithCharset(24, "abcdefghijklmnopqrstuvwxyz")

	torrentEntity := dto.NewTorrentEntity(torrentName, "DOWNLOADING", "/tmp/", "/tmp/"+transmissionOutputDirectoryName)
	torrentDao.SaveTorrent(torrentEntity)

	mockCliRunner.On("transmission-remote", []string{"-l"}, "    ID   Done       Have  ETA           Up    Down  Ratio  Status       Name\n     1   100%    2.20 GB  Done         0.0     0.0   0.00  Idle         "+torrentName+"\nSum:             2.20 GB               0.0     0.0", nil)
	mockCliRunner.On("mv", []string{"/tmp/" + transmissionOutputDirectoryName, "/tmp/" + torrentName}, "", nil)

	downloadedTorrentsScanner.Scan()
	commandRun := mockCliRunner.WasCommandRun("mv /tmp/" + transmissionOutputDirectoryName + " /tmp/" + torrentName)
	if !commandRun {
		log.Fatalln("Command was not run")
	}
}

func TestWaitBeforeMovingDownloadWithIncorrectlyReportedDoneStatus(t *testing.T) {
	mockCliRunner := cli_domain.MockCliRunner{}
	filesystemService := filesystem.NewFilesystemService(filesystem_domain.NewFilesystemManager("/tmp", "/tmp", "/tmp"))
	torrentParser := torrent_domain.NewTorrentParser()
	torrentDao := persistence.NewTorrentDao()
	transmissionClient := transmission_domain.NewTransmissionClient(&mockCliRunner)
	transmissionService := transmission.NewTransmissionService(transmissionClient)
	torrentManager := torrent_domain.NewTorrentManager(torrentDao, torrentParser, transmissionService, filesystemService)
	torrentService := torrent.NewTorrentService(torrentManager)
	downloadedTorrentsScanner := communication.NewDownloadedTorrentsScanner(torrentService, filesystemService, &mockCliRunner)

	transmissionOutputDirectoryName := common.StringWithCharset(24, "abcdefghijklmnopqrstuvwxyz")
	err := os.Mkdir("/tmp/"+transmissionOutputDirectoryName, 0777)
	if err != nil {
		log.Fatalln("Transmission output directory was not created: " + err.Error())
	}
	_, err = os.Create("/tmp/" + transmissionOutputDirectoryName + "/1.part")
	if err != nil {
		log.Fatalln("Error while creating a test output .part file")
	}

	torrentName := common.StringWithCharset(24, "abcdefghijklmnopqrstuvwxyz")
	mockCliRunner.On("transmission-remote", []string{"-l"}, "    ID   Done       Have  ETA           Up    Down  Ratio  Status       Name\n     1   100%    2.20 GB  Done         0.0     0.0   0.00  Idle         "+torrentName+"\nSum:             2.20 GB               0.0     0.0", nil)

	torrentEntity := dto.NewTorrentEntity(torrentName, "DOWNLOADING", "/tmp/", "/tmp/"+transmissionOutputDirectoryName)
	torrentDao.SaveTorrent(torrentEntity)

	transmissionData := transmission_dto.TransmissionTorrentData{Done: 100}

	aggregatedTorrentData := dto.AggregatedTorrentData{torrentEntity, &transmissionData}

	isDownloaded, _ := downloadedTorrentsScanner.IsTorrentFullyDownloaded(&aggregatedTorrentData)
	if isDownloaded {
		log.Fatalln("Torrent has incorrectly reported download status, but still marked as ready to move")
	}

	err = os.Rename("/tmp/"+transmissionOutputDirectoryName+"/1.part", "/tmp/"+transmissionOutputDirectoryName+"1.tmp")
	if err != nil {
		log.Fatalln("Error while renaming the test file")
	}

	isDownloaded, _ = downloadedTorrentsScanner.IsTorrentFullyDownloaded(&aggregatedTorrentData)
	if !isDownloaded {
		log.Fatalln("Torrent is now marked as ready to move, but incorrectly reported the download status")
	}
}
