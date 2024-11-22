package communication

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/mail"
	"strconv"
	"strings"
	"tm/src/authentication"
	http_dto "tm/src/http/dto"
	"tm/src/torrent"
	"tm/src/user"
	user_dto "tm/src/user/dto"
)

type HttpServer struct {
	authService    *authentication.AuthService
	userService    *user.UserService
	torrentService *torrent.TorrentService
}

func NewHttpServer(
	authService *authentication.AuthService,
	userService *user.UserService,
	torrentService *torrent.TorrentService,

) *HttpServer {
	return &HttpServer{
		authService:    authService,
		userService:    userService,
		torrentService: torrentService,
	}
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (s *HttpServer) jsonResponse(w http.ResponseWriter, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}

func (s *HttpServer) requireAuthenticatedUser(w http.ResponseWriter, r *http.Request) *user_dto.User {
	authHeaderValue := r.Header.Get("Authorization")

	if len(authHeaderValue) > 5 && authHeaderValue[:6] == "Bearer" {
		authHeaderValue = authHeaderValue[7:]
	}

	user, err := s.authService.GetCurrentUser(authHeaderValue)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return nil
	}

	return user
}

func (s *HttpServer) requireAuthenticatedAdmin(w http.ResponseWriter, r *http.Request) *user_dto.User {
	user := s.requireAuthenticatedUser(w, r)
	if user == nil {
		return nil
	}

	if user.Role != "admin" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return nil
	}

	return user
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

func (s *HttpServer) getNumericUrlParam(w http.ResponseWriter, r *http.Request, name string) (int, error) {
	strValue := r.URL.Query().Get(name)
	if strValue == "" {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return 0, errors.New("Value cannot be empty")
	}

	intValue, err := strconv.Atoi(strValue)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return 0, err
	}

	return intValue, nil
}

func (s *HttpServer) getMultipartFormDataFile(r *http.Request, w http.ResponseWriter, field string) ([]byte, error) {
	reader, err := r.MultipartReader()
	if err != nil {
		http.Error(w, "Error creating multipart reader", http.StatusInternalServerError)
		return nil, err
	}

	var fileBytes []byte
	for {
		part, err := reader.NextPart()
		if err == io.EOF {
			break
		}
		if err != nil {
			http.Error(w, "Error reading multipart data", http.StatusInternalServerError)
			return nil, err
		}

		if part.FormName() == field {
			buf := bytes.NewBuffer(nil)
			if _, err := io.Copy(buf, part); err != nil {
				http.Error(w, "Error reading file content", http.StatusInternalServerError)
				return nil, err
			}
			fileBytes = buf.Bytes()
		}
		part.Close()
	}

	return fileBytes, nil
}

func (s *HttpServer) handleCurrentUser(w http.ResponseWriter, r *http.Request) {
	user := s.requireAuthenticatedUser(w, r)
	if user == nil {
		return
	}

	s.jsonResponse(w, user)
}

func (s *HttpServer) handleLogin(w http.ResponseWriter, r *http.Request) {
	var req http_dto.LoginRequest
	err := s.decodeRequestPayload(w, r, &req)
	if err != nil {
		return
	}

	user, token, err := s.authService.Login(req.Email, req.Password)
	resp := http_dto.LoginResponse{
		IsSuccess:           err == nil,
		AuthenticationToken: token,
		User:                user,
	}

	s.jsonResponse(w, resp)
}

func (s *HttpServer) handleGetUsers(w http.ResponseWriter, r *http.Request) {
	user := s.requireAuthenticatedAdmin(w, r)
	if user == nil {
		return
	}

	sort := r.URL.Query().Get("sort")
	if sort != "id" && sort != "name" {
		http.Error(w, "Bad Request: sort should be 'id' or 'name'", http.StatusBadRequest)
		return
	}

	limit, err := s.getNumericUrlParam(w, r, "limit")
	if err != nil {
		http.Error(w, "Bad Request: limit should be a numeric value", http.StatusBadRequest)
		return
	}

	page, err := s.getNumericUrlParam(w, r, "page")
	if err != nil {
		http.Error(w, "Bad Request: page should be numeric value", http.StatusBadRequest)
		return
	}

	usersList := s.userService.GetUsersList(sort, page, limit)

	s.jsonResponse(w, usersList)
}

func (s *HttpServer) handleGetUser(w http.ResponseWriter, r *http.Request) {
	admin := s.requireAuthenticatedAdmin(w, r)
	if admin == nil {
		return
	}

	userId, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "Bad request: user id should be int value", http.StatusBadRequest)
		return
	}

	readUser := s.userService.GetUserById(userId)
	if readUser == nil {
		http.Error(w, "User does not exist", http.StatusNotFound)
		return
	}

	s.jsonResponse(w, readUser)
}

func (s *HttpServer) handleAddUser(w http.ResponseWriter, r *http.Request) {
	user := s.requireAuthenticatedAdmin(w, r)
	if user == nil {
		return
	}

	var req http_dto.SaveUserRequest
	err := s.decodeRequestPayload(w, r, &req)
	if err != nil {
		return
	}

	_, err = mail.ParseAddress(req.Email)
	if err != nil {
		http.Error(w, "Bad request: email should be a correct email", http.StatusBadRequest)
		return
	}
	if len(req.Password) < 6 {
		http.Error(w, "Bad request: password should be at least 6 characters long", http.StatusBadRequest)
		return
	}
	if len(strings.TrimSpace(req.Name)) == 0 {
		http.Error(w, "Bad request: name should not be empty", http.StatusBadRequest)
		return
	}
	if req.Role != "admin" && req.Role != "user" {
		http.Error(w, "Bad request: role should be either 'user' or 'admin'", http.StatusBadRequest)
		return
	}

	s.userService.CreateUser(req.Name, req.Email, req.Password, req.Role)
}

func (s *HttpServer) handleUpdateUser(w http.ResponseWriter, r *http.Request) {
	admin := s.requireAuthenticatedAdmin(w, r)
	if admin == nil {
		return
	}

	var req http_dto.SaveUserRequest
	err := s.decodeRequestPayload(w, r, &req)
	if err != nil {
		return
	}

	_, err = mail.ParseAddress(req.Email)
	if err != nil {
		http.Error(w, "Bad request: email should be a correct email", http.StatusBadRequest)
		return
	}
	if len(req.Password) < 6 && req.Password != "" {
		http.Error(w, "Bad request: password should be at least 6 characters long", http.StatusBadRequest)
		return
	}
	if len(strings.TrimSpace(req.Name)) == 0 {
		http.Error(w, "Bad request: name should not be empty", http.StatusBadRequest)
		return
	}
	if req.Role != "admin" && req.Role != "user" {
		http.Error(w, "Bad request: role should be either 'user' or 'admin'", http.StatusBadRequest)
		return
	}

	userId, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "Bad request: user id should be int value", http.StatusBadRequest)
		return
	}

	editedUser := s.userService.GetUserById(userId)
	if editedUser == nil {
		http.Error(w, "User does not exist", http.StatusNotFound)
		return
	}

	s.userService.UpdateUser(userId, req.Name, req.Email, req.Password, req.Role)
}

func (s *HttpServer) handleDeleteUser(w http.ResponseWriter, r *http.Request) {
	user := s.requireAuthenticatedAdmin(w, r)
	if user == nil {
		return
	}

	userId, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "Bad request: user id should be int value", http.StatusBadRequest)
		return
	}

	s.userService.DeleteUser(userId)
}

func (s *HttpServer) handleTorrentList(w http.ResponseWriter, r *http.Request) {
	user := s.requireAuthenticatedUser(w, r)
	if user == nil {
		return
	}

	sort := r.URL.Query().Get("sort")
	if sort != "id" && sort != "name" {
		http.Error(w, "Bad Request: sort should be 'id' or 'name'", http.StatusBadRequest)
		return
	}

	page, err := s.getNumericUrlParam(w, r, "page")
	if err != nil {
		http.Error(w, "Bad Request: page should be a numeric value", http.StatusBadRequest)
		return
	}

	limit, err := s.getNumericUrlParam(w, r, "limit")
	if err != nil {
		http.Error(w, "Bad Request: pagesize should be numeric value", http.StatusBadRequest)
		return
	}

	torrentsList := s.torrentService.GetTorrentsList(sort, page, limit)
	s.jsonResponse(w, torrentsList)

}

func (s *HttpServer) handleAddTorrent(w http.ResponseWriter, r *http.Request) {
	user := s.requireAuthenticatedUser(w, r)
	if user == nil {
		return
	}

	file, err := s.getMultipartFormDataFile(r, w, "file")
	if err != nil {
		return
	}

	newTorrent, inputError, systemError := s.torrentService.AddTorrent(file)
	if systemError != nil {
		http.Error(w, "Error while adding torrent: "+err.Error(), http.StatusInternalServerError)
		return
	}
	if inputError != nil {
		http.Error(w, "Error while adding torrent: "+err.Error(), http.StatusBadRequest)
		return
	}
	
	s.jsonResponse(w, newTorrent)
}

func (s *HttpServer) handleDeleteTorrent(w http.ResponseWriter, r *http.Request) {
	user := s.requireAuthenticatedUser(w, r)
	if user == nil {
		return
	}

	torrentId, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "Bad request: user id should be int value", http.StatusBadRequest)
		return
	}

	err = s.torrentService.DeleteTorrent(torrentId)
	if err != nil {
		http.Error(w, "Error while removing torrent: "+err.Error(), http.StatusInternalServerError)
	}
}

func (s *HttpServer) handleGetSpace(w http.ResponseWriter, r *http.Request) {
	user := s.requireAuthenticatedUser(w, r)
	if user == nil {
		return
	}
}

func (server *HttpServer) Start() {
	http.HandleFunc("POST /login", server.handleLogin)

	http.HandleFunc("GET /users/current", server.handleCurrentUser)
	http.HandleFunc("GET /users", server.handleGetUsers)
	http.HandleFunc("GET /users/{id}", server.handleGetUser)
	http.HandleFunc("POST /users", server.handleAddUser)
	http.HandleFunc("PUT /users/{id}", server.handleUpdateUser)
	http.HandleFunc("DELETE /users/{id}", server.handleDeleteUser)

	http.HandleFunc("GET /torrents", server.handleTorrentList)
	http.HandleFunc("POST /torrents", server.handleAddTorrent)
	http.HandleFunc("DELETE /torrents/{id}", server.handleDeleteTorrent)

	http.HandleFunc("GET /space", server.handleGetSpace)

	fmt.Println("Starting http server at :8080...")
	handlerWithCors := corsMiddleware(http.DefaultServeMux)
	go http.ListenAndServe(":8080", handlerWithCors)
	fmt.Println("http server started at :8080")
}
