package http

import (
	"encoding/json"
	"net/http"
	"strconv"

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
	router.HandleFunc("/api/v1/posts", handler.CreatePost).Methods("POST")
	router.HandleFunc("/api/v1/posts/{id}", handler.UpdatePost).Methods("PUT")
	router.HandleFunc("/api/v1/posts/{id}", handler.DeletePost).Methods("DELETE")
	router.HandleFunc("/api/v1/posts/{id}", handler.GetPost).Methods("GET")
	router.HandleFunc("/api/v1/posts", handler.GetPostsPaged).Methods("GET")
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
	var post entity.Post
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		JSONResponse(w, http.StatusBadRequest, "error", "Invalid post ID", nil)
		return
	}
	json.NewDecoder(r.Body).Decode(&post)
	post.ID = id
	data, err := h.postUseCase.UpdatePost(&post)
	if err != nil {
		JSONResponse(w, http.StatusInternalServerError, "error", err.Error(), nil)
		return
	}
	JSONResponse(w, http.StatusOK, "success", "Post updated successfully", data)
}

func (h *PostHandler) DeletePost(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		JSONResponse(w, http.StatusBadRequest, "error", "Invalid post ID", nil)
		return
	}
	err = h.postUseCase.DeletePost(id)
	if err != nil {
		JSONResponse(w, http.StatusInternalServerError, "error", err.Error(), nil)
		return
	}
	JSONResponse(w, http.StatusOK, "success", "Post deleted successfully", nil)
}

func (h *PostHandler) GetPost(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		JSONResponse(w, http.StatusBadRequest, "error", "Invalid post ID", nil)
		return
	}
	post, err := h.postUseCase.GetPostByID(id)
	if err != nil {
		JSONResponse(w, http.StatusInternalServerError, "error", err.Error(), nil)
		return
	}
	JSONResponse(w, http.StatusOK, "success", "Post retrieved successfully", post)
}

func (h *PostHandler) GetPostsPaged(w http.ResponseWriter, r *http.Request) {
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	size, _ := strconv.Atoi(r.URL.Query().Get("size"))
	posts, err := h.postUseCase.GetPostsPaged(page, size)
	if err != nil {
		JSONResponse(w, http.StatusInternalServerError, "error", err.Error(), nil)
		return
	}
	JSONResponse(w, http.StatusOK, "success", "Posts retrieved successfully", posts)
}
