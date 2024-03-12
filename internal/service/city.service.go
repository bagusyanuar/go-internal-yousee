package service

import (
	"context"

	"github.com/bagusyanuar/go-internal-yousee/internal/model"
	"github.com/bagusyanuar/go-internal-yousee/internal/model/transformer"
	"github.com/bagusyanuar/go-internal-yousee/internal/repositories"
	"github.com/sirupsen/logrus"
)

type (
	CityService interface {
		FindAll(ctx context.Context, queryString model.QueryString[string]) (model.Response[[]model.CityResponse], error)
		FindByID(ctx context.Context, id string) (*model.CityResponse, error)
	}

	city struct {
		CityRepository repositories.CityRepository
		Log            *logrus.Logger
	}
)

// FindAll implements CityService.
func (service *city) FindAll(ctx context.Context, queryString model.QueryString[string]) (model.Response[[]model.CityResponse], error) {
	var cities []model.CityResponse
	response, err := service.CityRepository.FindAll(ctx, queryString)
	if err != nil {
		return model.Response[[]model.CityResponse]{}, err
	}
	cities = transformer.ToCities(response.Data)
	return model.Response[[]model.CityResponse]{Data: cities, Meta: response.Meta}, nil
}

// FindByID implements CityService.
func (service *city) FindByID(ctx context.Context, id string) (*model.CityResponse, error) {
	entity, err := service.CityRepository.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return transformer.ToCity(entity), nil
}

func NewCityService(cityRepository repositories.CityRepository, log *logrus.Logger) CityService {
	return &city{
		CityRepository: cityRepository,
		Log:            log,
	}
}
