package service

import (
	"context"

	"github.com/bagusyanuar/go-internal-yousee/common"
	"github.com/bagusyanuar/go-internal-yousee/internal/model"
	"github.com/bagusyanuar/go-internal-yousee/internal/repositories"
	"github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
)

type (
	DashboardService interface {
		GetDashboardStatisticInfo(ctx context.Context) model.InterfaceResponse[[]model.DashboardStatisticInfoResponse]
	}

	dashboardStruct struct {
		DashboardRepository repositories.DashboardRepository
		Log                 *logrus.Logger
	}
)

// GetDashboardStatisticInfo implements DashboardService.
func (service *dashboardStruct) GetDashboardStatisticInfo(ctx context.Context) model.InterfaceResponse[[]model.DashboardStatisticInfoResponse] {
	response := model.InterfaceResponse[[]model.DashboardStatisticInfoResponse]{
		Status: common.StatusInternalServerError,
		Error:  common.ErrUnknown,
	}

	var itemCount int64
	var vendorCount int64
	var eg errgroup.Group

	eg.Go(func() error {
		repositoryResponse := service.DashboardRepository.GetCountItem(ctx)
		if repositoryResponse.Error != nil {
			return repositoryResponse.Error
		}
		itemCount = repositoryResponse.Data
		response.Error = repositoryResponse.Error
		response.Status = repositoryResponse.Status
		return nil
	})

	eg.Go(func() error {
		repositoryResponse := service.DashboardRepository.GetCountVendor(ctx)
		if repositoryResponse.Error != nil {
			return repositoryResponse.Error
		}
		vendorCount = repositoryResponse.Data
		response.Error = repositoryResponse.Error
		response.Status = repositoryResponse.Status
		return nil
	})

	if err := eg.Wait(); err != nil {
		response.Error = err
		return response
	}

	var data []model.DashboardStatisticInfoResponse
	data = append(data, model.DashboardStatisticInfoResponse{
		Name:  "item",
		Value: itemCount,
	})

	data = append(data, model.DashboardStatisticInfoResponse{
		Name:  "vendor",
		Value: vendorCount,
	})
	response.Status = common.StatusOK
	response.Data = data
	response.Error = nil
	return response
}

func NewDashboardService(dashboardRepository repositories.DashboardRepository, log *logrus.Logger) DashboardService {
	return &dashboardStruct{
		DashboardRepository: dashboardRepository,
		Log:                 log,
	}
}
