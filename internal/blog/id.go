package blog

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/knmsh08200/Blog_test/internal/model"
)

func (r *blogRepository) GetAllId(ctx context.Context) ([]model.ID, error) {

	rows, err := r.db.Query("SELECT id,name FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ids []model.ID //не знаю как без var, как я понимаю, тут нельзя применить make
	for rows.Next() {
		var id model.ID
		if err := rows.Scan(&id.ID, &id.Name); err != nil {
			return nil, err
		}
		ids = append(ids, id)
	}
	return ids, err
}

func (r *blogRepository) CreateID(ctx context.Context, id model.CreateID) (int, error) {
	existingUserID := 0

	err := r.db.QueryRow("SELECT id FROM users WHERE name = $1", id.Name).Scan(&existingUserID)
	if err != nil && err != sql.ErrNoRows {
		return 0, fmt.Errorf("ошибка при проверке существующего пользователя: %w", err)
	}

	if existingUserID > 0 {
		return -1, fmt.Errorf("пользователь с именем '%s' уже существует", id.Name)

	}

	iden := 0
	error := r.db.QueryRow("INSERT INTO  users  (id,name) VALUES ($1, $2) RETURNING id", id.ID, id.Name).Scan(&id.ID)
	return iden, error // вопрос что здесь будет в этой переменной, сделал по пример list
}

func (r *blogRepository) DeleteID(ctx context.Context, id int) (int64, error) {
	result, err := r.db.Exec("DELETE FROM users WHERE id = $1", id)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}
