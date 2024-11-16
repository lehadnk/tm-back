package communication

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"tm/src/authentication"
	http_dto "tm/src/http/dto"
	user_dto "tm/src/user/dto"
)

type HttpServer struct {
	authService *authentication.AuthService
}

func NewHttpServer(authService *authentication.AuthService) *HttpServer {
	return &HttpServer{
		authService: authService,
	}
}

func (s *HttpServer) jsonResponse(w http.ResponseWriter, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}

func (s *HttpServer) getAuthenticatedUser(w http.ResponseWriter, r *http.Request) *user_dto.User {
	authHeaderValue := r.Header.Get("Authorization")

	if authHeaderValue[:6] == "Bearer" {
		authHeaderValue = authHeaderValue[7:]
	}

	user, err := s.authService.GetCurrentUser(authHeaderValue)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return nil
	}

	return user
}

func (s *HttpServer) handleCurrent(w http.ResponseWriter, r *http.Request) {
	user := s.getAuthenticatedUser(w, r)
	if user == nil {
		return
	}

	s.jsonResponse(w, user)
}

func (s *HttpServer) decodeRequestPayload(w http.ResponseWriter, r *http.Request, obj any) error {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&obj)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return errors.New("Could not decode the request object")
	}

	return nil
}

func (s *HttpServer) handleLogin(w http.ResponseWriter, r *http.Request) {
	var req http_dto.LoginRequest
	err := s.decodeRequestPayload(w, r, &req)
	if err != nil {
		return
	}

	user, token, err := s.authService.Login(req.Email, req.Password)
	resp := http_dto.LoginResponse{
		IsSuccess:           err != nil,
		AuthenticationToken: token,
		User:                user,
	}

	s.jsonResponse(w, resp)
}

func (server *HttpServer) Start() {
	http.HandleFunc("POST /login", server.handleLogin)
	http.HandleFunc("GET /user/current", server.handleCurrent)

	fmt.Println("Starting http server at :8080...")
	go http.ListenAndServe(":8080", nil)
	fmt.Println("http server started at :8080")
}
