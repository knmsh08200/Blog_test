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

func ConvertDBtoResponse(ids []ID) []IDResponse {
	var responses []IDResponse

	for _, id := range ids {
		responses = append(responses, IDResponse{Name: id.Name})
	}
	return responses
}
