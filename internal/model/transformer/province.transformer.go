package transformer

import (
	"github.com/bagusyanuar/go-internal-yousee/internal/entity"
	"github.com/bagusyanuar/go-internal-yousee/internal/model"
)

func ProvinceToResponse(province *entity.Province) *model.ProvinceResponse {
	return &model.ProvinceResponse{
		ID:   province.ID,
		Name: province.Name,
	}
}

func ToProvinces(entities []entity.Province) []model.ProvinceResponse {
	var provinces []model.ProvinceResponse
	for _, entity := range entities {
		t := model.ProvinceResponse{
			ID:   entity.ID,
			Name: entity.Name,
		}
		provinces = append(provinces, t)
	}
	return provinces
}
