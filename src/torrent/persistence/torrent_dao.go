package persistence

import (
	"awesomeProject/src/common"
	"awesomeProject/src/torrent/dto"
	"errors"
	_ "github.com/lib/pq"
	"log"
)

type TorrentDao struct {
	common.AbstractDao
}

func NewTorrentDao() *TorrentDao {
	var newTorrentDao = TorrentDao{}
	newTorrentDao.Connect()
	return &newTorrentDao
}

func (dbc *TorrentDao) SaveTorrent(torrent *dto.Torrent) {
	var torrentId int
	err := dbc.Db.QueryRow(
		"INSERT INTO torrents(name, status, filepath, created, updated) VALUES ($1, $2, $3, $4, $5) RETURNING id",
		torrent.Name, torrent.Status, torrent.Filepath, torrent.Created, torrent.Updated).Scan(&torrentId)
	if err != nil {
		log.Fatalln(errors.New("could not create torrent"))
	}
	torrent.Id = torrentId
}

func (dbc *TorrentDao) GetTorrentById(torrentId int) *dto.Torrent {
	torrent := dto.Torrent{}
	err := dbc.Db.Get(
		&torrent,
		"SELECT * from torrents WHERE id = $1",
		torrentId)
	if err != nil {
		log.Fatalln(errors.New("could not get torrent"))
	}
	return &torrent
}

func (dbc *TorrentDao) GetTorrentByStatus(torrentStatus string) *dto.Torrent {
	torrent := dto.Torrent{}
	err := dbc.Db.Get(
		&torrent, "SELECT * from torrents WHERE status = $1", torrentStatus)
	if err != nil {
		log.Fatalln(errors.New("could not get torrent in status " + torrentStatus))
	}
	return &torrent
}

func (dbc *TorrentDao) GetTorrentsList(sort string, page int, pageSize int) []*dto.Torrent {
	var torrents []*dto.Torrent
	var offset = (page - 1) * pageSize

	err := dbc.Db.Select(
		&torrents, "SELECT * from torrents ORDER BY $1 LIMIT $2 OFFSET $3", sort, pageSize, offset)
	if err != nil {
		log.Fatalln(errors.New("could not get torrents"))
	}
	return torrents
}

func (dbc *TorrentDao) GetCountOfTorrents() int {
	var torrentsCount int
	err := dbc.Db.Get(
		&torrentsCount, "SELECT COUNT(*) from torrents")
	if err != nil {
		log.Fatalln(errors.New("could not get count"))
	}
	return torrentsCount
}

func (dbc *TorrentDao) DeleteTorrentById(ids []int) {
	err := dbc.Db.QueryRow(
		"DELETE from torrents WHERE id IN($1)", ids)
	if err != nil {
		log.Fatalln(errors.New("could not delete torrents"))
	}
}

func (dbc *TorrentDao) DeleteAllTorrents() {
	err := dbc.Db.QueryRow(
		"TRUNCATE TABLE torrents")
	if err.Err() != nil {
		log.Fatalln(errors.New("could not delete all torrents"))
	}
}
