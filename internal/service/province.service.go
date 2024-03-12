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
		FindAll(ctx context.Context, queryString model.QueryString[string]) (model.Response[[]model.ProvinceResponse], error)
	}

	province struct {
		ProvinceRepository repositories.ProvinceRepository
		Log                *logrus.Logger
	}
)

// FindAll implements ProvinceService.
func (service *province) FindAll(ctx context.Context, queryString model.QueryString[string]) (model.Response[[]model.ProvinceResponse], error) {
	var provinces []model.ProvinceResponse
	response, err := service.ProvinceRepository.FindAll(ctx, queryString)
	if err != nil {
		return model.Response[[]model.ProvinceResponse]{}, err
	}
	provinces = transformer.ToProvinces(response.Data)
	return model.Response[[]model.ProvinceResponse]{Data: provinces, Meta: response.Meta}, nil
}

func NewProvinceService(provinceRepository repositories.ProvinceRepository, log *logrus.Logger) ProvinceService {
	return &province{
		ProvinceRepository: provinceRepository,
		Log:                log,
	}
}
