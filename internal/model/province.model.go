package model

type ProvinceResponse struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
}

type PaginationProvincesResponse struct {
	Data       []ProvinceResponse
	Pagination *MetaPagination
}
