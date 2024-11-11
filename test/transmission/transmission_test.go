package transmission

import (
	"log"
	"testing"
	cli_domain "tm/src/cli/domain"
	transmission_domain "tm/src/transmission/domain"
)

func TestAddTorrentFile(t *testing.T) {
	runner := &cli_domain.MockCliRunner{}
	transmissionClient := transmission_domain.NewTransmissionClient(runner)

	runner.On("transmission-remote", []string{"-a"}, "localhost:9091/transmission/rpc/ responded: success", nil)
	addTorrentResult := transmissionClient.AddTorrentFile("/media/torrents/1.torrent", "/media/downloads/1")
	if addTorrentResult == false {
		log.Fatalln("Torrent was not added")
	}

	if !runner.WasCommandRun("transmission-remote -a /media/torrents/1.torrent -w /media/downloads/1") {
		log.Fatalln("Torrent add command was not run")
	}
}

func TestSeparate(t *testing.T) {
	var inputString = "   ID   Done       Have  ETA           Up    Down  Ratio  Status       Name\n     1*  100%    2.20 GB  Done         0.0     0.0    1.4  Stopped      Ochen.strashnoe.kino.2000.RUS.BDRip.XviD.AC3.-HQCLUB\n     2*  100%   11.24 GB  Done         0.0     0.0   13.5  Stopped      Catch.Me.If.You.Can.2002.WEB-DL.1080p.Open.Matte.mkv\n     3*  100%   11.47 GB  Done         0.0     0.0    0.1  Stopped      Margin.Call.2011.BluRay.1080p.Rus.Eng.DTS.x264-CHD.mkv\n     4*  100%   29.33 GB  Done         0.0     0.0    0.0  Stopped      Король Стейтен-Айленда.2020.WEB-DL.2160p.mkv\n     5*  100%   10.84 GB  Done         0.0     0.0    0.9  Stopped      Jay.and.Silent.Bob.Strike.Back.2001.Open.Matte.1080p.WEB-DL.mkv\n     6*  100%    4.23 GB  Done         0.0     0.0    0.0  Stopped      99 francs '07 [Wrnr].mkv\n     7*  100%    3.50 GB  Done         0.0     0.0   11.6  Stopped      Click\n     8   100%   13.41 GB  Done         0.0     0.0    4.1  Idle         Black.Coal.Thin.Ice.2014.BDRip.1080p.Sub.Rus.Eng.mkv\n     9*  100%    3.01 GB  Done         0.0     0.0    0.0  Stopped      COMING TO AMERICA_1988_BDRIP_СУКА-ПАДЛА_МИХАЛЁВ_ENG_KBC.mkv\n    10*  100%    5.56 GB  Done         0.0     0.0    0.0  Stopped      Robot Chicken (Seasons 1, 2, 3) AVO (Goblin)\n    11*  100%   28.91 GB  Done         0.0     0.0    0.0  Stopped      Dazed.and.Confused.1993.1080p.Bluray.AVC.Remux.mkv\n    12*  100%    8.19 GB  Done         0.0     0.0    0.0  Stopped      Scary.Movie.2000.720p.BluRay.5xRus.Eng.HDCLUB-SbR.mkv\n    13*  100%   19.90 GB  Done         0.0     0.0    0.0  Stopped      Scary.Movie.1991.BDRemux.1080p.mkv\n    14*  100%    4.17 GB  Done         0.0     0.0    3.2  Stopped      *Batteries Not Included (HD).m4v\n"

	parser := transmission_domain.TransmissionParser{}
	testResult := parser.SeparateToLines(inputString)

	if len(testResult) != 14 {
		t.Fatal("Length is not expected")
	}

	if testResult[0].Id != 1 ||
		testResult[0].Done != "100%" ||
		testResult[0].Down != "0.0" ||
		testResult[0].ETA != "Done" ||
		testResult[0].Name != "Ochen.strashnoe.kino.2000.RUS.BDRip.XviD.AC3.-HQCLUB" {
		t.Fatal("Data is not exist or not expected")
	}
}

func TestGetTorrentFileList(t *testing.T) {
	runner := &cli_domain.MockCliRunner{}
	transmissionClient := transmission_domain.NewTransmissionClient(runner)

	runner.On("transmission-remote", []string{"-t"}, "Ochen.strashnoe.kino.2000.RUS.BDRip.XviD.AC3.-HQCLUB (6 files):\n  #  Done Priority Get      Size  Name\n  0: 100% Normal   Yes 1.56 GB    Ochen.strashnoe.kino.2000.RUS.BDRip.XviD.AC3.-HQCLUB/Ochen.strashnoe.kino.2000.RUS.BDRip.XviD.AC3.-HQCLUB.avi\n  1: 100% Normal   Yes 254.3 MB   Ochen.strashnoe.kino.2000.RUS.BDRip.XviD.AC3.-HQCLUB/Ochen.strashnoe.kino.2000.RUS.BDRip.XviD.AC3.-HQCLUB.AVO.GRANKIN.ac3\n  2: 100% Normal   Yes 127.0 MB   Ochen.strashnoe.kino.2000.RUS.BDRip.XviD.AC3.-HQCLUB/Ochen.strashnoe.kino.2000.RUS.BDRip.XviD.AC3.-HQCLUB.AVO.SANAEV.ac3\n  3: 100% Normal   Yes 254.3 MB   Ochen.strashnoe.kino.2000.RUS.BDRip.XviD.AC3.-HQCLUB/Ochen.strashnoe.kino.2000.RUS.BDRip.XviD.AC3.-HQCLUB.ENG.ac3\n  4: 100% Normal   Yes 98.41 kB   Ochen.strashnoe.kino.2000.RUS.BDRip.XviD.AC3.-HQCLUB/Ochen.strashnoe.kino.2000.RUS.BDRip.XviD.AC3.-HQCLUB.ENG.SDH.srt\n  5: 100% Normal   Yes 70.94 kB   Ochen.strashnoe.kino.2000.RUS.BDRip.XviD.AC3.-HQCLUB/Ochen.strashnoe.kino.2000.RUS.BDRip.XviD.AC3.-HQCLUB.RUS.srt", nil)
	fileList := transmissionClient.GetFileList(1)
	if !runner.WasCommandRun("transmission-remote -t 1 -f") {
		log.Fatalln("List torrent files command was not run")
	}

	if len(fileList) != 6 {
		log.Fatalln("Returned list doesn't contain all files")
	}

	if fileList[0].Id != 0 {
		log.Fatalln("First list element has an incorrect ID")
	}

	if fileList[5].Name != "Ochen.strashnoe.kino.2000.RUS.BDRip.XviD.AC3.-HQCLUB/Ochen.strashnoe.kino.2000.RUS.BDRip.XviD.AC3.-HQCLUB.RUS.srt" {
		log.Fatalln("Last list element has an incorrect name")
	}
}
