package communication

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
	"tm/src/cli/domain"
	"tm/src/filesystem"
	"tm/src/torrent"
	"tm/src/torrent/dto"
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

func (scanner *DownloadedTorrentsScanner) IsTorrentFullyDownloaded(torrentData *dto.AggregatedTorrentData) (bool, error) {
	if torrentData.Entity.Status != "DOWNLOADING" || torrentData.TransmissionData.Done != 100 {
		return false, nil
	}

	/* The issue with transmission-remote is that it reports 100% done before doing the checksum verification. And until the verification is finished, file names have .part suffix. So we have to also check output filenames to ensure that downloading process is -actually- done.*/
	hasPartFiles := false

	err := filepath.Walk(torrentData.Entity.OutputDirectory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err // propagate the error
		}
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".part") {
			hasPartFiles = true
			return filepath.SkipDir
		}
		return nil
	})

	if err != nil {
		return false, err
	}
	return !hasPartFiles, nil
}

func (scanner *DownloadedTorrentsScanner) Scan() {
	downloadingTorrentsList := scanner.torrentService.GetDownloadingTorrentsList()
	mediaDir := os.Getenv("TM_MEDIA_DIR")

	fmt.Println("Scanning downloaded torrents...")

	for i := 0; i < downloadingTorrentsList.Count; i++ {
		isDownloaded, scanError := scanner.IsTorrentFullyDownloaded(downloadingTorrentsList.Torrents[i])
		if scanError != nil {
			fmt.Println("Error during the torrent output directory post-download scanning process: " + scanError.Error())
		}
		if !isDownloaded {
			continue
		}

		run, err := scanner.cli.Run("mv", []string{downloadingTorrentsList.Torrents[i].Entity.OutputDirectory, mediaDir + "/" + downloadingTorrentsList.Torrents[i].Entity.Name})
		if err != nil {
			fmt.Println("Error moving file to filesystem: ", run)
		}
	}
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
