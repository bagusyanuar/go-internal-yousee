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
