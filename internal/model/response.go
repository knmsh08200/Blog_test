package model

// здесь опишу переменные связанные с  ответом  на вопрос

type ListResponse struct {
	ID      int    `json:"id"`
	UserID  int    `json:"user_id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type IDResponse struct {
	Name string `json:"name"`
}

type IDsResponse struct {
	IDs []IDResponse `json:"ids"`
}

func ConvertDBtoResponse(ids []ID) *IDsResponse {

	responses := make([]IDResponse, len(ids))

	for i, id := range ids {
		responses[i] = IDResponse{Name: id.Name}
	}
	return &IDsResponse{IDs: responses}
}
