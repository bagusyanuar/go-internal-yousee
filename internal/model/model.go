package model

type APIResponse[T any] struct {
	Code    int    `json:"code"`
	Data    T      `json:"data"`
	Message string `json:"message"`
}

type PaginationQuery struct {
	Page    int `json:"page"`
	PerPage int `json:"per_page"`
}
