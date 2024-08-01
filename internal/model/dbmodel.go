package model

// здесь опишу переменные связанные с бд
type List struct {
	ID      int    `json:"id"`
	USER_ID int    `json:"user_id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type ID struct {
	Name string `json:"name"`
	ID   int    `json:"id"`
}
