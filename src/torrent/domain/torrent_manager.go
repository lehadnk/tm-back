package domain

import (
	"errors"
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

func (torrentManager *TorrentManager) AddTorrent(file []byte) (*dto.Torrent, error, error) {
	filename := common.StringWithCharset(24, "abcdefghijklmnopqrstuvwxyz")

	torrentFilePath, err := torrentManager.Filesystemservice.SaveTorrentFile(file, filename+".torrent")
	if err != nil {
		log.Fatalln("Error while saving file:" + err.Error())
		return nil, nil, err
	}
	transmissionTorrentName, err := torrentManager.TorrentParser.GetTorrentNameFromBencode(torrentFilePath)
	torrentFromDB := torrentManager.TorrentDao.GetTorrentByName(transmissionTorrentName)
	if torrentFromDB != nil {
		return nil, errors.New("Torrent already exists"), nil
	}

	outputDirectory, err := torrentManager.Filesystemservice.CreateTorrentOutputDirectory(filename)
	if err != nil {
		log.Fatalln("Error while creating output directory:" + err.Error())
		return nil, nil, err
	}

	torrentManager.TransmissionService.AddTransmissionTorrentFile(torrentFilePath, outputDirectory)
	torrentDto := dto.NewTorrent(transmissionTorrentName, "DOWNLOADING", torrentFilePath, outputDirectory)
	torrentManager.TorrentDao.SaveTorrent(torrentDto)

	return torrentDto, nil, nil
}

func (torrentManager *TorrentManager) DeleteTorrent(torrentId int) error {
	torrent := torrentManager.TorrentDao.GetTorrentById(torrentId)
	if torrent == nil {
		return errors.New("Torrent not found")
	}

	transmissionTorrent := torrentManager.TransmissionService.GetTransmissionTorrentByName(torrent.Name)
	if transmissionTorrent != nil {
		err := torrentManager.TransmissionService.DeleteTransmissionTorrent(transmissionTorrent.Id)
		if err != nil {
			return err
		}
	}

	torrentManager.TorrentDao.DeleteTorrentById(torrentId)
	return nil
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

	var finalTorrents = make([]*dto.FinalTorrent, 0)
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
		FinalTorrentCount: len(finalTorrents),
	}
	return finalTorrentsList
}
