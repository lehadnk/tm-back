package filesystem

import "tm/src/filesystem/domain"

type FilesystemService struct {
	filesystemManager *domain.FilesystemManager
}

func NewFilesystemService(filesystemManager *domain.FilesystemManager) *FilesystemService {
	var newNewFilesystemService = FilesystemService{
		filesystemManager: filesystemManager,
	}
	return &newNewFilesystemService
}

func (filesystemservice *FilesystemService) SaveTorrentFile(file []byte, fileName string) (string, error) {
	return filesystemservice.filesystemManager.WriteTorrentFile(file, fileName)
}

func (filesystemservice *FilesystemService) CreateTorrentFileDirectory(directoryPath string) error {
	return filesystemservice.filesystemManager.CreateTorrentFileDirectory(directoryPath)
}

func (filesystemservice *FilesystemService) CreateTorrentOutputDirectory(directoryPath string) (string, error) {
	return filesystemservice.filesystemManager.CreateTorrentOutputDirectory(directoryPath)
}

func (filesystemservice *FilesystemService) CreateMediaDirectory(directoryPath string) error {
	return filesystemservice.filesystemManager.CreateMediaDirectory(directoryPath)
}

func (filesystemservice *FilesystemService) MoveFile(sourcePath string, destinationPath string) error {
	return filesystemservice.filesystemManager.MoveFile(sourcePath, destinationPath)
}

func (fs *FilesystemService) GetFreeSpaceLeft() uint64 {
	return fs.filesystemManager.GetFreeSpaceLeft()
}
