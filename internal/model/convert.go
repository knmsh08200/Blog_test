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

			Title: list.Title,
		}
	}

	return responses
}

func ConvertFindListtoResponse(article FindList) FindListResponse {

	return FindListResponse{
		Id:      article.ID,
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
