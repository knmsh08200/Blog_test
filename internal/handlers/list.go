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

type BlogHandler struct {
	Service blog.ListRepository
}

func (h *BlogHandler) BlogFindListHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	switch r.Method {
	case http.MethodGet:
		h.handleFindList(r.Context(), w, r)
	}
}

func (h *BlogHandler) BlogListHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	switch r.Method {
	case http.MethodGet:
		h.handleGetList(r.Context(), w, r)
	case http.MethodPost:
		h.handlePostList(r.Context(), w, r)
	case http.MethodDelete:
		h.handleDelList(r.Context(), w, r)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *BlogHandler) handleFindList(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	idQuery := strings.TrimPrefix(r.URL.Path, "/blog/list/")

	id, err := strconv.Atoi(idQuery)
	if err != nil {
		http.Error(w, "Error fetching data", http.StatusBadRequest)
		log.Printf("Error fetching ID of blog:%v", err)
	}

	fmt.Println(id)

	if id == 0 {
		http.Error(w, "Title query parameter is required", http.StatusBadRequest)
		return
	}

	article, err := h.Service.FindBlog(ctx, id)

	if err != nil {
		log.Printf("Error fetching blogs: %v", err)
		http.Error(w, "Error fetching data", http.StatusInternalServerError)
		return
	}

	arcticleResp := model.ConvertFindListtoResponse(article)
	jsonData, err := json.Marshal(arcticleResp)

	if err != nil {
		http.Error(w, "Error marshalling JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)

}

func (h *BlogHandler) handleGetList(ctx context.Context, w http.ResponseWriter, r *http.Request) {

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}
	fmt.Println("Before unmarshall", string(body))

	var newMeta model.MetaRequest
	if err := json.Unmarshal(body, &newMeta); err != nil {
		http.Error(w, "Error unmarshalling JSON", http.StatusBadRequest)
		log.Printf("Error:%v", err)
		return
	}
	// TODO убрать в get all blogs
	offset := (newMeta.Page - 1) * newMeta.Limit

	lists, meta, err := h.Service.GetAllBlogs(ctx, newMeta.Limit, offset)

	if err != nil {
		log.Printf("Error fetching blogs: %v", err)
		http.Error(w, "Error fetching data", http.StatusBadRequest)
		return
	}
	// todo убрать чтобы не возвращало null в реквест --- написать Матвею когда разберусь
	BlogResponse := model.ConvertBlogListResponse(meta, lists)
	jsonData, err := json.Marshal(BlogResponse)

	if err != nil {
		http.Error(w, "Error marshalling JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

func (h *BlogHandler) handlePostList(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}
	fmt.Println("Before unmarshall", string(body))

	var newList model.List
	if err := json.Unmarshal(body, &newList); err != nil {
		http.Error(w, "Error unmarshalling JSON", http.StatusBadRequest)
		return
	}

	fmt.Println("Result unmurshalling", newList.Content)

	id, err := h.Service.CreateBlog(ctx, newList)
	if err != nil {
		http.Error(w, "Error inserting data", http.StatusInternalServerError)
		return
	}
	newList.ID = id

	w.WriteHeader(http.StatusCreated)
}

func (h *BlogHandler) handleDelList(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/blog/list/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	rowsAffected, err := h.Service.DeleteBlog(ctx, id)
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

func (h *BlogHandler) BlogCounterHandler(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()

	userIDStr := params.Get("user_id")

	fmt.Println(userIDStr)

	if userIDStr == "" {
		http.Error(w, "User ID parameter is required", http.StatusBadRequest)
		return
	}

	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	count, err := h.Service.CounterUserBlog(userID)
	if err != nil {
		http.Error(w, "Error fetching user articles count", http.StatusInternalServerError)
		return
	}

	jsonResponse := map[string]interface{}{
		"user_id":        userID,
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
