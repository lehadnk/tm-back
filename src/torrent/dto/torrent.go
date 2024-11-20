package dto

import "time"

type Torrent struct {
	Id              int       `db:"id" json:"id"`
	Name            string    `db:"name" json:"name"`
	Status          string    `db:"status" json:"status"`
	Filepath        string    `db:"filepath"`
	OutputDirectory string    `db:"output_directory"`
	Created         time.Time `db:"created"`
	Updated         time.Time `db:"updated"`
}

func NewTorrent(name string, status string, filepath string, outputdirectory string) *Torrent {
	return &Torrent{
		Name:            name,
		Status:          status,
		Filepath:        filepath,
		OutputDirectory: outputdirectory,
		Created:         time.Now(),
		Updated:         time.Now(),
	}
}

type TorrentsListFromDB struct {
	TorrentsArray []*Torrent
	TorrentsCount int
}
