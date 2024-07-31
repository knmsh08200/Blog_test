package model

// здесь опишу переменные связанные с  ответом  на вопрос

type ListResponse struct {
	ID      int    `json:"id"`
	USER_ID int    `json:"user_id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type IDResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
