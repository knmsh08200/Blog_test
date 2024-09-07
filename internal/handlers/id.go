package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/knmsh08200/Blog_test/internal/blog"
	"github.com/knmsh08200/Blog_test/internal/model"
)

// BlogHandler defines the handler structure.

type IDHandler struct {
	IDProvider blog.IDRepository
}

func (h *IDHandler) BlogIDHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	switch r.Method {
	case http.MethodGet:
		h.handleGetID(r.Context(), w)
	case http.MethodPost:
		h.handlePostID(r.Context(), w, r)
	case http.MethodDelete:
		h.handleDelID(r.Context(), w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *IDHandler) handleGetID(ctx context.Context, w http.ResponseWriter) {
	ids, err := h.IDProvider.GetAllId(ctx)
	if err != nil {
		log.Printf("Error fetching blogs: %v", err)
		http.Error(w, "Error fetching data", http.StatusInternalServerError)
		return
	}

	jsonData, err := json.Marshal(ids)
	if err != nil {
		http.Error(w, "Error marshalling JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

func (h *IDHandler) handlePostID(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}
	fmt.Println("Before unmarshall", string(body))

	var newID model.CreateID
	if err := json.Unmarshal(body, &newID); err != nil {
		http.Error(w, "Error unmarshalling JSON", http.StatusBadRequest)
		return
	}

	fmt.Println("Result unmurshalling", newID.ID, newID.Name)

	id, err := h.IDProvider.CreateID(ctx, newID)
	if err != nil && id == -1 {
		http.Error(w, "Already exist", http.StatusInternalServerError)
	}

	newID.ID = id

	w.WriteHeader(http.StatusCreated)
}

func (h *IDHandler) handleDelID(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/blog/id/")
	id, err := strconv.Atoi(idStr)

	fmt.Println(id)

	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	rowsAffected, err := h.IDProvider.DeleteID(ctx, id)
	if err != nil {
		http.Error(w, "Error deleting data", http.StatusInternalServerError)
		return
	}

	if rowsAffected == 0 {
		http.Error(w, "Blog not found", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// NewBlogHandler returns a new instance of BlogHandler.
func NewBlogHandler(s blog.ListRepository, t blog.IDRepository) (*BlogHandler, *IDHandler) {
	return &BlogHandler{Service: s}, &IDHandler{IDProvider: t}
}

// BlogListHandler handles the /blog/list endpoint.
