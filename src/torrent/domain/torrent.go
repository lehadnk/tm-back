package domain

type Torrent struct {
	Id     int    `db:"id"`
	Name   string `db:"name"`
	Status string `db:"status"`
}

func NewTorrent(name string, status string) *Torrent {
	return &Torrent{
		Name:   name,
		Status: status,
	}
}

type TorrentsList struct {
	TorrentsArray []Torrent
	TorrentsCount int
}
