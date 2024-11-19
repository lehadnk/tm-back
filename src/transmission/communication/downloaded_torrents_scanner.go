package communication

import (
	"fmt"
	"time"
	"tm/src/filesystem"
	"tm/src/torrent"
)

type DownloadedTorrentsScanner struct {
	torrentService    *torrent.TorrentService
	filesystemService *filesystem.FilesystemService
}

func NewDownloadedTorrentsScanner(torrentService *torrent.TorrentService, filesystemService *filesystem.FilesystemService) *DownloadedTorrentsScanner {
	return &DownloadedTorrentsScanner{
		torrentService,
		filesystemService,
	}
}

func (scanner *DownloadedTorrentsScanner) Scan() {
	//activeTorrentsFromDB := scanner.torrentService.GetActiveTorrentsList()
	//scanner.filesystemService.CreateMediaDirectory()
	//
	//for i := 0; i < activeTorrentsFromDB.FinalTorrentCount; i++ {
	//	if activeTorrentsFromDB.FinalTorrentArray[i].Torrent.Status == "DOWNLOADING" && activeTorrentsFromDB.FinalTorrentArray[i].TransmissionTorrent.Done == 100 {
	//		err := filesystem.FilesystemService.MoveFile()
	//		if err != nil {
	//			fmt.Println("Error moving file: ", err)
	//		}
	//	}
	//}
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
