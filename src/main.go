package main

import (
	"log"
	"os"
	"tm/src/authentication"
	authentication_domain "tm/src/authentication/domain"
	cli_domain "tm/src/cli/domain"
	"tm/src/filesystem"
	"tm/src/filesystem/domain"
	"tm/src/http"
	http_communication "tm/src/http/communication"
	"tm/src/torrent"
	torrent_domain "tm/src/torrent/domain"
	torrent_persistence "tm/src/torrent/persistence"
	"tm/src/transmission"
	transmission_communication "tm/src/transmission/communication"
	transmission_domain "tm/src/transmission/domain"
	"tm/src/user"
	user_persistence "tm/src/user/persistence"
)

func main() {
	torrentFileDir := os.Getenv("TM_TORRENT_FILE_DIR")
	if torrentFileDir == "" {
		log.Fatalln("No output directory specified for writing .torrent files. Please set up the TM_TORRENT_FILE_DIR environment variable.")
	}

	torrentOutputDir := os.Getenv("TM_TORRENT_OUTPUT_DIR")
	if torrentOutputDir == "" {
		log.Fatalln("No output directory specified for writing transmission download files. Please set up the TM_TORRENT_OUTPUT_DIR environment variable.")
	}

	mediaDir := os.Getenv("TM_MEDIA_DIR")
	if mediaDir == "" {
		log.Fatalln("No output directory specified as media directory. Please set up the TM_MEDIA_DIR environment variable.")
	}

	userDao := user_persistence.NewUserDao()
	userService := user.NewUserService(userDao)

	jwtManager := authentication_domain.NewJwtManager(userService)
	authService := authentication.NewAuthService(userService, jwtManager)

	filesystemManager := domain.NewFilesystemManager(torrentFileDir, torrentOutputDir, mediaDir)
	filesystemService := filesystem.NewFilesystemService(filesystemManager)

	torrentDao := torrent_persistence.NewTorrentDao()
	torrentParser := torrent_domain.NewTorrentParser()
	cliRunner := cli_domain.CliRunner{}
	transmissionClient := transmission_domain.NewTransmissionClient(cliRunner)
	transmissionService := transmission.NewTransmissionService(transmissionClient)
	torrentManager := torrent_domain.NewTorrentManager(torrentDao, torrentParser, transmissionService, filesystemService)
	torrentService := torrent.NewTorrentService(torrentManager)

	httpServer := http_communication.NewHttpServer(authService, userService, torrentService)
	httpService := http.NewHttpService(httpServer)
	httpService.Start()

	downloadedTorrentsScanner := transmission_communication.NewDownloadedTorrentsScanner(torrentService, filesystemService, &cliRunner)
	downloadedTorrentsScanner.Start()

	select {}
}
