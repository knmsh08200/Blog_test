package model

func ConvertDBtoResponse(ids []ID) []IDResponse {
	var responses []IDResponse

	for _, id := range ids {
		responses = append(responses, IDResponse{Name: id.Name})
	}
	return responses
}
