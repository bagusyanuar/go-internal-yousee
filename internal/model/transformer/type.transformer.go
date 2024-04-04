package transformer

import (
	"fmt"

	"github.com/bagusyanuar/go-internal-yousee/internal/entity"
	"github.com/bagusyanuar/go-internal-yousee/internal/model"
)

func ToType(itemType *entity.Type) *model.TypeResponse {
	var icon *string
	if itemType.Icon != nil {
		iconPath := fmt.Sprintf("http://127.0.0.1:8000/%s", *itemType.Icon)
		icon = &iconPath
	}
	return &model.TypeResponse{
		ID:   itemType.ID,
		Name: itemType.Name,
		Icon: icon,
	}
}

func ToTypes(entities []entity.Type) []model.TypeResponse {
	types := make([]model.TypeResponse, 0)
	for _, entity := range entities {
		t := *ToType(&entity)
		types = append(types, t)
	}
	return types
}
