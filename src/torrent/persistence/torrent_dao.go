package persistence

import (
	_ "github.com/lib/pq"
	"log"
	"tm/src/common"
	"tm/src/torrent/dto"
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
		"INSERT INTO torrents(name, status, filepath, output_directory, created, updated) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id",
		torrent.Name, torrent.Status, torrent.Filepath, torrent.OutputDirectory, torrent.Created, torrent.Updated,
	).Scan(&torrentId)
	if err != nil {
		log.Fatalln("Could not create torrent: " + err.Error())
	}
	torrent.Id = torrentId
}

func (dbc *TorrentDao) GetTorrentById(torrentId int) *dto.Torrent {
	torrent := dto.Torrent{}
	err := dbc.Db.Get(&torrent, "SELECT * from torrents WHERE id = $1", torrentId)
	if err != nil {
		log.Println("Could not select torrent: " + err.Error())
	}
	return &torrent
}

func (dbc *TorrentDao) GetActiveTorrentList() []*dto.Torrent {
	var torrents []*dto.Torrent
	torrentStatus := "DOWNLOADING"

	err := dbc.Db.Select(
		&torrents, "SELECT * from torrents WHERE status = $1", torrentStatus)
	if err != nil {
		log.Fatalln("Could not find active torrents: " + err.Error())
	}
	return torrents
}

func (dbc *TorrentDao) GetCountOfActiveTorrents() int {
	var torrentsCount int
	torrentStatus := "DOWNLOADING"

	err := dbc.Db.Get(&torrentsCount, "SELECT COUNT(*) from torrents WHERE status = $1", torrentStatus)
	if err != nil {
		log.Fatalln("Could not obtain active torrent count information: " + err.Error())
	}
	return torrentsCount
}

func (dbc *TorrentDao) GetTorrentsList(sort string, page int, pageSize int) []*dto.Torrent {
	var torrents []*dto.Torrent
	var offset = (page - 1) * pageSize

	err := dbc.Db.Select(
		&torrents, "SELECT * from torrents ORDER BY "+sort+" LIMIT $1 OFFSET $2", pageSize, offset)
	if err != nil {
		log.Fatalln("Could not select torrents: " + err.Error())
	}
	return torrents
}

func (dbc *TorrentDao) GetCountOfTorrents() int {
	var torrentsCount int
	err := dbc.Db.Get(&torrentsCount, "SELECT COUNT(*) from torrents")
	if err != nil {
		log.Fatalln("Could not obtain torrent count information: " + err.Error())
	}
	return torrentsCount
}

func (dbc *TorrentDao) DeleteTorrentById(id int) {
	_, err := dbc.Db.Exec("DELETE FROM torrents WHERE id = $1", id)
	if err != nil {
		log.Println("Could not delete torrent: " + err.Error())
	}
}

func (dbc *TorrentDao) DeleteAllTorrents() {
	err := dbc.Db.QueryRow("TRUNCATE TABLE torrents")
	if err.Err() != nil {
		log.Fatalln("Could not truncate torrents table: " + err.Err().Error())
	}
}
