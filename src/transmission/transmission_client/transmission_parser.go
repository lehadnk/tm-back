package transmission_client

import (
	"regexp"
	"strings"
)

func separateToLines(stringToSplit string) []string {

	stringResult := strings.SplitAfter(stringToSplit, "\n")
	return stringResult
	}
}

// сделать другой торрент transmission torrent
func parseLine(line string) *TransmissionTorrent {
	re := regexp.MustCompile(`\s{2,}`)
	parts := re.Split(strings.TrimSpace(line), -1)

	if len(parts) < 8 {
		return nil
	}

	name := strings.Join(parts[8:], " ")

	return &TransmissionTorrent{
		Id:     parts[0][:len(parts[0])-1],
		Done:   parts[1],
		Have:   parts[2],
		ETA:    parts[3],
		Up:     parts[4],
		Down:   parts[5],
		Ratio:  parts[6],
		Status: parts[7],
		Name:   name,
	}
}
