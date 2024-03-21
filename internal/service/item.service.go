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
		Create(ctx context.Context, request *model.ItemRequest) (any, error)
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
func (service *itemStruct) Create(ctx context.Context, request *model.ItemRequest) (any, error) {
	//validate form request
	errValidation, msg := common.Validate(service.Validator, request)
	if errValidation != nil {
		return msg, common.ErrBadRequest
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

	entity := &entity.Item{
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
	err := service.ItemRepository.Create(ctx, entity)
	if err != nil {
		return nil, err
	}
	return nil, nil
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
