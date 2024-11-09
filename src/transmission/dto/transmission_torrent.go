package dto

type TransmissionTorrent struct {
	Id     int
	Done   string
	Have   string
	ETA    string
	Up     string
	Down   string
	Ratio  int
	Status string
	Name   string
}
