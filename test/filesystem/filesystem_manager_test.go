package filesystem

import (
	"os"
	"reflect"
	"testing"
	"tm/src/common"
	"tm/src/filesystem"
	"tm/src/filesystem/domain"
)

func TestSaveTorrentFile(t *testing.T) {
	filesystemManager := domain.NewFilesystemManager("/tmp", "/tmp", "/tmp")
	filesystemService := filesystem.NewFilesystemService(filesystemManager)

	file, _ := filesystemManager.ReadFile("../test.torrent")
	if file == nil {
		t.Errorf("Error reading file")
	}

	testFileName := common.StringWithCharset(10, "abcdefghijklmnopqrstuvwxyz") + ".torrent"
	testFilePath, err := filesystemService.SaveTorrentFile(file, testFileName)
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
	var directoryPath = "/tmp/" + common.StringWithCharset(10, "abcdefghijklmnopqrstuvwxyz")
	filesystemManager := domain.NewFilesystemManager("/tmp", "/tmp", "/tmp")

	err := filesystemManager.CreateDirectory(directoryPath)
	if err != nil {
		t.Errorf("Error creating directory")
	}
	_, err = os.Stat(directoryPath)
	if os.IsNotExist(err) {
		t.Errorf("Directory does not exist")
	}
}

func TestMoveFile(t *testing.T) {
	filesystemManager := domain.NewFilesystemManager("/tmp", "/tmp", "/tmp")
	filesystemService := filesystem.NewFilesystemService(filesystemManager)

	file, _ := filesystemManager.ReadFile("../test.torrent")
	if file == nil {
		t.Errorf("Error reading file")
	}
	testFilePath, err := filesystemService.SaveTorrentFile(file, "newtest.torrent")
	if err != nil {
		t.Errorf("Error while saving file")
	}

	err = filesystemManager.MoveFile(testFilePath, "/tmp"+"/newtest.torrent")
	if err != nil {
		t.Errorf("Error while moving file")
	}
	_, err = os.Stat("/tmp" + "/newtest.torrent")
	if os.IsNotExist(err) {
		t.Errorf("Directory does not exist")
	}
}
