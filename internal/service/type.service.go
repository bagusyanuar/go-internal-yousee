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
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

const (
	IconPath = "assets/type"
)

type (
	TypeService interface {
		FindAll(ctx context.Context, queryString model.QueryString[string]) model.InterfaceResponse[[]model.TypeResponse]
		FindByID(ctx context.Context, id string) model.InterfaceResponse[*model.TypeResponse]
		Create(ctx context.Context, request *model.TypeRequest) model.InterfaceResponse[*model.TypeResponse]
		Patch(ctx context.Context, id string, request *model.TypeRequest) model.InterfaceResponse[*model.TypeResponse]
		Delete(ctx context.Context, id string) error
	}

	typeStruct struct {
		TypeRepository repositories.TypeRepository
		Log            *logrus.Logger
		Validator      *validator.Validate
	}
)

// FindAll implements TypeService.
func (service *typeStruct) FindAll(ctx context.Context, queryString model.QueryString[string]) model.InterfaceResponse[[]model.TypeResponse] {
	// var types []model.TypeResponse
	response := service.TypeRepository.FindAll(ctx, queryString)
	if response.Error != nil {
		return model.InterfaceResponse[[]model.TypeResponse]{
			Status:         response.Status,
			Error:          response.Error,
			MetaPagination: response.MetaPagination,
		}
	}
	types := transformer.ToTypes(response.Data)
	return model.InterfaceResponse[[]model.TypeResponse]{
		Data:           types,
		Status:         response.Status,
		Error:          response.Error,
		MetaPagination: response.MetaPagination,
	}
}

// FindByID implements TypeService.
func (service *typeStruct) FindByID(ctx context.Context, id string) model.InterfaceResponse[*model.TypeResponse] {
	response := service.TypeRepository.FindByID(ctx, id)
	if response.Error != nil {
		return model.InterfaceResponse[*model.TypeResponse]{
			Status: response.Status,
			Error:  response.Error,
		}
	}
	data := transformer.ToType(response.Data)
	return model.InterfaceResponse[*model.TypeResponse]{
		Data:   data,
		Status: response.Status,
		Error:  response.Error,
	}
}

// Create implements TypeService.
func (service *typeStruct) Create(ctx context.Context, request *model.TypeRequest) model.InterfaceResponse[*model.TypeResponse] {
	errValidation, msg := common.Validate(service.Validator, request)
	if errValidation != nil {
		return model.InterfaceResponse[*model.TypeResponse]{
			Status:     common.StatusBadRequest,
			Error:      errValidation,
			Validation: msg,
		}
	}
	//upload icon
	icon, err := service.upload(request.Icon)
	if err != nil {
		return model.InterfaceResponse[*model.TypeResponse]{
			Status: common.StatusInternalServerError,
			Error:  err,
		}
	}

	name := request.Name
	entity := &entity.Type{
		Name: name,
		Icon: icon,
	}

	response := service.TypeRepository.Create(ctx, entity)
	if response.Error != nil {
		return model.InterfaceResponse[*model.TypeResponse]{
			Status: response.Status,
			Error:  response.Error,
		}
	}
	data := transformer.ToType(response.Data)
	return model.InterfaceResponse[*model.TypeResponse]{
		Status: response.Status,
		Error:  nil,
		Data:   data,
	}
}

// Patch implements TypeService.
func (service *typeStruct) Patch(ctx context.Context, id string, request *model.TypeRequest) model.InterfaceResponse[*model.TypeResponse] {
	errValidation, msg := common.Validate(service.Validator, request)
	if errValidation != nil {
		return model.InterfaceResponse[*model.TypeResponse]{
			Status:     common.StatusBadRequest,
			Error:      errValidation,
			Validation: msg,
		}
	}

	//upload icon
	icon, err := service.upload(request.Icon)
	if err != nil {
		return model.InterfaceResponse[*model.TypeResponse]{
			Status: common.StatusInternalServerError,
			Error:  err,
		}
	}

	name := request.Name
	data := map[string]interface{}{
		"name": name,
	}
	if icon != nil {
		data = map[string]interface{}{
			"name": name,
			"icon": icon,
		}
	}

	response := service.TypeRepository.Patch(ctx, id, data)
	if response.Error != nil {
		return model.InterfaceResponse[*model.TypeResponse]{
			Status: response.Status,
			Error:  response.Error,
		}
	}
	item := transformer.ToType(response.Data)
	return model.InterfaceResponse[*model.TypeResponse]{
		Status: response.Status,
		Error:  nil,
		Data:   item,
	}
}

// Delete implements TypeService.
func (service *typeStruct) Delete(ctx context.Context, id string) error {
	err := service.TypeRepository.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func (service *typeStruct) upload(icon *multipart.FileHeader) (*string, error) {

	var iconName *string
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

func NewItemTypeService(typeRepository repositories.TypeRepository, log *logrus.Logger, validator *validator.Validate) TypeService {
	return &typeStruct{
		TypeRepository: typeRepository,
		Log:            log,
		Validator:      validator,
	}
}
