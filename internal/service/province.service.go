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
		FindAll(ctx context.Context, param string, paginateQuery model.PaginationQuery) ([]model.ProvinceResponse, error)
	}

	province struct {
		ProvinceRepository repositories.ProvinceRepository
		Log                *logrus.Logger
	}
)

// FindAll implements ProvinceService.
func (service *province) FindAll(ctx context.Context, param string, paginateQuery model.PaginationQuery) ([]model.ProvinceResponse, error) {
	var results []model.ProvinceResponse
	entities, err := service.ProvinceRepository.FindAll(ctx, param, paginateQuery)
	if err != nil {
		return results, err
	}
	for _, entity := range entities {
		t := *transformer.ProvinceToResponse(&entity)
		results = append(results, t)
	}
	return results, nil
}

func NewProvinceService(provinceRepository repositories.ProvinceRepository, log *logrus.Logger) ProvinceService {
	return &province{
		ProvinceRepository: provinceRepository,
		Log:                log,
	}
}
