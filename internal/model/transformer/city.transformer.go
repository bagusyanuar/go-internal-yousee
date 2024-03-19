package transformer

import (
	"github.com/bagusyanuar/go-internal-yousee/internal/entity"
	"github.com/bagusyanuar/go-internal-yousee/internal/model"
)

func ToCity(city *entity.City) *model.CityResponse {
	var province *model.CityProvince
	if city.Province != nil {
		province = &model.CityProvince{
			ID:   city.Province.ID,
			Name: city.Province.Name,
		}
	}

	return &model.CityResponse{
		ID:         city.ID,
		Name:       city.Name,
		ProvinceID: city.ProvinceID,
		Province:   province,
	}
}

func ToCities(entities []entity.City) []model.CityResponse {
	cities := make([]model.CityResponse, 0)
	for _, entity := range entities {
		t := *ToCity(&entity)
		cities = append(cities, t)
	}
	return cities
}
