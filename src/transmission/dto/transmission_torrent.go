package dto

type TransmissionTorrent struct {
	Id     int
	Done   int
	Have   string
	ETA    string
	Up     string
	Down   string
	Ratio  int
	Status string
	Name   string
}

type TransmissionTorrentFile struct {
	Id       int
	Done     int
	Priority string
	Get      string
	Size     string
	Name     string
}
