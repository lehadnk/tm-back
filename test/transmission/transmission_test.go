package transmission

import (
	"awesomeProject/src/transmission/cli_runner"
	"awesomeProject/src/transmission/transmission_client"
	"fmt"
	"log"
	"testing"
)

func TestAddTorrentFile(t *testing.T) {
	runner := cli_runner.CliRunner{}
	transmissionClient := transmission_client.NewTransmissionClient(runner)

	//runner.On("transmission-remote -l", "", "{status: Error}")
	testResult := transmissionClient.AddTorrentFile("https://test", "gttps://")

	if testResult == false {
		log.Fatalln("Torrent is not added")
	}
}


func TestParser(t *testing.T) {
	line := "     2*  100%   11.24 GB  Done         0.0     0.0   13.5  Stopped      Catch.Me.If.You.Can.2002.WEB-DL.1080p.Open.Matte.mkv\n"
	torrent := transmission_client..(line)
	fmt.Println(torrent.ETA)
	fmt.Println(torrent.Status)
	fmt.Println(torrent.ID)
}

