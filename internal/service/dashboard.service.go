package service

import (
	"context"
	"errors"
	"math/rand"
	"sync"
	"time"

	"github.com/bagusyanuar/go-internal-yousee/common"
	"github.com/bagusyanuar/go-internal-yousee/internal/model"
	"github.com/bagusyanuar/go-internal-yousee/internal/repositories"
	"github.com/sirupsen/logrus"
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
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	response := model.InterfaceResponse[[]model.DashboardStatisticInfoResponse]{
		Status: common.StatusInternalServerError,
		Error:  common.ErrUnknown,
	}

	var itemCount int64
	var vendorCount int64
	var wg sync.WaitGroup
	chanErrors := make(chan error, 4)
	wg.Add(4)

	go service.doCountItemWithError(&wg, chanErrors, false, cancel, ctx, 1)
	go service.doCountItemWithError(&wg, chanErrors, true, cancel, ctx, 2)
	go service.doCountItemWithError(&wg, chanErrors, false, cancel, ctx, 3)
	go service.doCountItemWithError(&wg, chanErrors, false, cancel, ctx, 4)
	go func() {
		wg.Wait()
		close(chanErrors)
	}()

	select {
	case <-ctx.Done():
		response.Error = ctx.Err()
	case err := <-chanErrors:
		response.Error = err
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

func (service *dashboardStruct) doCountItemWithError(wg *sync.WaitGroup, err chan error, stat bool, cancel context.CancelFunc, ctx context.Context, key int) {
	defer wg.Done()
	if stat {
		e := errors.New("asd")
		select {
		case err <- e:
			cancel()
			return
		case <-ctx.Done():
			return
		}
	} else {
		service.Log.Warnf("success fetch %+v", key)
	}
}

func (service *dashboardStruct) doCountVendorWithError(ctx context.Context) (int64, error) {
	// time.Sleep(time.Millisecond * 800)
	r := rand.Intn(100)
	time.Sleep(time.Duration(r) * time.Millisecond)
	return 15, common.ErrUnknown
}

func NewDashboardService(dashboardRepository repositories.DashboardRepository, log *logrus.Logger) DashboardService {
	return &dashboardStruct{
		DashboardRepository: dashboardRepository,
		Log:                 log,
	}
}
