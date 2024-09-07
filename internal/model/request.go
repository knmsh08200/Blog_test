package model

const defaultlimit = 5

// здесь опишу переменные связанные с получением запроса
type CreateList struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type CreateID struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type MetaRequest struct {
	Limit int `json:"limit"`
	Page  int `json:"page"`
}

// TODO make config file where --- config.yaml !!!! ++ vyper
func (m *MetaRequest) Validation() error {
	if m.Limit > defaultlimit {
		m.Limit = defaultlimit
	}
	return nil
}
