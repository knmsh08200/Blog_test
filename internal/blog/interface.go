package blog

import (
	"context"

	"github.com/knmsh08200/Blog_test/internal/model"
)

type Repository interface {
	GetAllBlogs(ctx context.Context) ([]model.List, error)
	CreateBlog(ctx context.Context, list model.List) (int, error)
	DeleteBlog(ctx context.Context, d int) (int64, error) // работа с бд
	CounterUserBlog(ctx context.Context, userID int) (int, error)
}
