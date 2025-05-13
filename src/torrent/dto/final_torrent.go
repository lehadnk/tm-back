package dto

import (
	"tm/src/transmission/dto"
)

type AggregatedTorrentData struct {
	Entity           *TorrentEntity               `json:"databaseData"`
	TransmissionData *dto.TransmissionTorrentData `json:"transmissionData"`
}

type AggregatedTorrentDataList struct {
	Torrents []*AggregatedTorrentData `json:"torrents"`
	Count    int                      `json:"count"`
}
