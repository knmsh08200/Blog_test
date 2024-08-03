package service

type BlogHandler interface {
	GetUsersByBlogID(blogID int) ([]int, error)
}
