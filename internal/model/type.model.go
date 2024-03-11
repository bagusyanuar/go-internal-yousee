package model

type TypeRequest struct {
	Name string  `json:"name"`
	Icon *string `json:"icon"`
}

type TypeResponse struct {
	ID   uint64  `json:"id"`
	Name string  `json:"name"`
	Icon *string `json:"icon"`
}
