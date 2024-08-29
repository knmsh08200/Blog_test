package blog

import (
	"context"
	"database/sql"

	"github.com/knmsh08200/Blog_test/internal/model"
)

type blogRepository struct {
	db *sql.DB
}

func NewRep(db *sql.DB) Repository {
	return &blogRepository{db: db}
}

func (r *blogRepository) GetAllBlogs(ctx context.Context) ([]model.List, error) {
	rows, err := r.db.Query("SELECT id,user_id,title,content FROM lists")
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

func (r *blogRepository) CreateBlog(ctx context.Context, list model.List) (int, error) {

	id := 0
	err := r.db.QueryRow("INSERT INTO lists (user_id,title, content) VALUES ($1, $2, $3) RETURNING id", list.UserID, list.Title, list.Content).Scan(&list.ID)
	return id, err
}

func (r *blogRepository) DeleteBlog(ctx context.Context, id int) (int64, error) {
	result, err := r.db.Exec("DELETE FROM lists WHERE id = $1", id)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}
func (r *blogRepository) CounterUserBlog(ctx context.Context, userID int) (int, error) {
	articlesCount := 0
	err := r.db.QueryRow("SELECT COUNT(*) FROM lists WHERE user_id=$1", userID).Scan(&articlesCount)
	if err != nil {
		return 0, err
	}
	return articlesCount, err
}

// // не знаю нужно или нет, думаю надо, так как ...
// type userDB struct {
// 	DB *sql.DB
// }
