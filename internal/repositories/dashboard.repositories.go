package repositories

import (
	"context"

	"github.com/bagusyanuar/go-internal-yousee/common"
	"github.com/bagusyanuar/go-internal-yousee/internal/entity"
	"github.com/bagusyanuar/go-internal-yousee/internal/model"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type (
	DashboardRepository interface {
		GetCountItem(ctx context.Context) model.InterfaceResponse[int64]
		GetCountVendor(ctx context.Context) model.InterfaceResponse[int64]
	}

	dashboardStruct struct {
		DB  *gorm.DB
		Log *logrus.Logger
	}
)

// GetCountItem implements DashboadRepository.
func (repository *dashboardStruct) GetCountItem(ctx context.Context) model.InterfaceResponse[int64] {
	var value int64
	response := model.InterfaceResponse[int64]{
		Status: common.StatusUnProccessableEntity,
	}

	tx := repository.DB.WithContext(ctx)
	if err := tx.Model(&entity.Item{}).Count(&value).Error; err != nil {
		repository.Log.Warnf("query failed : %+v", err)
		response.Error = err
		return response
	}
	response.Status = common.StatusOK
	response.Data = value
	return response
}

// GetCountVendor implements DashboadRepository.
func (repository *dashboardStruct) GetCountVendor(ctx context.Context) model.InterfaceResponse[int64] {
	var value int64
	response := model.InterfaceResponse[int64]{
		Status: common.StatusUnProccessableEntity,
	}

	tx := repository.DB.WithContext(ctx)
	if err := tx.Model(&entity.Vendor{}).Count(&value).Error; err != nil {
		repository.Log.Warnf("query failed : %+v", err)
		response.Error = err
		return response
	}
	response.Status = common.StatusOK
	response.Data = value
	return response
}

func NewDashboardRepository(db *gorm.DB, log *logrus.Logger) DashboardRepository {
	return &dashboardStruct{
		DB:  db,
		Log: log,
	}
}
