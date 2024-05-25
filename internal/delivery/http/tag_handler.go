package http

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/eulbyvan/blog_api/internal/entity"
	"github.com/eulbyvan/blog_api/internal/usecase"
	"github.com/gorilla/mux"
)

type TagHandler struct {
	tagUseCase usecase.TagUseCase
}

// routes
func NewTagHandler(router *mux.Router, tu usecase.TagUseCase) {
	handler := &TagHandler{tagUseCase: tu}
	router.HandleFunc("/api/v1/tags", handler.CreateTag).Methods("POST")
	router.HandleFunc("/api/v1/tags/{id}", handler.UpdateTag).Methods("PUT")
	router.HandleFunc("/api/v1/tags/{id}", handler.DeleteTag).Methods("DELETE")
	router.HandleFunc("/api/v1/tags/{id}", handler.GetTag).Methods("GET")
	router.HandleFunc("/api/v1/tags", handler.GetAllTags).Methods("GET")
}

func (h *TagHandler) CreateTag(w http.ResponseWriter, r *http.Request) {
	var tag entity.Tag
	json.NewDecoder(r.Body).Decode(&tag)
	data, err := h.tagUseCase.CreateTag(&tag)
	if err != nil {
		JSONResponse(w, http.StatusInternalServerError, "error", err.Error(), nil)
		return
	}
	JSONResponse(w, http.StatusCreated, "success", "Tag created successfully", data)
}

func (h *TagHandler) UpdateTag(w http.ResponseWriter, r *http.Request) {
	var tag entity.Tag
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		JSONResponse(w, http.StatusBadRequest, "error", "Invalid tag ID", nil)
		return
	}
	json.NewDecoder(r.Body).Decode(&tag)
	tag.ID = id
	data, err := h.tagUseCase.UpdateTag(&tag)
	if err != nil {
		JSONResponse(w, http.StatusInternalServerError, "error", err.Error(), nil)
		return
	}
	JSONResponse(w, http.StatusOK, "success", "Tag updated successfully", data)
}

func (h *TagHandler) DeleteTag(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		JSONResponse(w, http.StatusBadRequest, "error", "Invalid tag ID", nil)
		return
	}
	err = h.tagUseCase.DeleteTag(id)
	if err != nil {
		JSONResponse(w, http.StatusInternalServerError, "error", err.Error(), nil)
		return
	}
	JSONResponse(w, http.StatusOK, "success", "Tag deleted successfully", nil)
}

func (h *TagHandler) GetTag(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		JSONResponse(w, http.StatusBadRequest, "error", "Invalid tag ID", nil)
		return
	}
	tag, err := h.tagUseCase.GetTagByID(id)
	if err != nil {
		JSONResponse(w, http.StatusInternalServerError, "error", err.Error(), nil)
		return
	}
	JSONResponse(w, http.StatusOK, "success", "Tag retrieved successfully", tag)
}

func (h *TagHandler) GetAllTags(w http.ResponseWriter, r *http.Request) {
	tags, err := h.tagUseCase.GetAllTags()
	if err != nil {
		JSONResponse(w, http.StatusInternalServerError, "error", err.Error(), nil)
		return
	}
	JSONResponse(w, http.StatusOK, "success", "Tags retrieved successfully", tags)
}
