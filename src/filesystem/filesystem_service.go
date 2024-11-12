package filesystem

import "tm/src/filesystem/domain"

type FilesystemService struct {
	FilesystemManager *domain.FilesystemManager
}

func NewFilesystemService(fileSystemManager *domain.FilesystemManager) *FilesystemService {
	var newNewFilesystemService = FilesystemService{
		FilesystemManager: fileSystemManager,
	}
	return &newNewFilesystemService
}

func (filesystemservice *FilesystemService) SaveTorrentFile(file []byte, path string) error {
	return filesystemservice.FilesystemManager.WriteToFile(file, path)
}

func (filesystemservice *FilesystemService) CreateTorrentFileDirectory(directoryPath string) error {
	return filesystemservice.FilesystemManager.CreateTorrentFileDirectory(directoryPath)
}

func (filesystemservice *FilesystemService) CreateTorrentOutputDirectory(directoryPath string) error {
	return filesystemservice.FilesystemManager.CreateTorrentOutputDirectory(directoryPath)
}

func (filesystemservice *FilesystemService) CreateMediaDirectory(directoryPath string) error {
	return filesystemservice.FilesystemManager.CreateMediaDirectory(directoryPath)
}

func (filesystemservice *FilesystemService) MoveFile(sourcePath string, destinationPath string) error {
	return filesystemservice.FilesystemManager.MoveFile(sourcePath, destinationPath)
}
