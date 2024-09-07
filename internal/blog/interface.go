package blog

import (
	"context"

	"github.com/knmsh08200/Blog_test/internal/model"
)

type ListRepository interface {
	GetAllBlogs(ctx context.Context, limit, offset int) ([]model.ListResponse, model.Meta, error)
	CreateBlog(ctx context.Context, list model.List) (int, error)
	DeleteBlog(ctx context.Context, d int) (int64, error) // работа с бд
	FindBlog(ctx context.Context, title string) (model.FindList, error)
	CounterUserBlog(userID int) (int, error)
}

type IDRepository interface {
	GetAllId(ctx context.Context) ([]model.ID, error)
	CreateID(ctx context.Context, list model.CreateID) (int, error)
	DeleteID(ctx context.Context, d int) (int64, error) // работа с бд
}
