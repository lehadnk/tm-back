package communication

import (
	"fmt"
	"time"
)

type DownloadedTorrentsScanner struct {
}

func NewDownloadedTorrentsScanner() *DownloadedTorrentsScanner {
	return &DownloadedTorrentsScanner{}
}

func (scanner *DownloadedTorrentsScanner) Scan() {
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
