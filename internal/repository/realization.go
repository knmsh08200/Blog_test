package repository

import (
	"database/sql"

	"github.com/knmsh08200/Blog_test/internal/model"
)

type blogRepository struct {
	DB *sql.DB
}

func NewPostRep(db *sql.DB) BlogRepository {
	return &blogRepository{DB: db}
}

func (r *blogRepository) GetAllBlogs() ([]model.List, error) {
	rows, err := r.DB.Query("SELECT id,user_id,title,content FROM lists")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var lists []model.List //не знаю как без var, как я понимаю, тут нельзя применить make
	for rows.Next() {
		var list model.List
		if err := rows.Scan(&list.ID, &list.UserID, &list.Title, &list.Content); err != nil {
			return nil, err
		}
		lists = append(lists, list)
	}
	return lists, err
}

func (r *blogRepository) CreateBlog(list model.List) (int, error) {

	id := 0
	err := r.DB.QueryRow("INSERT INTO lists (user_id,title, content) VALUES ($1, $2, $3) RETURNING id", list.UserID, list.Title, list.Content).Scan(&list.ID)
	return id, err
}

func (r *blogRepository) DeleteBlog(id int) (int64, error) {
	result, err := r.DB.Exec("DELETE FROM lists WHERE id = $1", id)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}
func (r *blogRepository) CounterUserBlog(userID int) (int, error) {
	articlesCount := 0
	err := r.DB.QueryRow("SELECT COUNT(*) FROM lists WHERE user_id=$1", userID).Scan(&articlesCount)
	if err != nil {
		return 0, err
	}
	return articlesCount, err
}

// // не знаю нужно или нет, думаю надо, так как ...
// type userDB struct {
// 	DB *sql.DB
// }
