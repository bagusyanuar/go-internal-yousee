package transformer

import (
	"github.com/bagusyanuar/go-internal-yousee/internal/entity"
	"github.com/bagusyanuar/go-internal-yousee/internal/model"
)

func ToItem(item *entity.Item) *model.ItemResponse {

	return &model.ItemResponse{
		ID:        item.ID,
		TypeID:    item.TypeID,
		CityID:    item.CityID,
		VendorID:  item.VendorID,
		Name:      item.Name,
		Address:   item.Address,
		Latitude:  item.Latitude,
		Longitude: item.Longitude,
		Location:  item.Location,
		URL:       item.URL,
		Width:     item.Width,
		Height:    item.Height,
		Position:  item.Position,
		CreatedAt: item.CreatedAt,
		UpdatedAt: item.UpdatedAt,
	}
}

func ToItems(entities []entity.Item) []model.ItemResponse {
	var items []model.ItemResponse
	for _, entity := range entities {
		t := *ToItem(&entity)
		items = append(items, t)
	}
	return items
}
