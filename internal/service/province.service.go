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
	response := model.InterfaceResponse[[]model.ProvinceResponse]{
		Status: common.StatusInternalServerError,
		Error:  common.ErrUnknown,
	}

	repositoryResponse := service.ProvinceRepository.FindAll(ctx, queryString)
	if repositoryResponse.Error != nil {
		response.Status = repositoryResponse.Status
		response.Error = repositoryResponse.Error
		response.MetaPagination = repositoryResponse.MetaPagination
		return response
	}
	data := transformer.ToProvinces(repositoryResponse.Data)
	response.Status = repositoryResponse.Status
	response.MetaPagination = repositoryResponse.MetaPagination
	response.Data = data
	response.Error = nil
	return response
}

// FindByID implements ProvinceService.
func (service *province) FindByID(ctx context.Context, id string) model.InterfaceResponse[*model.ProvinceResponse] {
	response := model.InterfaceResponse[*model.ProvinceResponse]{
		Status: common.StatusInternalServerError,
		Error:  common.ErrUnknown,
	}

	repositoryResponse := service.ProvinceRepository.FindByID(ctx, id)
	if repositoryResponse.Error != nil {
		response.Status = repositoryResponse.Status
		response.Error = repositoryResponse.Error
		return response
	}

	data := transformer.ToProvince(repositoryResponse.Data)
	response.Status = repositoryResponse.Status
	response.Data = data
	response.Error = nil
	return response
}

func NewProvinceService(provinceRepository repositories.ProvinceRepository, log *logrus.Logger) ProvinceService {
	return &province{
		ProvinceRepository: provinceRepository,
		Log:                log,
	}
}
