package dto

type TransmissionTorrent struct {
	Id     int    `json:"id"`
	Done   int    `json:"done"`
	Have   string `json:"have"`
	ETA    string `json:"eta"`
	Up     string `json:"up"`
	Down   string `json:"down"`
	Ratio  int    `json:"ration"`
	Status string `json:"status"`
	Name   string `json:"name"`
}

type TransmissionTorrentFile struct {
	Id       int
	Done     int
	Priority string
	Get      string
	Size     string
	Name     string
}
