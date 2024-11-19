package domain

import (
	"fmt"
	"github.com/jackpal/bencode-go"
	"os"
	"tm/src/torrent/dto"
)

type TorrentParser struct {
}

func NewTorrentParser() *TorrentParser {
	return &TorrentParser{}
}

func (torrentParser *TorrentParser) GetTorrentNameFromBencode(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return "", err
	}
	defer file.Close()

	var torrent dto.TorrentBencodeStructure
	if err := bencode.Unmarshal(file, &torrent); err != nil {
		fmt.Printf("Error decoding .torrent file: %v\n", err)
		return "", err
	}
	return torrent.Info.Name, nil
}
