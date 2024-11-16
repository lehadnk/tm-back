package main

import (
	"time"
	"tm/src/authentication"
	"tm/src/authentication/domain"
	"tm/src/http"
	"tm/src/http/communication"
	"tm/src/user"
	"tm/src/user/persistence"
)

func main() {
	userDao := persistence.NewUserDao()
	userService := user.NewUserService(userDao)
	jwtManager := domain.NewJwtManager(userService)
	authService := authentication.NewAuthService(userService, jwtManager)

	httpServer := communication.NewHttpServer(authService)
	httpService := http.NewHttpService(httpServer)
	httpService.Start()

	for true {
		time.Sleep(10000)
	}
}
