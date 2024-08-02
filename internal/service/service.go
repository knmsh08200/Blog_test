package service

import "github.com/knmsh08200/Blog_test/internal/repository"

type BlogService struct {
	Repo repository.BlogRepository
}

func NewBlogService(repo repository.BlogRepository) *BlogService {
	return &BlogService{Repo: repo}
}
