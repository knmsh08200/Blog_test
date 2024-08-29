package model

// здесь опишу переменные связанные с бд
type List struct {
	ID      int
	UserID  int
	Title   string
	Content string
}

type ID struct {
	Name string
	ID   int
}
