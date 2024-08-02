package model

// здесь опишу переменные связанные с получением запроса
type СreateListRequest struct {
	ID      int    `json:"id"`
	UserID  int    `json:"user_id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type СreateIDRequest struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
