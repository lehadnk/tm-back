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
	TransmissionService *transmission.TransmissionService
	Filesystemservice   *filesystem.FilesystemService
}

func NewTorrentManager(
	torrentDao *persistence.TorrentDao,
	transmissionService *transmission.TransmissionService,
	filesystemservice *filesystem.FilesystemService,
) *TorrentManager {
	var newTorrentManager = TorrentManager{
		torrentDao,
		transmissionService,
		filesystemservice,
	}
	return &newTorrentManager
}

func (torrentManager *TorrentManager) GetTorrentList(sort string, page int, pageSize int) dto.FinalTorrentsList {
	torrentsListFromDB := torrentManager.TorrentDao.GetTorrentsList(sort, page, pageSize)
	torrentsCount := torrentManager.TorrentDao.GetCountOfTorrents()
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
		FinalTorrentCount: torrentsCount,
	}
	return finalTorrentsList
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
	torrentDto := dto.NewTorrent("123", "NEW", torrentFilePath, outputDirectory)
	torrentManager.TorrentDao.SaveTorrent(torrentDto)

	return torrentDto, err
}

func (torrentManager *TorrentManager) DeleteTorrent(torrentId int) {
	torrentManager.TransmissionService.DeleteTransmissionTorrent(torrentId)
	torrentManager.TorrentDao.DeleteTorrentById([]int{torrentId})
}
