package router

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"github.com/knmsh08200/Blog_test/internal/db"
	"github.com/knmsh08200/Blog_test/internal/middleware"
	"github.com/knmsh08200/Blog_test/internal/model"
	"github.com/knmsh08200/Blog_test/internal/service"
)

func NewRouter() *mux.Router {
	mux := mux.NewRouter()

	mux.HandleFunc("/blog/list/", blogListHandler).Methods(http.MethodGet, http.MethodPost, http.MethodDelete)
	mux.HandleFunc("/blog/id/", blogIDHandler).Methods(http.MethodGet, http.MethodPost, http.MethodDelete)
	mux.HandleFunc("/blog/id/counter", blogCounterHandler).Methods(http.MethodGet)

	return mux
}

func NewHandler() http.Handler {
	mux := NewRouter()
	handler := middleware.MetricsMiddleware(mux)
	return handler
}

type List struct {
	ID      int    `json:"id"`
	USER_ID int    `json:"user_id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type ID struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func blogCounterHandler(w http.ResponseWriter, r *http.Request) {
	// Получаем параметры из URL-запроса
	params := r.URL.Query()
	userIDStr := params.Get("user_id")

	if userIDStr == "" {
		http.Error(w, "User ID parameter is required", http.StatusBadRequest)
		return
	}

	// Преобразуем userIDStr в целочисленный тип
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	// Запрос для подсчета количества статей пользователя
	var articlesCount int
	err = db.Db.QueryRow("SELECT COUNT(*) FROM lists WHERE user_id=$1", userID).Scan(&articlesCount)
	if err != nil {
		http.Error(w, "Error fetching user articles count", http.StatusInternalServerError)
		return
	}

	// Формируем JSON-ответ с количеством статей пользователя
	jsonResponse := map[string]interface{}{
		"user_id":        userID,
		"articles_count": articlesCount,
	}

	jsonData, err := json.Marshal(jsonResponse)
	if err != nil {
		http.Error(w, "Error marshalling JSON", http.StatusInternalServerError)
		return
	}

	// Отправляем JSON-ответ клиенту
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

func blogListHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST,DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	switch r.Method {
	case http.MethodGet:
		handleGetList(w)
	case http.MethodPost:
		handlePostList(w, r)
	case http.MethodDelete:
		handleDelList(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func handleGetList(w http.ResponseWriter) {
	log.Println("received request to get ")
	// w.Header().Set("Access-Control-Allow-Origin", "*")
	// w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	// w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	var lists []List
	rows, err := db.Db.Query("SELECT id,user_id,title,content FROM lists")
	if err != nil {
		log.Printf("Error querying database: %v", err)
		http.Error(w, "Error fetching data", http.StatusInternalServerError)
		return
	}

	defer rows.Close()
	for rows.Next() {
		var list List
		if err := rows.Scan(&list.ID, &list.USER_ID, &list.Title, &list.Content); err != nil {
			log.Printf("error: %v", err)
			http.Error(w, "Error scanning data", http.StatusInternalServerError)
			return
		}
		lists = append(lists, list)
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

func handlePostList(w http.ResponseWriter, r *http.Request) {

	log.Println("Received request to add new list")

	// w.Header().Set("Access-Control-Allow-Origin", "*")
	// w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	// w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}

	var newList List
	if err := json.Unmarshal(body, &newList); err != nil {
		http.Error(w, "Error unmarshalling JSON", http.StatusBadRequest)
		return
	}

	var userExists bool
	err = db.Db.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE id=$1)", newList.USER_ID).Scan(&userExists)
	if err != nil {
		http.Error(w, "Error checking user existence", http.StatusInternalServerError)
		return
	}

	if !userExists {
		http.Error(w, "User not found", http.StatusBadRequest)
		return
	}
	err = db.Db.QueryRow("INSERT INTO lists (user_id,title, content) VALUES ($1, $2, $3) RETURNING id", newList.USER_ID, newList.Title, newList.Content).Scan(&newList.ID)
	if err != nil {
		log.Printf("Error inserting data: %v", err)
		http.Error(w, "Error inserting data", http.StatusInternalServerError)
		return
	}
	// mutex.Lock()
	// newList.ID = len(lists) + 1
	// lists = append(lists, newList)
	// mutex.Unlock()

	w.WriteHeader(http.StatusCreated)
}

func handleDelList(w http.ResponseWriter, r *http.Request) {
	// получить blog id из request

	ids, err := service.BlogHandler.GetUsersByBlogID(ctx, id)

	w.Write()
}

func blogIDHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST,DELETE,OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	switch r.Method {
	case http.MethodGet:
		fmt.Println("get id req")
		handleGetID(w)
	case http.MethodPost:
		fmt.Println("post  req")
		handlePostID(w, r)
	case http.MethodDelete:
		fmt.Println("delete req")
		handleDelID(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func handleGetID(w http.ResponseWriter) {
	// w.Header().Set("Access-Control-Allow-Origin", "*")
	// w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	// w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	var ids []model.ID
	rows, err := db.Db.Query("SELECT id,name FROM users")
	if err != nil {
		log.Printf("Error querying database: %v", err)
		http.Error(w, "Error fetching data", http.StatusInternalServerError)
		return
	}

	for rows.Next() {
		var id model.ID
		if err := rows.Scan(&id.ID, &id.Name); err != nil {

			http.Error(w, "Error scanning data", http.StatusInternalServerError)
			return
		}
		ids = append(ids, id)
	}

	defer rows.Close()

	idResponse := model.ConvertDBtoResponse(ids)

	jsonData, err := json.Marshal(idResponse)
	if err != nil {
		http.Error(w, "Error marshalling JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

func handlePostID(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Access-Control-Allow-Origin", "*")
	// w.Header().Set("Access-Control-Allow-Methods", "GET, POST,DELETE,OPTIONS")
	// w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}

	var newID model.ID
	if err := json.Unmarshal(body, &newID); err != nil {
		http.Error(w, "Error unmarshalling JSON", http.StatusBadRequest)
		return
	}

	err = db.Db.QueryRow("INSERT INTO users (name) VALUES ($1) RETURNING id", newID.Name).Scan(&newID.ID)
	if err != nil {
		log.Printf("Error inserting data: %v", err)
		http.Error(w, "Error inserting data", http.StatusInternalServerError)
		return
	}

	// mutex.Lock()
	// newID.ID = len(ids) + 1
	// ids = append(ids, newID)
	// mutex.Unlock()

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Inserted ID: %d", newID.ID)
}

func handleDelID(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Access-Control-Allow-Origin", "*")
	// w.Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE, OPTIONS")
	// w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	fmt.Println("Вызов handleDelID") // Отладочное сообщение
	fmt.Println("URL.Path:", r.URL.Path)
	idStr := strings.TrimPrefix(r.URL.Path, "/blog/id/")
	id, err := strconv.Atoi(idStr)
	fmt.Println(id)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	result, err := db.Db.Exec("DELETE FROM users WHERE id = $1", id)
	if err != nil {
		http.Error(w, "Error deleting data", http.StatusInternalServerError)
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("error :%v", err)
		http.Error(w, "Error getting rows affected", http.StatusInternalServerError)
		return
	}

	fmt.Println(strconv.Itoa(int(rowsAffected)))
	// mutex.Lock()
	// defer mutex.Unlock()

	// for i, ident := range ids {
	// 	if ident.ID == id {
	// 		ids = append(ids[:i], ids[i+1:]...)
	// 		w.WriteHeader(http.StatusOK)
	// 		return
	// 	}
	// }

	// http.Error(w, "ID not found", http.StatusNotFound) ИНТЕРЕССССС

}

// Middleware part
