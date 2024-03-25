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
	eg, ctx := errgroup.WithContext(ctx)

	eg.Go(func() error {
		for {
			select {
			case <-ctx.Done():
				service.Log.Warnf("Sleep 1 cancel..................")
				return nil
			default:
				v, e := service.doJob(ctx, 400)
				itemCount = v
				service.Log.Warnf("Sleep 1 go..................")
				return e
			}
		}

	})

	eg.Go(func() error {
		for {
			select {
			case <-ctx.Done():
				service.Log.Warnf("Sleep 2 cancel..................")
				return nil
			default:
				v, e := service.doJob(ctx, 200)
				itemCount = v
				service.Log.Warnf("Sleep 2 go..................")
				return e
			}
		}
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

func (service *dashboardStruct) doJob(ctx context.Context, key int) (int64, error) {
	// time.Sleep(time.Duration(key) * time.Millisecond)
	if key > 300 {
		return 0, common.ErrBadRequest
	}
	return 0, nil
}

func NewDashboardService(dashboardRepository repositories.DashboardRepository, log *logrus.Logger) DashboardService {
	return &dashboardStruct{
		DashboardRepository: dashboardRepository,
		Log:                 log,
	}
}
