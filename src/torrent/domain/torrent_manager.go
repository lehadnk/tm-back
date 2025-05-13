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

func (torrentManager *TorrentManager) AddTorrent(file []byte) (*dto.TorrentEntity, error, error) {
	filename := common.StringWithCharset(24, "abcdefghijklmnopqrstuvwxyz")

	torrentFilePath, err := torrentManager.Filesystemservice.SaveTorrentFile(file, filename+".torrent")
	if err != nil {
		log.Fatalln("Error while saving file:" + err.Error())
		return nil, nil, err
	}
	transmissionTorrentName, err := torrentManager.TorrentParser.GetTorrentNameFromBencode(torrentFilePath)
	torrentFromDB := torrentManager.TorrentDao.GetTorrentByName(transmissionTorrentName)
	if torrentFromDB != nil {
		return nil, errors.New("TorrentEntity already exists"), nil
	}

	outputDirectory, err := torrentManager.Filesystemservice.CreateTorrentOutputDirectory(filename)
	if err != nil {
		log.Fatalln("Error while creating output directory:" + err.Error())
		return nil, nil, err
	}

	wasAddedToTransmission, err := torrentManager.TransmissionService.AddTorrentFileToTransmission(torrentFilePath, outputDirectory)
	if !wasAddedToTransmission {
		return nil, nil, errors.New("Torrent was not added to transmission. See application logs for error")
	}

	torrentDto := dto.NewTorrentEntity(transmissionTorrentName, "DOWNLOADING", torrentFilePath, outputDirectory)
	torrentManager.TorrentDao.SaveTorrent(torrentDto)

	return torrentDto, nil, nil
}

func (torrentManager *TorrentManager) DeleteTorrent(torrentId int) error {
	torrent := torrentManager.TorrentDao.GetTorrentById(torrentId)
	if torrent == nil {
		return errors.New("TorrentEntity not found")
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

func (torrentManager *TorrentManager) GetTorrentsListFromDatabase(sort string, page int, pageSize int) dto.AggregatedTorrentDataList {
	torrentsListFromDB := torrentManager.TorrentDao.GetTorrentsList(sort, page, pageSize)
	return torrentManager.hydrateTorrentListWithTransmissionData(torrentsListFromDB)
}

func (torrentManager *TorrentManager) GetDownloadingTorrentsList() dto.AggregatedTorrentDataList {
	torrentsListFromDB := torrentManager.TorrentDao.GetDownloadingTorrentList()
	return torrentManager.hydrateTorrentListWithTransmissionData(torrentsListFromDB)
}

func (torrentManager *TorrentManager) hydrateTorrentListWithTransmissionData(torrentsListFromDB []*dto.TorrentEntity) dto.AggregatedTorrentDataList {
	torrentsListFromTransmission := torrentManager.TransmissionService.GetTransmissionTorrentList()

	var aggregatedTorrentDataList = make([]*dto.AggregatedTorrentData, 0)
	for i := 0; i < len(torrentsListFromDB); i++ {
		for j := 0; j < len(torrentsListFromTransmission); j++ {
			if torrentsListFromDB[i].Name != torrentsListFromTransmission[j].Name {
				continue
			}

			aggregatedData := dto.AggregatedTorrentData{
				Entity:           torrentsListFromDB[i],
				TransmissionData: torrentsListFromTransmission[j],
			}
			aggregatedTorrentDataList = append(aggregatedTorrentDataList, &aggregatedData)
		}
	}

	finalTorrentsList := dto.AggregatedTorrentDataList{
		Torrents: aggregatedTorrentDataList,
		Count:    len(aggregatedTorrentDataList),
	}
	return finalTorrentsList
}
