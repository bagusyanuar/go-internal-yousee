package service

import (
	"context"
	"time"

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
	serviceContext, cancel := context.WithCancel(context.Background())
	defer cancel()
	eg, errorGroupCtx := errgroup.WithContext(serviceContext)

	eg.Go(func() error {
		val, err := service.doCountItemWithError(errorGroupCtx)
		if err != nil {
			cancel()
			return err
		}
		service.Log.Warnf("do count item")
		itemCount = val
		return nil
	})
	eg.Go(func() error {
		val, err := service.doCountVendorWithError(errorGroupCtx)
		if err != nil {
			cancel()
			return err
		}
		service.Log.Warnf("do count vendor")
		vendorCount = val
		return nil
	})

	if err := eg.Wait(); err != nil {
		response.Error = err
		response.Status = common.StatusInternalServerError
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

func (service *dashboardStruct) doCountItemWithError(ctx context.Context) (int64, error) {
	time.Sleep(time.Millisecond * 400)
	return 10, common.ErrUnknown
}

func (service *dashboardStruct) doCountVendorWithError(ctx context.Context) (int64, error) {
	time.Sleep(time.Millisecond * 800)
	return 15, nil
}

func NewDashboardService(dashboardRepository repositories.DashboardRepository, log *logrus.Logger) DashboardService {
	return &dashboardStruct{
		DashboardRepository: dashboardRepository,
		Log:                 log,
	}
}
