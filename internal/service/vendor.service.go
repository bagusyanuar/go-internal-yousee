package service

import (
	"context"

	"github.com/bagusyanuar/go-internal-yousee/internal/model"
	"github.com/bagusyanuar/go-internal-yousee/internal/model/transformer"
	"github.com/bagusyanuar/go-internal-yousee/internal/repositories"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
)

type (
	VendorService interface {
		FindAll(ctx context.Context, queryString model.QueryString[string]) (model.Response[[]model.VendorResponse], error)
	}

	vendor struct {
		VendorRepository repositories.VendorRepository
		Log              *logrus.Logger
		Validator        *validator.Validate
	}
)

// FindAll implements VendorService.
func (service *vendor) FindAll(ctx context.Context, queryString model.QueryString[string]) (model.Response[[]model.VendorResponse], error) {
	var vendors []model.VendorResponse
	response, err := service.VendorRepository.FindAll(ctx, queryString)
	if err != nil {
		return model.Response[[]model.VendorResponse]{}, err
	}
	vendors = transformer.ToVendors(response.Data)
	return model.Response[[]model.VendorResponse]{Data: vendors, Meta: response.Meta}, nil
}

func NewVendorService(
	vendorRepository repositories.VendorRepository,
	log *logrus.Logger,
	validator *validator.Validate,
) VendorService {
	return &vendor{
		VendorRepository: vendorRepository,
		Log:              log,
		Validator:        validator,
	}
}
