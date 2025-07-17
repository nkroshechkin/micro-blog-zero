package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/nkroshechkin/micro-blog-zero/internal/service"
)

type PostHandler struct {
	service service.PostService
}

func NewPostHandler(p service.PostService) *PostHandler {
	return &PostHandler{service: p}
}

func (h *PostHandler) GetPosts(res http.ResponseWriter, req *http.Request) {
	if id := req.URL.Query().Get("id"); id != "" {
		post, err := h.service.GetPost(id)
		if err != nil {
			http.Error(res, err.Error(), http.StatusBadRequest)
			return
		}

		res.WriteHeader(http.StatusFound)
		json.NewEncoder(res).Encode(post)
		return
	}
	post, err := h.service.GetAllPost()
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	res.WriteHeader(http.StatusFound)
	json.NewEncoder(res).Encode(post)

}
func (h *PostHandler) CreatePosts(res http.ResponseWriter, req *http.Request) {
	var request struct {
		AuthorId string `json:"authorId"`
		Text     string `json:"text"`
	}

	if err := json.NewDecoder(req.Body).Decode(&request); err != nil {
		http.Error(res, "Неверное тело запроса", http.StatusBadRequest)
		return
	}

	idPost, err := h.service.CreatePost(request.AuthorId, request.Text)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
	}

	res.WriteHeader(http.StatusCreated)
	res.Write([]byte(idPost))

}
func (h *PostHandler) LikePosts(res http.ResponseWriter, req *http.Request) {
	pathParts := strings.Split(req.URL.Path, "/")

	postId := pathParts[2]

	var request struct {
		UserId string `json:"userId"`
	}

	if err := json.NewDecoder(req.Body).Decode(&request); err != nil || postId == "" {
		http.Error(res, "Неверный запрос", http.StatusBadRequest)
		return
	}

	status, err := h.service.LikePost(request.UserId, postId)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
	}

	res.WriteHeader(http.StatusOK)
	res.Write([]byte(status))
}
