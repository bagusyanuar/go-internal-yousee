package transformer

import (
	"github.com/bagusyanuar/go-internal-yousee/internal/entity"
	"github.com/bagusyanuar/go-internal-yousee/internal/model"
)

func ToType(itemType *entity.Type) *model.TypeResponse {
	return &model.TypeResponse{
		ID:   itemType.ID,
		Name: itemType.Name,
		Icon: itemType.Icon,
	}
}

func ToTypes(entities []entity.Type) []model.TypeResponse {
	var types []model.TypeResponse
	for _, entity := range entities {
		t := *ToType(&entity)
		types = append(types, t)
	}
	return types
}
