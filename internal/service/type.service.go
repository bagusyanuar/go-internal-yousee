package service

import (
	"context"
	"fmt"
	"mime/multipart"
	"path/filepath"

	"github.com/bagusyanuar/go-internal-yousee/common"
	"github.com/bagusyanuar/go-internal-yousee/internal/entity"
	"github.com/bagusyanuar/go-internal-yousee/internal/model"
	"github.com/bagusyanuar/go-internal-yousee/internal/model/transformer"
	"github.com/bagusyanuar/go-internal-yousee/internal/repositories"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

const (
	IconPath = "assets/type"
)

type (
	TypeService interface {
		FindAll(ctx context.Context, param string) ([]model.TypeResponse, error)
		FindByID(ctx context.Context, id string) (*model.TypeResponse, error)
		Create(ctx context.Context, request *model.TypeRequest) error
		Patch(ctx context.Context, id string, request *model.TypeRequest) error
		Delete(ctx context.Context, id string) error
	}

	itemType struct {
		TypeRepository repositories.TypeRepository
		Log            *logrus.Logger
	}
)

// Delete implements TypeService.
func (service *itemType) Delete(ctx context.Context, id string) error {
	err := service.TypeRepository.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

// Patch implements TypeService.
func (service *itemType) Patch(ctx context.Context, id string, request *model.TypeRequest) error {
	name := request.Name
	entity := &entity.Type{
		Name: name,
	}

	if request.Icon != nil {
		icon, err := service.upload(request.Icon)
		if err != nil {
			return err
		}
		entity.Icon = icon
	}

	err := service.TypeRepository.Patch(ctx, id, entity)
	if err != nil {
		return err
	}
	return nil
}

// FindByID implements TypeService.
func (service *itemType) FindByID(ctx context.Context, id string) (*model.TypeResponse, error) {
	entity, err := service.TypeRepository.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return transformer.TypeToResponse(entity), nil
}

// Create implements TypeService.
func (service *itemType) Create(ctx context.Context, request *model.TypeRequest) error {
	name := request.Name

	icon, err := service.upload(request.Icon)

	if err != nil {
		return err
	}

	entity := &entity.Type{
		Name: name,
		Icon: icon,
	}

	err = service.TypeRepository.Create(ctx, entity)
	if err != nil {
		return err
	}
	return nil
}

// FindAll implements TypeService.
func (service *itemType) FindAll(ctx context.Context, param string) ([]model.TypeResponse, error) {
	var results []model.TypeResponse
	entities, err := service.TypeRepository.FindAll(ctx, param)
	if err != nil {
		return results, err
	}
	for _, entity := range entities {
		t := *transformer.TypeToResponse(&entity)
		results = append(results, t)
	}
	return results, nil
}

func (service *itemType) upload(icon *multipart.FileHeader) (*string, error) {

	iconName := new(string)
	if icon != nil {
		fileSystem := common.FileSystem{
			File: icon,
		}
		if err := fileSystem.CheckPath(IconPath); err != nil {
			return nil, err
		}

		ext := filepath.Ext(icon.Filename)
		fileName := fmt.Sprintf("%s/%s%s", IconPath, uuid.New().String(), ext)
		iconName = &fileName
		err := fileSystem.Upload(fileName)
		if err != nil {
			return nil, err
		}
	}
	return iconName, nil
}

func NewItemTypeService(typeRepository repositories.TypeRepository, log *logrus.Logger) TypeService {
	return &itemType{
		TypeRepository: typeRepository,
		Log:            log,
	}
}
