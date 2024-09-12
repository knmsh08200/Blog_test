package blog

import (
	"context"
	"database/sql"

	"fmt"

	"github.com/knmsh08200/Blog_test/internal/model"
	"github.com/pkg/errors"
)

type blogRepository struct {
	db *sql.DB
}

//	func NewRep(rep Repository) (*blogRepository, error) {
//		if br, ok := rep.(*blogRepository); ok {
//			return br, nil
//		}
//		return nil, errors.New("provided rep failed")
//	}
func NewRep(db *sql.DB) *blogRepository {
	return &blogRepository{db: db}
}

func (r *blogRepository) GetAllBlogs(ctx context.Context, limit, offset int) ([]model.ListResponse, model.Meta, error) {

	var total int
	err := r.db.QueryRowContext(ctx, "SELECT COUNT(*) FROM lists").Scan(&total)
	if err != nil {

		return nil, model.Meta{}, errors.Wrap(err, "GetAllBlogs  QueryRowContext")
	}

	query := "SELECT  title FROM lists LIMIT $1 OFFSET $2" // +select id

	rows, err := r.db.QueryContext(ctx, query, limit, offset)
	if err != nil {
		return nil, model.Meta{}, err
	}
	defer rows.Close()

	lists := []model.ListResponse{}
	for rows.Next() {
		var list model.ListResponse
		if err := rows.Scan(&list.Title); err != nil {
			return nil, model.Meta{}, err
		}
		lists = append(lists, list)
	}

	meta := model.Meta{
		Limit: limit,
		Page:  offset/limit + 1,
		Total: total,
	}

	return lists, meta, err
}

func (r *blogRepository) CreateBlog(ctx context.Context, list model.List) (int, error) {

	id := 0
	err := r.db.QueryRowContext(ctx, "INSERT INTO lists (user_id, title, content) VALUES ($1, $2, $3) RETURNING id", list.UserID, list.Title, list.Content).Scan(&id)
	return id, err
}

func (r *blogRepository) DeleteBlog(ctx context.Context, id int) (int64, error) {
	result, err := r.db.Exec("DELETE FROM lists WHERE id = $1", id)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}
func (r *blogRepository) CounterUserBlog(userID int) (int, error) {
	articlesCount := 0
	err := r.db.QueryRow("SELECT COUNT(*) FROM lists WHERE user_id=$1", userID).Scan(&articlesCount)
	if err != nil {
		return 0, err
	}
	return articlesCount, err
}

func (r *blogRepository) FindBlog(ctx context.Context, id int) (model.FindList, error) {
	query := `
        SELECT id,content FROM lists WHERE id = $1
    `
	var list model.FindList
	err := r.db.QueryRowContext(ctx, query, id).Scan(&list.ID, &list.Content)
	if err != nil {
		if err == sql.ErrNoRows {
			return list, fmt.Errorf("article not found")
		}
		return list, err
	}

	return list, nil
}

// // не знаю нужно или нет, думаю надо, так как ...
// type userDB struct {
// 	DB *sql.DB
// }
