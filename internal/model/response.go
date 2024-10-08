package model

// здесь опишу переменные связанные с  ответом  на вопрос

type ListResponse struct {
	Title   string `json:"title,omitempty"`
	Content string `json:"content,omitempty"`
}

type FindListResponse struct {
	Name    string `json:"name"`
	Title   string `json:"title,omitempty"`
	Content string `json:"content,omitempty"`
}

type IDResponse struct {
	Name string `json:"name"`
}

type IDsResponse struct {
	IDs []IDResponse `json:"ids"`
}

type Meta struct {
	Limit int `json:"limit"`
	Page  int `json:"page"`
	Total int `json:"total"`
}

type BlogResponse struct {
	Meta  Meta           `json:"meta"`
	Blogs []ListResponse `json:"blogs"`
}
