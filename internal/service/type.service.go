package service

import (
	"context"

	"github.com/bagusyanuar/go-internal-yousee/internal/entity"
	"github.com/bagusyanuar/go-internal-yousee/internal/model"
	"github.com/bagusyanuar/go-internal-yousee/internal/model/transformer"
	"github.com/bagusyanuar/go-internal-yousee/internal/repositories"
)

type (
	TypeService interface {
		FindAll(ctx context.Context) ([]model.TypeResponse, error)
		Create(ctx context.Context, request *model.TypeRequest) error
	}

	itemType struct {
		TypeRepository repositories.TypeRepository
	}
)

// Create implements TypeService.
func (service *itemType) Create(ctx context.Context, request *model.TypeRequest) error {
	name := request.Name
	icon := request.Icon

	entity := &entity.Type{
		Name: name,
		Icon: icon,
	}

	err := service.TypeRepository.Create(ctx, entity)
	if err != nil {
		return err
	}
	return nil
}

// FindAll implements TypeService.
func (service *itemType) FindAll(ctx context.Context) ([]model.TypeResponse, error) {
	var results []model.TypeResponse
	entities, err := service.TypeRepository.FindAll(ctx)
	if err != nil {
		return results, err
	}
	for _, entity := range entities {
		results = append(results, transformer.TypeToResponse(entity))
	}
	return results, nil
}

func NewItemTypeService(typeRepository repositories.TypeRepository) TypeService {
	return &itemType{
		TypeRepository: typeRepository,
	}
}
