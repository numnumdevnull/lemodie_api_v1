package model

type DictionaryTypes struct {
	ID    uint64  `json:"id"`
	Value string  `json:"value"`
	Meta  *string `json:"meta"`
}

type PaginatedResponse[T any] struct {
	Data       []T `json:"data"`
	Total      int `json:"total"`
	Page       int `json:"page"`
	Limit      int `json:"limit"`
	TotalPages int `json:"total_pages"`
}
