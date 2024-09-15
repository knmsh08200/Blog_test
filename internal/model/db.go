package model

// здесь опишу переменные связанные с бд
type List struct {
	ID      int
	UserID  int
	Title   string
	Content string
}

type ID struct {
	ID   int
	Name string
}

type FindList struct {
	ID      int
	Content string
}
