package filesystem

import (
	"os"
	"reflect"
	"testing"
	"tm/src/filesystem"
	"tm/src/filesystem/domain"
	"tm/test"
)

func TestSaveTorrentFile(t *testing.T) {
	filesystemManager := domain.NewFilesystemManager("/tmp", "/tmp", "/tmp")
	filesystemService := filesystem.NewFilesystemService(filesystemManager)

	file, _ := filesystemManager.ReadFile("../test.torrent")
	if file == nil {
		t.Errorf("Error reading file")
	}

	testFilePath := "/tmp/" + test.StringWithCharset(10, "abcdefghijklmnopqrstuvwxyz") + ".torrent"
	err := filesystemService.SaveTorrentFile(file, testFilePath)
	if err != nil {
		t.Errorf("Error while saving file")
	}

	savedAndReadFile, _ := filesystemManager.ReadFile(testFilePath)
	if savedAndReadFile == nil {
		t.Errorf("Error reading file")
	}

	if !reflect.DeepEqual(file, savedAndReadFile) {
		t.Errorf("Files are not equal")
	}
}

func TestCreateOutputDirectory(t *testing.T) {
	var directoryPath = "/tmp/" + test.StringWithCharset(10, "abcdefghijklmnopqrstuvwxyz")
	filesystemManager := domain.NewFilesystemManager("/tmp", "/tmp", "/tmp")

	err := filesystemManager.CreateOutputDirectory(directoryPath)
	if err != nil {
		t.Errorf("Error creating directory")
	}
	_, err = os.Stat(directoryPath)
	if os.IsNotExist(err) {
		t.Errorf("Directory does not exist")
	}
}

func TestMoveFile(t *testing.T) {
	sourcePath := "/tmp/" + test.StringWithCharset(10, "abcdefghijklmnopqrstuvwxyz")
	destinationPath := "/tmp/" + test.StringWithCharset(10, "abcdefghijklmnopqrstuvwxyz1234567890")

	filesystemManager := domain.NewFilesystemManager("/tmp", "/tmp", "/tmp")
	filesystemService := filesystem.NewFilesystemService(filesystemManager)

	err := filesystemManager.CreateOutputDirectory(sourcePath)
	if err != nil {
		t.Errorf("Error creating source directory")
	}
	err = filesystemManager.CreateOutputDirectory(destinationPath)
	if err != nil {
		t.Errorf("Error creating destination directory")
	}
	file, _ := filesystemManager.ReadFile("../test.torrent")
	if file == nil {
		t.Errorf("Error reading file")
	}
	err = filesystemService.SaveTorrentFile(file, sourcePath+"/newtest.torrent")
	if err != nil {
		t.Errorf("Error while saving file")
	}

	err = filesystemManager.MoveFile(sourcePath+"/newtest.torrent", destinationPath+"/newtest.torrent")
	if err != nil {
		t.Errorf("Error while moving file")
	}
	_, err = os.Stat(destinationPath + "/newtest.torrent")
	if os.IsNotExist(err) {
		t.Errorf("Directory does not exist")
	}
}
