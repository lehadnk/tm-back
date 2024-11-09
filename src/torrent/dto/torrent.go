package dto

import "time"

type Torrent struct {
	Id       int       `db:"id"`
	Name     string    `db:"name"`
	Status   string    `db:"status"`
	Filepath string    `db:"filepath"`
	Created  time.Time `db:"created"`
	Updated  time.Time `db:"updated"`
}

func NewTorrent(name string, status string, filepath string) *Torrent {
	return &Torrent{
		Name:     name,
		Status:   status,
		Filepath: filepath,
		Created:  time.Now(),
		Updated:  time.Now(),
	}
}

type TorrentsListFromDB struct {
	TorrentsArray []*Torrent
	TorrentsCount int
}
