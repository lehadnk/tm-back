package domain

import (
	"log"
	"tm/src/common"
	"tm/src/filesystem"
	"tm/src/torrent/dto"
	"tm/src/torrent/persistence"
	"tm/src/transmission"
)

type TorrentManager struct {
	TorrentDao          *persistence.TorrentDao
	TorrentParser       *TorrentParser
	TransmissionService *transmission.TransmissionService
	Filesystemservice   *filesystem.FilesystemService
}

func NewTorrentManager(
	torrentDao *persistence.TorrentDao,
	torrentParser *TorrentParser,
	transmissionService *transmission.TransmissionService,
	filesystemservice *filesystem.FilesystemService,
) *TorrentManager {
	var newTorrentManager = TorrentManager{
		torrentDao,
		torrentParser,
		transmissionService,
		filesystemservice,
	}
	return &newTorrentManager
}

func (torrentManager *TorrentManager) AddTorrent(file []byte) (*dto.Torrent, error) {
	filename := common.StringWithCharset(24, "abcdefghijklmnopqrstuvwxyz")

	torrentFilePath, err := torrentManager.Filesystemservice.SaveTorrentFile(file, filename+".torrent")
	if err != nil {
		log.Fatalln("Error while saving file:" + err.Error())
		return nil, err
	}

	outputDirectory, err := torrentManager.Filesystemservice.CreateTorrentOutputDirectory(filename)
	if err != nil {
		log.Fatalln("Error while creating output directory:" + err.Error())
		return nil, err
	}

	torrentManager.TransmissionService.AddTransmissionTorrentFile(torrentFilePath, outputDirectory)
	transmissionTorrentName, err := torrentManager.TorrentParser.GetTorrentNameFromBencode(torrentFilePath)
	torrentDto := dto.NewTorrent(transmissionTorrentName, "NEW", torrentFilePath, outputDirectory)
	torrentManager.TorrentDao.SaveTorrent(torrentDto)

	return torrentDto, err
}

func (torrentManager *TorrentManager) DeleteTorrent(torrentId int) {
	torrentManager.TransmissionService.DeleteTransmissionTorrent(torrentId)
	torrentManager.TorrentDao.DeleteTorrentById(torrentId)
}
func (torrentManager *TorrentManager) GetTorrentsList(sort string, page int, pageSize int) dto.FinalTorrentsList {
	torrentsListFromDB := torrentManager.TorrentDao.GetTorrentsList(sort, page, pageSize)
	torrentsCount := torrentManager.TorrentDao.GetCountOfTorrents()
	return torrentManager.buildFinalTorrentList(torrentsListFromDB, torrentsCount)
}

func (torrentManager *TorrentManager) GetActiveTorrentsList() dto.FinalTorrentsList {
	torrentsListFromDB := torrentManager.TorrentDao.GetActiveTorrentList()
	torrentsCount := torrentManager.TorrentDao.GetCountOfActiveTorrents()
	return torrentManager.buildFinalTorrentList(torrentsListFromDB, torrentsCount)
}

func (torrentManager *TorrentManager) buildFinalTorrentList(torrentsListFromDB []*dto.Torrent, count int) dto.FinalTorrentsList {
	torrentsListFromTransmission := torrentManager.TransmissionService.GetTransmissionTorrentList()

	var finalTorrents []*dto.FinalTorrent
	for i := 0; i < len(torrentsListFromDB); i++ {
		for j := 0; j < len(torrentsListFromTransmission); j++ {
			if torrentsListFromDB[i].Name != torrentsListFromTransmission[j].Name {
				continue
			}

			finalTorrent := dto.FinalTorrent{
				Torrent:             torrentsListFromDB[i],
				TransmissionTorrent: torrentsListFromTransmission[j],
			}
			finalTorrents = append(finalTorrents, &finalTorrent)
		}
	}
	finalTorrentsList := dto.FinalTorrentsList{
		FinalTorrentArray: finalTorrents,
		FinalTorrentCount: count,
	}
	return finalTorrentsList
}
