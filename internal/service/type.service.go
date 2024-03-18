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
		FindAll(ctx context.Context, queryString model.QueryString[string]) (model.Response[[]model.TypeResponse], int, error)
		FindByID(ctx context.Context, id string) (*model.TypeResponse, int, error)
		Create(ctx context.Context, request *model.TypeRequest) (code int, validationMessage any, e error)
		Patch(ctx context.Context, id string, request *model.TypeRequest) error
		Delete(ctx context.Context, id string) error
	}

	typeStruct struct {
		TypeRepository repositories.TypeRepository
		Log            *logrus.Logger
		Validator      *validator.Validate
	}
)

// FindAll implements TypeService.
func (service *typeStruct) FindAll(ctx context.Context, queryString model.QueryString[string]) (model.Response[[]model.TypeResponse], int, error) {
	var types []model.TypeResponse
	response, code, err := service.TypeRepository.FindAll(ctx, queryString)
	if err != nil {
		return model.Response[[]model.TypeResponse]{}, code, err
	}
	types = transformer.ToTypes(response.Data)
	return model.Response[[]model.TypeResponse]{
		Data: types,
		Meta: response.Meta,
	}, code, nil
}

// FindByID implements TypeService.
func (service *typeStruct) FindByID(ctx context.Context, id string) (*model.TypeResponse, int, error) {
	response, code, err := service.TypeRepository.FindByID(ctx, id)
	if err != nil {
		return nil, code, err
	}
	return transformer.ToType(response), code, nil
}

// Create implements TypeService.
func (service *typeStruct) Create(ctx context.Context, request *model.TypeRequest) (int, any, error) {
	errValidation, msg := common.Validate(service.Validator, request)
	if errValidation != nil {
		return common.StatusBadRequest, msg, common.ErrBadRequest
	}
	//upload icon
	icon, err := service.upload(request.Icon)
	if err != nil {
		return common.StatusInternalServerError, nil, err
	}

	name := request.Name
	entity := &entity.Type{
		Name: name,
		Icon: icon,
	}

	code, err := service.TypeRepository.Create(ctx, entity)
	if err != nil {
		return code, nil, err
	}
	return common.StatusOK, nil, nil
}

// Delete implements TypeService.
func (service *typeStruct) Delete(ctx context.Context, id string) error {
	err := service.TypeRepository.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

// Patch implements TypeService.
func (service *typeStruct) Patch(ctx context.Context, id string, request *model.TypeRequest) error {
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

func (service *typeStruct) upload(icon *multipart.FileHeader) (*string, error) {

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

func NewItemTypeService(typeRepository repositories.TypeRepository, log *logrus.Logger, validator *validator.Validate) TypeService {
	return &typeStruct{
		TypeRepository: typeRepository,
		Log:            log,
		Validator:      validator,
	}
}
