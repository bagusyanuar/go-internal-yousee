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
	ItemService interface {
		FindAll(ctx context.Context, queryString model.QueryString[string]) (model.Response[[]model.ItemResponse], error)
		FindByID(ctx context.Context, id string) (*model.ItemResponse, error)
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
