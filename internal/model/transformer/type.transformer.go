package transformer

import (
	"github.com/bagusyanuar/go-internal-yousee/internal/entity"
	"github.com/bagusyanuar/go-internal-yousee/internal/model"
)

func TypeToResponse(itemType *entity.Type) *model.TypeResponse {
	return &model.TypeResponse{
		ID:   itemType.ID,
		Name: itemType.Name,
		Icon: itemType.Icon,
	}
}
