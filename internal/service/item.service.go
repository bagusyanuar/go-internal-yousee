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
		FindAll(ctx context.Context, queryString model.QueryString[string]) (model.Response[[]model.ItemResponse], error)
		FindByID(ctx context.Context, id string) (*model.ItemResponse, error)
		Create(ctx context.Context, request *model.ItemRequest) (any, error)
	}

	itemStruct struct {
		ItemRepository repositories.ItemRepository
		Log            *logrus.Logger
		Validator      *validator.Validate
	}
)

// FindAll implements ItemService.
func (service *itemStruct) FindAll(ctx context.Context, queryString model.QueryString[string]) (model.Response[[]model.ItemResponse], error) {
	var items []model.ItemResponse
	response, err := service.ItemRepository.FindAll(ctx, queryString)
	if err != nil {
		return model.Response[[]model.ItemResponse]{}, err
	}
	items = transformer.ToItems(response.Data)
	return model.Response[[]model.ItemResponse]{Data: items, Meta: response.Meta}, nil
}

// FindByID implements ItemService.
func (service *itemStruct) FindByID(ctx context.Context, id string) (*model.ItemResponse, error) {
	entity, err := service.ItemRepository.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return transformer.ToItem(entity), nil
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
