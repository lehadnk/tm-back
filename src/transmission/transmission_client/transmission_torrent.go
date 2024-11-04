package transmission_client

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

func NewTransmissionTorrent(
	id int, done string, have string, eta string, up string, down string, ratio string, status string, name string) *TransmissionTorrent {

	return &TransmissionTorrent{}
}
