package model

type APIResponse[T any] struct {
	Code    int    `json:"code"`
	Data    T      `json:"data"`
	Message string `json:"message"`
}

type Response[T any] struct {
	Data T
	Meta *MetaPagination
}

type InterfaceResponse[T any] struct {
	Status         int
	Data           T
	MetaPagination *MetaPagination
	Validation     any
	Error          error
}

type QueryString[T any] struct {
	Query           T
	QueryPagination PaginationQuery
}

type PaginationQuery struct {
	Page    int `json:"page"`
	PerPage int `json:"per_page"`
}

type MetaPagination struct {
	Page      int   `json:"page"`
	PerPage   int   `json:"per_page"`
	TotalPage int   `json:"total_page"`
	TotalRows int64 `json:"total_rows"`
}
