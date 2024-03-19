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
	ProvinceRepository interface {
		FindAll(ctx context.Context, queryString model.QueryString[string]) model.InterfaceResponse[[]entity.Province]
		FindByID(ctx context.Context, id string) model.InterfaceResponse[*entity.Province]
	}

	provinceStruct struct {
		DB  *gorm.DB
		Log *logrus.Logger
	}
)

// FindAll implements ProvinceRepository.
func (repository *provinceStruct) FindAll(ctx context.Context, queryString model.QueryString[string]) model.InterfaceResponse[[]entity.Province] {
	var data []entity.Province
	metaPagination := new(model.MetaPagination)
	paginate := &common.Pagination{
		Limit: queryString.QueryPagination.PerPage,
		Page:  queryString.QueryPagination.Page,
	}
	response := model.InterfaceResponse[[]entity.Province]{
		Status:         common.StatusUnProccessableEntity,
		MetaPagination: metaPagination,
	}

	tx := repository.DB.WithContext(ctx)
	if queryString.Query != "" {
		q := "%" + queryString.Query + "%"
		tx = tx.Where("name LIKE ?", q)
	}

	if err := tx.
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

// FindByID implements ProvinceRepository.
func (repository *provinceStruct) FindByID(ctx context.Context, id string) model.InterfaceResponse[*entity.Province] {
	var data *entity.Province
	response := model.InterfaceResponse[*entity.Province]{
		Status: common.StatusUnProccessableEntity,
	}

	tx := repository.DB.WithContext(ctx)
	if err := tx.
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

func NewProvinceRepository(db *gorm.DB, log *logrus.Logger) ProvinceRepository {
	return &provinceStruct{
		DB:  db,
		Log: log,
	}
}
