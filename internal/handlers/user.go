package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/nkroshechkin/micro-blog-zero/internal/service"
)

type UserHandler struct {
	service service.UserService
}

func NewUserHandler(s service.UserService) *UserHandler {
	return &UserHandler{service: s}
}

func (h *UserHandler) GetUsers(res http.ResponseWriter, req *http.Request) {
	if id := req.URL.Query().Get("id"); id != "" {
		user, err := h.service.GetUser(id)

		if err != nil {
			http.Error(res, err.Error(), http.StatusBadRequest)
			return
		}

		json.NewEncoder(res).Encode(user)
		return
	}

	user, err := h.service.GetAllUser()
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(res).Encode(user)

}

func (h *UserHandler) CreateUser(res http.ResponseWriter, req *http.Request) {
	var request struct {
		Username string `json:"username"`
	}

	if err := json.NewDecoder(req.Body).Decode(&request); err != nil {
		http.Error(res, "Неверное тело запроса", http.StatusBadRequest)
		return
	}

	idUser, err := h.service.CreateUser(request.Username)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
	}

	res.WriteHeader(http.StatusCreated)
	res.Write([]byte(idUser))
}
