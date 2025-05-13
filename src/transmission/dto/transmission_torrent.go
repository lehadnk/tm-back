package dto

type TransmissionTorrentData struct {
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
