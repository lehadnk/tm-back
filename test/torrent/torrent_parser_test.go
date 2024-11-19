package torrent

import (
	_ "regexp"
	"testing"
	torrent_domain "tm/src/torrent/domain"
)

func TestGetTorrentName(t *testing.T) {
	torrentParser := torrent_domain.NewTorrentParser()

	filePath := "../test.torrent"

	torrentName := "Ochen.strashnoe.kino.2000.RUS.BDRip.XviD.AC3.-HQCLUB"
	torrentNameBencode, err := torrentParser.GetTorrentNameFromBencode(filePath)
	if err != nil {
		t.Fatal("Error getting torrent name from bencode")
	}
	if torrentName != torrentNameBencode {
		t.Fatal("Expected", torrentName, "got", torrentNameBencode)
	}
}
