package transformer

import (
	"github.com/bagusyanuar/go-internal-yousee/internal/entity"
	"github.com/bagusyanuar/go-internal-yousee/internal/model"
)

func ToItem(item *entity.Item) *model.ItemResponse {

	city := toItemCity(item.City)
	itemType := toItemType(item.Type)
	vendor := toItemVendor(item.Vendor)
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
		City:      city,
		Type:      itemType,
		Vendor:    vendor,
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

func toItemCity(city *entity.City) *model.ItemCity {
	if city != nil {
		return &model.ItemCity{
			ID:   city.ID,
			Name: city.Name,
		}
	}
	return nil
}

func toItemType(itemType *entity.Type) *model.ItemType {
	if itemType != nil {
		return &model.ItemType{
			ID:   itemType.ID,
			Name: itemType.Name,
		}
	}
	return nil
}

func toItemVendor(vendor *entity.Vendor) *model.ItemVendor {
	if vendor != nil {
		return &model.ItemVendor{
			ID:       vendor.ID,
			Email:    vendor.Email,
			Name:     vendor.Name,
			Address:  vendor.Address,
			Phone:    vendor.Phone,
			Brand:    vendor.Brand,
			PICName:  vendor.PICName,
			PICPhone: vendor.PICPhone,
			LastSeen: vendor.LastSeen,
		}
	}
	return nil
}
