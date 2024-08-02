package service

import "github.com/knmsh08200/Blog_test/internal/repository"

type BlogService struct {
	blogProvider repository.BlogRepository
	userProvider repository.UserRepository
}

func NewBlogService(repo repository.BlogRepository) *BlogService {
	return &BlogService{Repo: repo}
}

// пример
func (b *BlogService) GetUsersByBlogID(blogID int) ([]int, error) {
	b.userProvider.GetAllUsers()
	b.blogProvider.CounterUserBlog()

	return id, nil
}
