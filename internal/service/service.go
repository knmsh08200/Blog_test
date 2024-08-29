package service

import (
	"context"

	"github.com/knmsh08200/Blog_test/internal/blog"
	"github.com/knmsh08200/Blog_test/internal/model"
)

type BlogService struct {
<<<<<<< HEAD
	Repo blog.Repository
=======
	blogProvider repository.BlogRepository
	userProvider repository.UserRepository
>>>>>>> 13f7ffce47aec85e97a6946bac9796e6625f8111
}

func NewBlogService(repo blog.Repository) *BlogService {
	return &BlogService{Repo: repo}
}

<<<<<<< HEAD
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
=======
// пример
func (b *BlogService) GetUsersByBlogID(blogID int) ([]int, error) {
	b.userProvider.GetAllUsers()
	b.blogProvider.CounterUserBlog()

	return id, nil
>>>>>>> 13f7ffce47aec85e97a6946bac9796e6625f8111
}
