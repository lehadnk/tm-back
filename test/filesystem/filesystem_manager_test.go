package filesystem

import (
	"reflect"
	"testing"
	"tm/src/filesystem"
	"tm/src/filesystem/domain"
	"tm/test"
)

func TestSaveTorrentFile(t *testing.T) {
	file, _ := domain.ReadFile("../test.torrent")
	if file == nil {
		t.Errorf("Error reading file")
	}

	testFilePath := "/tmp/" + test.StringWithCharset(10, "abcdefghijklmnopqrstuvwxyz") + ".torrent"
	err := filesystem.NewFilesystemService().SaveTorrentFile(file, testFilePath)
	if err != nil {
		t.Errorf("Error while saving file")
	}

	savedAndReadFile, _ := domain.ReadFile(testFilePath)
	if savedAndReadFile == nil {
		t.Errorf("Error reading file")
	}

	if !reflect.DeepEqual(file, savedAndReadFile) {
		t.Errorf("Files are not equal")
	}
}
