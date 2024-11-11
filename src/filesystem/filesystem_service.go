package filesystem

import "tm/src/filesystem/domain"

type FilesystemService struct {
}

func NewFilesystemService() *FilesystemService {
	var newNewFilesystemService = FilesystemService{}
	return &newNewFilesystemService
}

func (filesystemservice *FilesystemService) SaveTorrentFile(file []byte, path string) error {
	return domain.WriteToFile(file, path)
}
