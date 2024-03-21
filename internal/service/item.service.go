package service

import (
	"context"

	"github.com/bagusyanuar/go-internal-yousee/common"
	"github.com/bagusyanuar/go-internal-yousee/internal/entity"
	"github.com/bagusyanuar/go-internal-yousee/internal/model"
	"github.com/bagusyanuar/go-internal-yousee/internal/model/transformer"
	"github.com/bagusyanuar/go-internal-yousee/internal/repositories"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
)

type (
	ItemService interface {
		FindAll(ctx context.Context, queryString model.QueryString[model.ItemQueryString]) model.InterfaceResponse[[]model.ItemResponse]
		FindByID(ctx context.Context, id string) model.InterfaceResponse[*model.ItemResponse]
		Create(ctx context.Context, request *model.ItemRequest) model.InterfaceResponse[*model.ItemResponse]
		Patch(ctx context.Context, id string, request *model.ItemRequest) model.InterfaceResponse[*model.ItemResponse]
		ValidateFormRequest(ctx context.Context, request *model.ItemRequest) model.InterfaceResponse[any]
	}

	itemStruct struct {
		ItemRepository repositories.ItemRepository
		Log            *logrus.Logger
		Validator      *validator.Validate
	}
)

// FindAll implements ItemService.
func (service *itemStruct) FindAll(ctx context.Context, queryString model.QueryString[model.ItemQueryString]) model.InterfaceResponse[[]model.ItemResponse] {
	response := model.InterfaceResponse[[]model.ItemResponse]{
		Status: common.StatusInternalServerError,
		Error:  common.ErrUnknown,
	}
	repositoryResponse := service.ItemRepository.FindAll(ctx, queryString)
	if repositoryResponse.Error != nil {
		response.Status = repositoryResponse.Status
		response.Error = repositoryResponse.Error
		response.MetaPagination = repositoryResponse.MetaPagination
		return response
	}

	data := transformer.ToItems(repositoryResponse.Data)
	response.Status = repositoryResponse.Status
	response.MetaPagination = repositoryResponse.MetaPagination
	response.Data = data
	response.Error = nil
	return response
}

// FindByID implements ItemService.
func (service *itemStruct) FindByID(ctx context.Context, id string) model.InterfaceResponse[*model.ItemResponse] {
	response := model.InterfaceResponse[*model.ItemResponse]{
		Status: common.StatusInternalServerError,
		Error:  common.ErrUnknown,
	}

	repositoryResponse := service.ItemRepository.FindByID(ctx, id)
	if repositoryResponse.Error != nil {
		response.Status = repositoryResponse.Status
		response.Error = repositoryResponse.Error
		return response
	}
	data := transformer.ToItem(repositoryResponse.Data)
	response.Status = repositoryResponse.Status
	response.Data = data
	response.Error = nil
	return response
}

// Create implements ItemService.
func (service *itemStruct) Create(ctx context.Context, request *model.ItemRequest) model.InterfaceResponse[*model.ItemResponse] {
	response := model.InterfaceResponse[*model.ItemResponse]{
		Status: common.StatusInternalServerError,
		Error:  common.ErrUnknown,
	}

	typeID := request.TypeID
	cityID := request.CityID
	vendorID := request.VendorID
	name := request.Name
	address := request.Address
	latitude := request.Latitude
	longitude := request.Longitude
	url := request.URL
	width := request.Width
	height := request.Height
	position := request.Position

	entry := &entity.Item{
		TypeID:    typeID,
		CityID:    cityID,
		VendorID:  vendorID,
		Name:      name,
		Address:   address,
		Latitude:  latitude,
		Longitude: longitude,
		URL:       &url,
		Width:     width,
		Height:    height,
		Position:  position,
	}
	repositoryResponse := service.ItemRepository.Create(ctx, entry)
	if repositoryResponse.Error != nil {
		response.Status = repositoryResponse.Status
		response.Error = repositoryResponse.Error
		return response
	}
	response.Status = repositoryResponse.Status
	response.Error = nil
	return response
}

// Patch implements ItemService.
func (service *itemStruct) Patch(ctx context.Context, id string, request *model.ItemRequest) model.InterfaceResponse[*model.ItemResponse] {
	response := model.InterfaceResponse[*model.ItemResponse]{
		Status: common.StatusInternalServerError,
		Error:  common.ErrUnknown,
	}

	typeID := request.TypeID
	cityID := request.CityID
	vendorID := request.VendorID
	name := request.Name
	address := request.Address
	latitude := request.Latitude
	longitude := request.Longitude
	url := request.URL
	width := request.Width
	height := request.Height
	position := request.Position

	entry := map[string]interface{}{
		"type_id":   typeID,
		"city_id":   cityID,
		"vendor_id": vendorID,
		"name":      name,
		"address":   address,
		"latitude":  latitude,
		"longitude": longitude,
		"url":       url,
		"width":     width,
		"height":    height,
		"position":  position,
	}

	repositoryResponse := service.ItemRepository.Patch(ctx, id, entry)
	if repositoryResponse.Error != nil {
		response.Status = repositoryResponse.Status
		response.Error = repositoryResponse.Error
		return response
	}
	response.Status = repositoryResponse.Status
	response.Error = nil
	return response
}

// ValidateFormRequest implements ItemService.
func (service *itemStruct) ValidateFormRequest(ctx context.Context, request *model.ItemRequest) model.InterfaceResponse[any] {
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

func NewItemService(
	itemRepository repositories.ItemRepository,
	log *logrus.Logger,
	validator *validator.Validate,
) ItemService {
	return &itemStruct{
		ItemRepository: itemRepository,
		Log:            log,
		Validator:      validator,
	}
}
