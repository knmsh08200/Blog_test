package service

import (
	"github.com/knmsh08200/Blog_test/internal/model"
)

// ConvertUserToResponse конвертирует структуру User из базы данных в UserResponse для ответа
func ConvertIDToResponse(user model.ID) model.IDResponse {
	return model.IDResponse{
		ID:   user.ID,
		Name: user.Name,
	}
}
