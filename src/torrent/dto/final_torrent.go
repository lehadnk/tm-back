package dto

import (
	"tm/src/transmission/dto"
)

type FinalTorrent struct {
	Torrent             *Torrent                 `json:"torrent"`
	TransmissionTorrent *dto.TransmissionTorrent `json:"transmissionTorrent"`
}

type FinalTorrentsList struct {
	FinalTorrentArray []*FinalTorrent `json:"torrents"`
	FinalTorrentCount int             `json:"count"`
}
