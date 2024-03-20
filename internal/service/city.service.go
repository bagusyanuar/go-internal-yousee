package service

import (
	"context"

	"github.com/bagusyanuar/go-internal-yousee/common"
	"github.com/bagusyanuar/go-internal-yousee/internal/model"
	"github.com/bagusyanuar/go-internal-yousee/internal/model/transformer"
	"github.com/bagusyanuar/go-internal-yousee/internal/repositories"
	"github.com/sirupsen/logrus"
)

type (
	CityService interface {
		FindAll(ctx context.Context, queryString model.QueryString[model.CityQueryString]) model.InterfaceResponse[[]model.CityResponse]
		FindByID(ctx context.Context, id string) model.InterfaceResponse[*model.CityResponse]
	}

	city struct {
		CityRepository repositories.CityRepository
		Log            *logrus.Logger
	}
)

// FindAll implements CityService.
func (service *city) FindAll(ctx context.Context, queryString model.QueryString[model.CityQueryString]) model.InterfaceResponse[[]model.CityResponse] {
	response := model.InterfaceResponse[[]model.CityResponse]{
		Status: common.StatusInternalServerError,
		Error:  common.ErrUnknown,
	}
	repositoryResponse := service.CityRepository.FindAll(ctx, queryString)
	if repositoryResponse.Error != nil {
		response.Status = repositoryResponse.Status
		response.Error = repositoryResponse.Error
		response.MetaPagination = repositoryResponse.MetaPagination
		return response
	}
	data := transformer.ToCities(repositoryResponse.Data)

	response.Status = repositoryResponse.Status
	response.MetaPagination = repositoryResponse.MetaPagination
	response.Data = data
	response.Error = nil
	return response
}

// FindByID implements CityService.
func (service *city) FindByID(ctx context.Context, id string) model.InterfaceResponse[*model.CityResponse] {
	response := model.InterfaceResponse[*model.CityResponse]{
		Status: common.StatusInternalServerError,
		Error:  common.ErrUnknown,
	}
	repositoryResponse := service.CityRepository.FindByID(ctx, id)
	if repositoryResponse.Error != nil {
		response.Status = repositoryResponse.Status
		response.Error = repositoryResponse.Error
		return response
	}

	data := transformer.ToCity(repositoryResponse.Data)
	response.Status = repositoryResponse.Status
	response.Data = data
	response.Error = nil
	return response
}

func NewCityService(cityRepository repositories.CityRepository, log *logrus.Logger) CityService {
	return &city{
		CityRepository: cityRepository,
		Log:            log,
	}
}
