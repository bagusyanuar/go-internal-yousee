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
	"gorm.io/gorm/clause"
)

type (
	ItemRepository interface {
		FindAll(ctx context.Context, queryString model.QueryString[model.ItemQueryString]) model.InterfaceResponse[[]entity.Item]
		FindByID(ctx context.Context, id string) model.InterfaceResponse[*entity.Item]
		Create(ctx context.Context, entry *entity.Item) model.InterfaceResponse[*entity.Item]
		Patch(ctx context.Context, id string, entry map[string]interface{}) model.InterfaceResponse[*entity.Item]
		// Delete(ctx context.Context, id string) error
	}

	itemStruct struct {
		DB  *gorm.DB
		Log *logrus.Logger
	}
)

// FindAll implements ItemRepository.
func (repository *itemStruct) FindAll(ctx context.Context, queryString model.QueryString[model.ItemQueryString]) model.InterfaceResponse[[]entity.Item] {
	var data []entity.Item
	queryParam := queryString.Query.Param
	queryType := queryString.Query.TypeID
	queryCity := queryString.Query.CityID
	queryVendor := queryString.Query.VendorID
	metaPagination := new(model.MetaPagination)
	paginate := &common.Pagination{
		Limit: queryString.QueryPagination.PerPage,
		Page:  queryString.QueryPagination.Page,
	}
	response := model.InterfaceResponse[[]entity.Item]{
		Status:         common.StatusUnProccessableEntity,
		MetaPagination: metaPagination,
	}
	tx := repository.DB.WithContext(ctx).
		Preload("Type").
		Preload("City").
		Preload("Vendor")

	if queryParam != "" {
		q := "%" + queryParam + "%"
		tx = tx.
			Where("name LIKE ?", q).
			Or("address LIKE ?", q).
			Or("location LIKE ?", q)
	}

	if queryCity != "" {
		tx = tx.Where("city_id = ?", queryCity)
	}

	if queryType != "" {
		tx = tx.Where("type_id = ?", queryType)
	}

	if queryVendor != "" {
		tx = tx.Where("vendor_id = ?", queryVendor)
	}

	if err := tx.
		Scopes(common.Paginate(data, paginate, tx)).
		Find(&data).Error; err != nil {
		repository.Log.Warnf("query failed : %+v", err)
		response.Error = err
		return response
	}

	response.Status = common.StatusOK
	response.Data = data
	response.MetaPagination = transformer.ToMetaPagination(paginate)
	return response
}

// FindByID implements ItemRepository.
func (repository *itemStruct) FindByID(ctx context.Context, id string) model.InterfaceResponse[*entity.Item] {
	var data *entity.Item
	response := model.InterfaceResponse[*entity.Item]{
		Status: common.StatusUnProccessableEntity,
	}

	tx := repository.DB.WithContext(ctx)
	if err := tx.
		Preload("Type").
		Preload("City").
		Preload("Vendor").
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

// Create implements ItemRepository.
func (repository *itemStruct) Create(ctx context.Context, entry *entity.Item) model.InterfaceResponse[*entity.Item] {
	response := model.InterfaceResponse[*entity.Item]{
		Status: common.StatusUnProccessableEntity,
	}

	tx := repository.DB.WithContext(ctx)
	if err := tx.Omit(clause.Associations).Create(&entry).Error; err != nil {
		repository.Log.Warnf("query failed : %+v", err)
		response.Error = err
		return response
	}
	response.Status = common.StatusCreated
	return response
}

// Patch implements ItemRepository.
func (repository *itemStruct) Patch(ctx context.Context, id string, entry map[string]interface{}) model.InterfaceResponse[*entity.Item] {
	response := model.InterfaceResponse[*entity.Item]{
		Status: common.StatusUnProccessableEntity,
	}
	tx := repository.DB.WithContext(ctx)

	data := new(entity.Item)
	if err := tx.
		Where("id = ?", id).
		First(&data).Error; err != nil {
		repository.Log.Warnf("query failed : %+v", err)
		if errors.Is(gorm.ErrRecordNotFound, err) {
			response.Error = err
			response.Status = common.StatusNotFound
			return response
		}
		response.Error = err
		return response
	}

	if err := tx.Model(&data).
		Omit(clause.Associations).
		Where("id = ?", id).
		Updates(&entry).Error; err != nil {
		response.Error = err
		return response
	}
	response.Status = common.StatusOK
	return response
}

func NewItemRepository(db *gorm.DB, log *logrus.Logger) ItemRepository {
	return &itemStruct{
		DB:  db,
		Log: log,
	}
}
