package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/knmsh08200/Blog_test/internal/model"
	"github.com/knmsh08200/Blog_test/internal/service"
)

type BlogHandler struct {
	Service *service.BlogService
}

func NewBlogHandler(s *service.BlogService) *BlogHandler {
	return &BlogHandler{Service: s}
}

func (h *BlogHandler) BlogListHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST,DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	switch r.Method {
	case http.MethodGet:
		h.handleGetList(w, r)
	case http.MethodPost:
		h.handlePostList(w, r)
	case http.MethodDelete:
		h.handleDelList(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *BlogHandler) handleGetList(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	lists, err := h.Service.GetAllBlogs(ctx)
	if err != nil {
		log.Printf("Error fetchinig blogs: %v", err)
		http.Error(w, "Error fetching data", http.StatusInternalServerError)
		return
	}

	jsonData, err := json.Marshal(lists)
	if err != nil {
		http.Error(w, "Error marshalling JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)

}

func (h *BlogHandler) handlePostList(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}

	var newList model.List
	if err := json.Unmarshal(body, &newList); err != nil {
		http.Error(w, "Error unmarshalling JSON", http.StatusBadRequest)
	}

	id, err := h.Service.CreateBlog(ctx, newList)
	if err != nil {
		http.Error(w, "Error inserting data", http.StatusInternalServerError)
		return
	}
	newList.ID = id

	w.WriteHeader(http.StatusCreated)
}

func (h *BlogHandler) handleDelList(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	idStr := strings.TrimPrefix(r.URL.Path, "/blog/list/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	rowsAffected, err := h.Service.DeleteBlog(ctx, id)
	if err != nil {
		http.Error(w, "Error deleting data", http.StatusInternalServerError)
	}

	if rowsAffected == 0 {
		http.Error(w, "Blog not found", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *BlogHandler) BlogCounterHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	params := r.URL.Query()

	userIDStr := params.Get("user_id")

	if userIDStr == "" {
		http.Error(w, "User ID parametr is required", http.StatusBadRequest)
		return
	}

	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	count, err := h.Service.CountBlogByUser(ctx, userID)
	if err != nil {
		http.Error(w, "Error fetching user articles count", http.StatusInternalServerError)
		return
	}

	jsonResponse := map[string]interface{}{
		"user_id ":       userID,
		"articles_count": count,
	}

	jsonData, err := json.Marshal(jsonResponse)
	if err != nil {
		http.Error(w, "Error marshalling JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}
