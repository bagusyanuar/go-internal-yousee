package service

import (
	"context"
	"time"

	"github.com/bagusyanuar/go-internal-yousee/common"
	"github.com/bagusyanuar/go-internal-yousee/internal/entity"
	"github.com/bagusyanuar/go-internal-yousee/internal/model"
	"github.com/bagusyanuar/go-internal-yousee/internal/model/transformer"
	"github.com/bagusyanuar/go-internal-yousee/internal/repositories"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
)

type (
	VendorService interface {
		FindAll(ctx context.Context, queryString model.QueryString[string]) (model.Response[[]model.VendorResponse], error)
		FindByID(ctx context.Context, id string) (*model.VendorResponse, error)
		Create(ctx context.Context, request *model.VendorRequest) (any, error)
		Patch(ctx context.Context, id string, request *model.VendorRequest) (any, error)
		Delete(ctx context.Context, id string) error
	}

	vendor struct {
		VendorRepository repositories.VendorRepository
		Log              *logrus.Logger
		Validator        *validator.Validate
	}
)

// Create implements VendorService.
func (service *vendor) Create(ctx context.Context, request *model.VendorRequest) (any, error) {

	//validate form request
	errValidation, msg := common.Validate(service.Validator, request)
	if errValidation != nil {
		return msg, common.ErrBadRequest
	}

	email := request.Email
	name := request.Name
	address := request.Address
	phone := request.Phone
	brand := request.Brand
	picName := request.PICName
	picPhone := request.PICPhone
	lastSeen := time.Now()

	data := &entity.Vendor{
		Name:     name,
		Email:    email,
		Address:  address,
		Phone:    phone,
		Brand:    brand,
		PICName:  picName,
		PICPhone: picPhone,
		LastSeen: &lastSeen,
	}

	err := service.VendorRepository.Create(ctx, data)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

// Patch implements VendorService.
func (service *vendor) Patch(ctx context.Context, id string, request *model.VendorRequest) (any, error) {
	//validate form request
	errValidation, msg := common.Validate(service.Validator, request)
	if errValidation != nil {
		return msg, common.ErrBadRequest
	}

	email := request.Email
	name := request.Name
	address := request.Address
	phone := request.Phone
	brand := request.Brand
	picName := request.PICName
	picPhone := request.PICPhone

	data := map[string]interface{}{
		"name":     name,
		"phone":    phone,
		"brand":    brand,
		"email":    email,
		"picName":  picName,
		"picPhone": picPhone,
		"address":  address,
	}

	err := service.VendorRepository.Patch(ctx, id, data)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

// Delete implements VendorService.
func (service *vendor) Delete(ctx context.Context, id string) error {
	panic("unimplemented")
}

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

// FindByID implements VendorService.
func (service *vendor) FindByID(ctx context.Context, id string) (*model.VendorResponse, error) {
	entity, err := service.VendorRepository.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return transformer.ToVendor(entity), nil
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
