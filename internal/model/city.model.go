package model

type CityResponse struct {
	ID         uint64        `json:"id"`
	ProvinceID uint64        `json:"province_id"`
	Name       string        `json:"name"`
	Province   *CityProvince `json:"province"`
}

type CityProvince struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
}
