package service

import (
	"context"
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
	response := model.InterfaceResponse[[]model.DashboardStatisticInfoResponse]{
		Status: common.StatusInternalServerError,
		Error:  common.ErrUnknown,
	}

	var wg sync.WaitGroup

	var chItemCount = make(chan int64)
	var chVendorCount = make(chan int64)

	wg.Add(2)
	go service.doCountItem(&wg, chItemCount)
	go service.doCountVendor(&wg, chVendorCount)
	// r1 := service.DashboardRepository.GetCountItem(ctx)
	// if r1.Error != nil {
	// 	response.Status = r1.Status
	// 	response.Error = r1.Error
	// 	return response
	// }
	// r2 := service.DashboardRepository.GetCountVendor(ctx)
	// if r1.Error != nil {
	// 	response.Status = r2.Status
	// 	response.Error = r2.Error
	// 	return response
	// }

	go func() {
		wg.Wait()
	}()

	itemCount := <-chItemCount
	vendorCount := <-chVendorCount
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

func (service *dashboardStruct) doCountItem(wg *sync.WaitGroup, value chan int64) {
	defer func() {
		close(value)
		wg.Done()
	}()
	time.Sleep(time.Millisecond * 300)
	value <- 10
	service.Log.Warnf("do count item")
}

func (service *dashboardStruct) doCountVendor(wg *sync.WaitGroup, value chan int64) {
	defer func() {
		close(value)
		wg.Done()
	}()
	time.Sleep(time.Millisecond * 500)
	value <- 15
	service.Log.Warnf("do count vendor")
}
func NewDashboardService(dashboardRepository repositories.DashboardRepository, log *logrus.Logger) DashboardService {
	return &dashboardStruct{
		DashboardRepository: dashboardRepository,
		Log:                 log,
	}
}
