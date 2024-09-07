package model

// сделано для удобства фронт разработки
func ConvertDBtoResponse(ids []ID) *IDsResponse {

	responses := make([]IDResponse, len(ids))

	for i, id := range ids {
		responses[i] = IDResponse{Name: id.Name}
	}
	return &IDsResponse{IDs: responses}
}

func ConvertListtoResponse(lists []List) []ListResponse {
	responses := make([]ListResponse, len(lists))

	for i, list := range lists {
		responses[i] = ListResponse{

			Title:   list.Title,
			Content: list.Content,
		}
	}

	return responses
}

func ConvertFindListtoResponse(article FindList) FindListResponse {

	return FindListResponse{
		Name:    article.Name,
		Title:   article.Title,
		Content: article.Content,
	}

}

func ConvertBlogListResponse(meta Meta, lists []ListResponse) BlogResponse {
	BlogResponse := BlogResponse{
		Meta:  meta,
		Blogs: lists,
	}
	return BlogResponse

}
