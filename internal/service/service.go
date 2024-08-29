package service

import (
	"context"

	"github.com/knmsh08200/Blog_test/internal/blog"
	"github.com/knmsh08200/Blog_test/internal/model"
)

type BlogService struct {
	Repo blog.Repository
}

func NewBlogService(repo blog.Repository) *BlogService {
	return &BlogService{Repo: repo}
}

func (s *BlogService) GetAllBlogs(ctx context.Context) ([]model.List, error) {
	return s.Repo.GetAllBlogs(ctx)
}

func (s *BlogService) CreateBlog(ctx context.Context, list model.List) (int, error) {
	return s.Repo.CreateBlog(ctx, list)
}
func (s *BlogService) DeleteBlog(ctx context.Context, id int) (int64, error) {
	return s.Repo.DeleteBlog(ctx, id)
}

func (s *BlogService) CountBlogByUser(ctx context.Context, userID int) (int, error) {
	return s.Repo.CounterUserBlog(ctx, userID)
}
