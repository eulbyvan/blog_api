package http

import (
	"encoding/json"
	"net/http"

	"github.com/eulbyvan/blog_api/internal/entity"
	"github.com/eulbyvan/blog_api/internal/usecase"
	"github.com/gorilla/mux"
)

type PostHandler struct {
	postUseCase usecase.PostUseCase
}

// routes
func NewPostHandler(router *mux.Router, pu usecase.PostUseCase) {
	handler := &PostHandler{postUseCase: pu}
	router.HandleFunc("/v1/api/posts", handler.CreatePost).Methods("POST")
	router.HandleFunc("/v1/api/posts/{id}", handler.UpdatePost).Methods("PUT")
	router.HandleFunc("/v1/api/posts/{id}", handler.DeletePost).Methods("DELETE")
	router.HandleFunc("/v1/api/posts/{id}", handler.GetPost).Methods("GET")
	router.HandleFunc("/v1/api/posts", handler.GetPostsPaged).Methods("GET")
}

func (h *PostHandler) CreatePost(w http.ResponseWriter, r *http.Request) {
	var post entity.Post
	json.NewDecoder(r.Body).Decode(&post)
	data, err := h.postUseCase.CreatePost(&post)
	if err != nil {
		JSONResponse(w, http.StatusInternalServerError, "error", err.Error(), nil)
		return
	}
	JSONResponse(w, http.StatusCreated, "success", "Post created successfully", data)
}

func (h *PostHandler) UpdatePost(w http.ResponseWriter, r *http.Request) {
	// TODO
}

func (h *PostHandler) DeletePost(w http.ResponseWriter, r *http.Request) {
	// TODO
}

func (h *PostHandler) GetPost(w http.ResponseWriter, r *http.Request) {
	// TODO
}

func (h *PostHandler) GetPostsPaged(w http.ResponseWriter, r *http.Request) {
	// TODO
}
