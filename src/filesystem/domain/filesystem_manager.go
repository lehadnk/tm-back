package domain

import (
	"fmt"
	"io"
	"log"
	"os"
)

type FilesystemManager struct {
	torrentFileDir   string
	torrentOutputDir string
	mediaDir         string
}

func NewFilesystemManager(torrentFileDir string, torrentOutputDir string, mediaDir string) *FilesystemManager {
	var newFilesystemManager = FilesystemManager{
		torrentFileDir:   torrentFileDir,
		torrentOutputDir: torrentOutputDir,
		mediaDir:         mediaDir,
	}
	return &newFilesystemManager
}

func (filesystemManager *FilesystemManager) ReadFile(path string) ([]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil, err
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatalln("Error reading file:" + err.Error())
	}

	return data, nil
}

func (filesystemManager *FilesystemManager) WriteToFile(file []byte, path string) error {
	err := os.WriteFile(path, file, 0777)
	if err != nil {
		log.Fatalln("Error writing to file: " + err.Error())
	}
	return err
}

func (filesystemManager *FilesystemManager) CreateOutputDirectory(directoryPath string) error {
	err := os.MkdirAll(directoryPath, 0777)
	if err != nil {
		log.Fatalln("Error creating directory:" + err.Error())
	}
	return err
}

func (filesystemManager *FilesystemManager) MoveFile(sourcePath string, destinationPath string) error {
	err := os.Rename(sourcePath, destinationPath)
	if err != nil {
		log.Fatalln("Error while moving file:" + err.Error())
	}
	return err
}

func (filesystemManager *FilesystemManager) CreateTorrentFileDirectory(directoryName string) error {
	err := filesystemManager.CreateOutputDirectory(filesystemManager.torrentFileDir + "/" + directoryName)
	if err != nil {
		log.Fatalln("Error creating torrent file directory:" + err.Error())
	}
	return err
}

func (filesystemManager *FilesystemManager) CreateTorrentOutputDirectory(directoryName string) error {
	err := filesystemManager.CreateOutputDirectory(filesystemManager.torrentOutputDir + "/" + directoryName)
	if err != nil {
		log.Fatalln("Error creating torrent output directory:" + err.Error())
	}
	return err
}

func (filesystemManager *FilesystemManager) CreateMediaDirectory(directoryName string) error {
	err := filesystemManager.CreateOutputDirectory(filesystemManager.mediaDir + "/" + directoryName)
	if err != nil {
		log.Fatalln("Error creating media directory:" + err.Error())
	}
	return err
}
