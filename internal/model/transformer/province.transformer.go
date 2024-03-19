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

func ToProvince(province *entity.Province) *model.ProvinceResponse {
	return &model.ProvinceResponse{
		ID:   province.ID,
		Name: province.Name,
	}
}

func ToProvinces(entities []entity.Province) []model.ProvinceResponse {
	provinces := make([]model.ProvinceResponse, 0)
	for _, entity := range entities {
		t := *ToProvince(&entity)
		provinces = append(provinces, t)
	}
	return provinces
}
