package repository

import "github.com/knmsh08200/Blog_test/internal/model"

type BlogRepository interface {
	GetAllBlogs() ([]model.List, error)
	CreateBlog(list model.List) (int, error)
	DeleteBlog(id int) (int64, error) // работа с бд
	CounterUserBlog(userID int) (int, error)
}

type UserRepository interface {
	GetAllUsers() ([]model.ID, error)
	CreateUser(user model.ID) (int, error) //работа с бд
	DeleteUser(id int) (int64, error)
}
