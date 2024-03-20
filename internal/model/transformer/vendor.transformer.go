package transformer

import (
	"github.com/bagusyanuar/go-internal-yousee/internal/entity"
	"github.com/bagusyanuar/go-internal-yousee/internal/model"
)

func ToVendor(vendor *entity.Vendor) *model.VendorResponse {
	return &model.VendorResponse{
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

func ToVendors(entities []entity.Vendor) []model.VendorResponse {
	vendors := make([]model.VendorResponse, 0)
	for _, entity := range entities {
		t := *ToVendor(&entity)
		vendors = append(vendors, t)
	}
	return vendors
}
