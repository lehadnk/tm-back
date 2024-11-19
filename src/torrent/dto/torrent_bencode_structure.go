package dto

type TorrentBencodeStructure struct {
	Info struct {
		Name string `bencode:"name"`
	} `bencode:"info"`
}
