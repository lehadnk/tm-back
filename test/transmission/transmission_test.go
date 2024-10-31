package transmission

import (
	"awesomeProject/src/transmission/cli_runner"
	"awesomeProject/src/transmission/transmission_client"
	"log"
	"testing"
)

func TestSomething(t *testing.T) {
	runner := cli_runner.CliRunner{}
	transmissionClient := transmission_client.NewTransmissionClient(runner)

	//runner.On("transmission-remote -l", "", "{status: Error}")
	testResult := transmissionClient.AddTorrentFile("https://test", "gttps://")

	if testResult == false {
		log.Fatalln("Torrent is not added")
	}
}
