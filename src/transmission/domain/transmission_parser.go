package domain

import (
	"awesomeProject/src/transmission/dto"
	"regexp"
	_ "regexp"
	"strconv"
	"strings"
)

type TransmissionParser struct {
}

func (parser *TransmissionParser) SeparateToLines(stringToSplit string) []*dto.TransmissionTorrent {
	separatedResult := strings.SplitAfter(stringToSplit, "\n")

	var lines []*dto.TransmissionTorrent
	for i := 1; i < len(separatedResult)-1; i++ {
		line := parseLine(separatedResult[i])
		lines = append(lines, line)
	}
	return lines
}

func parseLine(line string) *dto.TransmissionTorrent {
	re := regexp.MustCompile(`\s{2,}`)
	parts := re.Split(strings.TrimSpace(line), -1)

	if len(parts) < 8 {
		return nil
	}

	name := strings.Join(parts[8:], " ")

	id, _ := strconv.Atoi(parts[0][:len(parts[0])-1])
	ratio, _ := strconv.Atoi(parts[6])

	return &dto.TransmissionTorrent{
		Id:     id,
		Done:   parts[1],
		Have:   parts[2],
		ETA:    parts[3],
		Up:     parts[4],
		Down:   parts[5],
		Ratio:  ratio,
		Status: parts[7],
		Name:   name,
	}
}
