package main

import (
	"tm/src/authentication"
	"tm/src/authentication/domain"
	"tm/src/http"
	http_communication "tm/src/http/communication"
	transmission_communication "tm/src/transmission/communication"
	"tm/src/user"
	"tm/src/user/persistence"
)

func main() {
	userDao := persistence.NewUserDao()
	userService := user.NewUserService(userDao)
	jwtManager := domain.NewJwtManager(userService)
	authService := authentication.NewAuthService(userService, jwtManager)

	httpServer := http_communication.NewHttpServer(authService)
	httpService := http.NewHttpService(httpServer)
	httpService.Start()

	downloadedTorrentsScanner := transmission_communication.NewDownloadedTorrentsScanner()
	downloadedTorrentsScanner.Start()

	for true {

	}
}
