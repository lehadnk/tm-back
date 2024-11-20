package communication

import (
	"fmt"
	"os"
	"time"
	"tm/src/cli/domain"
	"tm/src/filesystem"
	"tm/src/torrent"
)

type DownloadedTorrentsScanner struct {
	torrentService    *torrent.TorrentService
	filesystemService *filesystem.FilesystemService
	cli               domain.CliRunnerInterface
}

func NewDownloadedTorrentsScanner(torrentService *torrent.TorrentService, filesystemService *filesystem.FilesystemService, cli domain.CliRunnerInterface) *DownloadedTorrentsScanner {
	return &DownloadedTorrentsScanner{
		torrentService,
		filesystemService,
		cli,
	}
}

func (scanner *DownloadedTorrentsScanner) Scan() {
	activeTorrentsFromDB := scanner.torrentService.GetActiveTorrentsList()
	mediaDir := os.Getenv("TM_MEDIA_DIR")

	for i := 0; i < activeTorrentsFromDB.FinalTorrentCount; i++ {
		if activeTorrentsFromDB.FinalTorrentArray[i].Torrent.Status == "DOWNLOADING" && activeTorrentsFromDB.FinalTorrentArray[i].TransmissionTorrent.Done == 100 {
			run, err := scanner.cli.Run("mv", []string{activeTorrentsFromDB.FinalTorrentArray[i].Torrent.OutputDirectory, mediaDir + "/" + activeTorrentsFromDB.FinalTorrentArray[i].Torrent.Name})
			if err != nil {
				fmt.Println("Error moving file to filesystem: ", run)
			}

			//filePath := activeTorrentsFromDB.FinalTorrentArray[i].Torrent.Filepath
			//mediaTorrentDir := scanner.filesystemService.CreateMediaDirectory(mediaDir + "/" + activeTorrentsFromDB.FinalTorrentArray[i].Torrent.Name)
			//err := scanner.filesystemService.MoveFile(filePath, mediaTorrentDir)
			//if err != nil {
			//	fmt.Println("Error moving file to filesystem: ", err)
			//}
		}
	}
	fmt.Println("Scanning downloaded torrents...")
}

func (scanner *DownloadedTorrentsScanner) Start() {
	go func() {
		ticker := time.NewTicker(5 * time.Second)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				scanner.Scan()
			}
		}
	}()
}
