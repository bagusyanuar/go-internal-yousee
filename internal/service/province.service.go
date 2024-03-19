package service

import (
	"context"

	"github.com/bagusyanuar/go-internal-yousee/internal/model"
	"github.com/bagusyanuar/go-internal-yousee/internal/model/transformer"
	"github.com/bagusyanuar/go-internal-yousee/internal/repositories"
	"github.com/sirupsen/logrus"
)

type (
	ProvinceService interface {
		FindAll(ctx context.Context, queryString model.QueryString[string]) model.InterfaceResponse[[]model.ProvinceResponse]
		FindByID(ctx context.Context, id string) model.InterfaceResponse[*model.ProvinceResponse]
	}

	province struct {
		ProvinceRepository repositories.ProvinceRepository
		Log                *logrus.Logger
	}
)

// FindAll implements ProvinceService.
func (service *province) FindAll(ctx context.Context, queryString model.QueryString[string]) model.InterfaceResponse[[]model.ProvinceResponse] {
	response := service.ProvinceRepository.FindAll(ctx, queryString)
	if response.Error != nil {
		return model.InterfaceResponse[[]model.ProvinceResponse]{
			Status:         response.Status,
			Error:          response.Error,
			MetaPagination: response.MetaPagination,
		}
	}
	provinces := transformer.ToProvinces(response.Data)
	return model.InterfaceResponse[[]model.ProvinceResponse]{
		Data:           provinces,
		Status:         response.Status,
		Error:          response.Error,
		MetaPagination: response.MetaPagination,
	}
}

// FindByID implements ProvinceService.
func (service *province) FindByID(ctx context.Context, id string) model.InterfaceResponse[*model.ProvinceResponse] {
	response := service.ProvinceRepository.FindByID(ctx, id)
	if response.Error != nil {
		return model.InterfaceResponse[*model.ProvinceResponse]{
			Status: response.Status,
			Error:  response.Error,
		}
	}
	data := transformer.ToProvince(response.Data)
	return model.InterfaceResponse[*model.ProvinceResponse]{
		Data:   data,
		Status: response.Status,
		Error:  response.Error,
	}
}

func NewProvinceService(provinceRepository repositories.ProvinceRepository, log *logrus.Logger) ProvinceService {
	return &province{
		ProvinceRepository: provinceRepository,
		Log:                log,
	}
}
