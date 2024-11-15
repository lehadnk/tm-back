package http

import "tm/src/http/communication"

type HttpService struct {
	httpServer *communication.HttpServer
}

func NewHttpService(httpServer *communication.HttpServer) *HttpService {
	return &HttpService{httpServer: httpServer}
}

func (httpService *HttpService) Start() {
	httpService.httpServer.Start()
}
