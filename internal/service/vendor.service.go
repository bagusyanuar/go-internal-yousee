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
		FindAll(ctx context.Context, queryString model.QueryString[string]) model.InterfaceResponse[[]model.VendorResponse]
		FindByID(ctx context.Context, id string) model.InterfaceResponse[*model.VendorResponse]
		Create(ctx context.Context, request *model.VendorRequest) model.InterfaceResponse[*model.VendorResponse]
		Patch(ctx context.Context, id string, request *model.VendorRequest) model.InterfaceResponse[*model.VendorResponse]
		Delete(ctx context.Context, id string) model.InterfaceResponse[any]
		ValidateFormRequest(ctx context.Context, request *model.VendorRequest) model.InterfaceResponse[any]
	}

	vendorStruct struct {
		VendorRepository repositories.VendorRepository
		Log              *logrus.Logger
		Validator        *validator.Validate
	}
)

// FindAll implements VendorService.
func (service *vendorStruct) FindAll(ctx context.Context, queryString model.QueryString[string]) model.InterfaceResponse[[]model.VendorResponse] {
	response := model.InterfaceResponse[[]model.VendorResponse]{
		Status: common.StatusInternalServerError,
		Error:  common.ErrUnknown,
	}
	repositoryResponse := service.VendorRepository.FindAll(ctx, queryString)
	if repositoryResponse.Error != nil {
		response.Status = repositoryResponse.Status
		response.Error = repositoryResponse.Error
		response.MetaPagination = repositoryResponse.MetaPagination
		return response
	}
	data := transformer.ToVendors(repositoryResponse.Data)
	response.Status = repositoryResponse.Status
	response.MetaPagination = repositoryResponse.MetaPagination
	response.Data = data
	response.Error = nil
	return response
}

// FindByID implements VendorService.
func (service *vendorStruct) FindByID(ctx context.Context, id string) model.InterfaceResponse[*model.VendorResponse] {
	response := model.InterfaceResponse[*model.VendorResponse]{
		Status: common.StatusInternalServerError,
		Error:  common.ErrUnknown,
	}
	repositoryResponse := service.VendorRepository.FindByID(ctx, id)
	if repositoryResponse.Error != nil {
		response.Status = repositoryResponse.Status
		response.Error = repositoryResponse.Error
		return response
	}
	data := transformer.ToVendor(repositoryResponse.Data)
	response.Status = repositoryResponse.Status
	response.Data = data
	response.Error = nil
	return response
}

// Create implements VendorService.
func (service *vendorStruct) Create(ctx context.Context, request *model.VendorRequest) model.InterfaceResponse[*model.VendorResponse] {
	response := model.InterfaceResponse[*model.VendorResponse]{
		Status: common.StatusInternalServerError,
		Error:  common.ErrUnknown,
	}

	email := request.Email
	name := request.Name
	address := request.Address
	phone := request.Phone
	brand := request.Brand
	picName := request.PICName
	picPhone := request.PICPhone
	lastSeen := time.Now()

	entry := &entity.Vendor{
		Name:     name,
		Email:    email,
		Address:  address,
		Phone:    phone,
		Brand:    brand,
		PICName:  picName,
		PICPhone: picPhone,
		LastSeen: &lastSeen,
	}

	repositoryResponse := service.VendorRepository.Create(ctx, entry)
	if repositoryResponse.Error != nil {
		response.Status = repositoryResponse.Status
		response.Error = repositoryResponse.Error
		return response
	}
	response.Status = repositoryResponse.Status
	response.Error = nil
	return response
}

// Patch implements VendorService.
func (service *vendorStruct) Patch(ctx context.Context, id string, request *model.VendorRequest) model.InterfaceResponse[*model.VendorResponse] {
	response := model.InterfaceResponse[*model.VendorResponse]{
		Status: common.StatusInternalServerError,
		Error:  common.ErrUnknown,
	}

	email := request.Email
	name := request.Name
	address := request.Address
	phone := request.Phone
	brand := request.Brand
	picName := request.PICName
	picPhone := request.PICPhone

	entry := map[string]interface{}{
		"name":     name,
		"phone":    phone,
		"brand":    brand,
		"email":    email,
		"picName":  picName,
		"picPhone": picPhone,
		"address":  address,
	}

	repositoryResponse := service.VendorRepository.Patch(ctx, id, entry)
	if repositoryResponse.Error != nil {
		response.Status = repositoryResponse.Status
		response.Error = repositoryResponse.Error
		return response
	}
	response.Status = repositoryResponse.Status
	response.Error = nil
	return response
}

// Delete implements VendorService.
func (service *vendorStruct) Delete(ctx context.Context, id string) model.InterfaceResponse[any] {
	response := model.InterfaceResponse[any]{
		Status: common.StatusInternalServerError,
		Error:  common.ErrUnknown,
	}

	repositoryResponse := service.VendorRepository.Delete(ctx, id)
	if repositoryResponse.Error != nil {
		response.Status = repositoryResponse.Status
		response.Error = repositoryResponse.Error
		return response
	}
	response.Status = repositoryResponse.Status
	response.Error = nil
	return response
}

// ValidateFormRequest implements VendorService.
func (service *vendorStruct) ValidateFormRequest(ctx context.Context, request *model.VendorRequest) model.InterfaceResponse[any] {
	response := model.InterfaceResponse[any]{
		Status: common.StatusInternalServerError,
		Error:  common.ErrValidateRequest,
	}

	err, msg := common.Validate(service.Validator, request)
	if err != nil {
		response.Status = common.StatusBadRequest
		response.Error = err
		response.Data = msg
		return response
	}
	response.Status = common.StatusOK
	response.Error = nil
	return response
}

func NewVendorService(
	vendorRepository repositories.VendorRepository,
	log *logrus.Logger,
	validator *validator.Validate,
) VendorService {
	return &vendorStruct{
		VendorRepository: vendorRepository,
		Log:              log,
		Validator:        validator,
	}
}
