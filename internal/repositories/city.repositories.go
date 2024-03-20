package repositories

import (
	"context"
	"errors"

	"github.com/bagusyanuar/go-internal-yousee/common"
	"github.com/bagusyanuar/go-internal-yousee/internal/entity"
	"github.com/bagusyanuar/go-internal-yousee/internal/model"
	"github.com/bagusyanuar/go-internal-yousee/internal/model/transformer"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type (
	CityRepository interface {
		FindAll(ctx context.Context, queryString model.QueryString[model.CityQueryString]) model.InterfaceResponse[[]entity.City]
		FindByID(ctx context.Context, id string) model.InterfaceResponse[*entity.City]
	}

	city struct {
		DB  *gorm.DB
		Log *logrus.Logger
	}
)

// FindAll implements CityRepository.
func (repository *city) FindAll(ctx context.Context, queryString model.QueryString[model.CityQueryString]) model.InterfaceResponse[[]entity.City] {
	var data []entity.City
	queryName := queryString.Query.Name
	queryProvince := queryString.Query.Province
	metaPagination := new(model.MetaPagination)
	paginate := &common.Pagination{
		Limit: queryString.QueryPagination.PerPage,
		Page:  queryString.QueryPagination.Page,
	}
	response := model.InterfaceResponse[[]entity.City]{
		Status:         common.StatusUnProccessableEntity,
		MetaPagination: metaPagination,
	}

	tx := repository.DB.WithContext(ctx)

	if queryName != "" {
		q := "%" + queryName + "%"
		tx = tx.Where("name LIKE ?", q)
	}

	if queryProvince != "" {
		tx = tx.Where("province_id = ?", queryProvince)
	}

	if err := tx.
		Preload("Province").
		Scopes(common.Paginate(data, paginate, tx)).
		Find(&data).Error; err != nil {
		response.Error = err
		return response
	}

	response.Status = common.StatusOK
	response.Data = data
	response.MetaPagination = transformer.ToMetaPagination(paginate)
	return response
}

// FindByID implements CityRepository.
func (repository *city) FindByID(ctx context.Context, id string) model.InterfaceResponse[*entity.City] {
	var data *entity.City
	response := model.InterfaceResponse[*entity.City]{
		Status: common.StatusUnProccessableEntity,
	}
	tx := repository.DB.WithContext(ctx)
	if err := tx.
		Preload("Province").
		Where("id = ?", id).
		First(&data).Error; err != nil {
		repository.Log.Warnf("query failed : %+v", err)
		if errors.Is(gorm.ErrRecordNotFound, err) {
			response.Status = common.StatusNotFound
			response.Error = common.ErrRecordNotFound
			return response
		}
		response.Error = err
		return response
	}
	response.Status = common.StatusOK
	response.Data = data
	return response
}

func NewCityRepository(db *gorm.DB, log *logrus.Logger) CityRepository {
	return &city{
		DB:  db,
		Log: log,
	}
}
