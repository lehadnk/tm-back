package torrent

import (
	"log"
	"strconv"
	"testing"
	cli_domain "tm/src/cli/domain"
	"tm/src/filesystem"
	filesystem_domain "tm/src/filesystem/domain"
	torrent_domain "tm/src/torrent/domain"
	"tm/src/torrent/dto"
	"tm/src/torrent/persistence"
	"tm/src/transmission"
	transmission_domain "tm/src/transmission/domain"
)

func TestGetTorrentList(t *testing.T) {
	mockCliRunner := cli_domain.MockCliRunner{}
	transmissionClient := transmission_domain.NewTransmissionClient(&mockCliRunner)
	transmissionService := transmission.NewTransmissionService(transmissionClient)
	filesystemService := filesystem.NewFilesystemService(
		filesystem_domain.NewFilesystemManager("/tmp", "/tmp", "/tmp"),
	)
	torrentDao := persistence.NewTorrentDao()
	torrentDao.DeleteAllTorrents()
	torrentManager := torrent_domain.NewTorrentManager(torrentDao, transmissionService, filesystemService)

	mockCliRunner.On("transmission-remote", []string{"-a"}, "localhost:9091/transmission/rpc/ responded: success", nil)

	for i := 0; i < 2; i++ {
		torrent := dto.NewTorrent("Test torrent "+strconv.Itoa(i), "NEW", "/media/torrents/"+strconv.Itoa(i)+".torrent", "/tmp")
		torrentDao.SaveTorrent(torrent)
	}

	mockCliRunner.On("transmission-remote", []string{"-l"}, "    ID   Done       Have  ETA           Up    Down  Ratio  Status       Name\n     1*  0%    2.20 GB  10 hrs         0.0     0.0    1.4  Stopped      Test torrent 1\n     2*  0%    2.20 GB  10 hrs         0.0     0.0    1.4  Downloading      Test torrent 2\n17     0%    3.11 MB  10 hrs       0.0   644.0   0.00  Downloading  TNG_s1", nil)
	torrentList := torrentManager.GetTorrentsList("id", 1, 2)
	if torrentList.FinalTorrentCount != 2 {
		log.Fatalln("Expected 2 torrents, got", torrentList.FinalTorrentCount)
	}
}

func TestAddTorrent(t *testing.T) {
	mockCliRunner := cli_domain.MockCliRunner{}
	transmissionClient := transmission_domain.NewTransmissionClient(&mockCliRunner)
	transmissionService := transmission.NewTransmissionService(transmissionClient)
	filesystemManager := filesystem_domain.NewFilesystemManager("/tmp", "/tmp", "/tmp")
	filesystemService := filesystem.NewFilesystemService(filesystemManager)
	torrentDao := persistence.NewTorrentDao()
	torrentManager := torrent_domain.NewTorrentManager(torrentDao, transmissionService, filesystemService)

	mockCliRunner.On("transmission-remote", []string{"-a"}, "localhost:9091/transmission/rpc/ responded: success", nil)

	testFile, _ := filesystemManager.ReadFile("../test.torrent")
	if testFile == nil {
		log.Fatalln("Error reading file")
	}

	testTorrent, _ := torrentManager.AddTorrent(testFile)

	torrentFromDB := torrentDao.GetTorrentById(testTorrent.Id)
	if torrentFromDB == nil {
		log.Fatalln("Torrent does not exist in database")
	}

	torrentFromDisk, _ := filesystemManager.ReadFile(testTorrent.Filepath)
	if torrentFromDisk == nil {
		log.Fatalln("Error reading file from disk")
	}

	commandRun := mockCliRunner.WasCommandRun("transmission-remote -a " + testTorrent.Filepath + " -w " + testTorrent.OutputDirectory)
	if !commandRun {
		log.Fatalln("Command was not run")
	}
}
