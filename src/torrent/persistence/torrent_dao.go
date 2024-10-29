package persistence

import (
	"awesomeProject/src/common"
	"awesomeProject/src/torrent/domain"
	"errors"
	_ "github.com/lib/pq"
	"log"
)

type TorrentDao struct {
	common.AbstractDao
}

func (dbc *TorrentDao) CreateTorrent(torrent domain.Torrent) {
	var torrentId int
	err := dbc.Db.QueryRow(
		"INSERT INTO torrents(name, status) VALUES ($1, $2) RETURNING id",
		torrent.Name, torrent.Status).Scan(&torrentId)
	if err != nil {
		log.Fatalln(errors.New("could not create torrent"))
	}
	torrent.Id = torrentId
}

func (dbc *TorrentDao) GetListOfTorrents(sort string, page int, pageSize int) []domain.Torrent {
	var torrents []domain.Torrent
	var offset = (page - 1) * pageSize

	err := dbc.Db.Select(
		&torrents,
		"SELECT * from torrents ORDER BY $1 LIMIT $2 OFFSET $3", sort, pageSize, offset)

	if err != nil {
		log.Fatalln(errors.New("could not create torrent"))
	}
	return torrents
}

func (dbc *TorrentDao) GetCountOfTorrents() int {
	var torrentsCount int
	err := dbc.Db.Select(
		&torrentsCount,
		"SELECT COUNT * from torrents")
	if err != nil {
		log.Fatalln(errors.New("could not get count"))
	}
	return torrentsCount
}
