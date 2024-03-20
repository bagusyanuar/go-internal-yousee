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
	VendorRepository interface {
		FindAll(ctx context.Context, queryString model.QueryString[string]) model.InterfaceResponse[[]entity.Vendor]
		FindByID(ctx context.Context, id string) model.InterfaceResponse[*entity.Vendor]
		Create(ctx context.Context, entry *entity.Vendor) model.InterfaceResponse[*entity.Vendor]
		Patch(ctx context.Context, id string, entry map[string]interface{}) model.InterfaceResponse[*entity.Vendor]
		Delete(ctx context.Context, id string) model.InterfaceResponse[any]
	}

	vendorStruct struct {
		DB  *gorm.DB
		Log *logrus.Logger
	}
)

// FindAll implements VendorRepository.
func (repository *vendorStruct) FindAll(ctx context.Context, queryString model.QueryString[string]) model.InterfaceResponse[[]entity.Vendor] {
	var data []entity.Vendor
	metaPagination := new(model.MetaPagination)
	paginate := &common.Pagination{
		Limit: queryString.QueryPagination.PerPage,
		Page:  queryString.QueryPagination.Page,
	}
	response := model.InterfaceResponse[[]entity.Vendor]{
		Status:         common.StatusUnProccessableEntity,
		MetaPagination: metaPagination,
	}

	tx := repository.DB.WithContext(ctx)

	if queryString.Query != "" {
		q := "%" + queryString.Query + "%"
		tx = tx.Where("name LIKE ?", q).
			Or("address LIKE ?", q).
			Or("brand LIKE ?", q)
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

// FindByID implements VendorRepository.
func (repository *vendorStruct) FindByID(ctx context.Context, id string) model.InterfaceResponse[*entity.Vendor] {
	var data *entity.Vendor
	response := model.InterfaceResponse[*entity.Vendor]{
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

// Create implements VendorRepository.
func (repository *vendorStruct) Create(ctx context.Context, entry *entity.Vendor) model.InterfaceResponse[*entity.Vendor] {
	response := model.InterfaceResponse[*entity.Vendor]{
		Status: common.StatusUnProccessableEntity,
	}

	tx := repository.DB.WithContext(ctx)
	if err := tx.Create(&entry).Error; err != nil {
		repository.Log.Warnf("query failed : %+v", err)
		response.Error = err
		return response
	}
	response.Status = common.StatusCreated
	return response
}

// Patch implements VendorRepository.
func (repository *vendorStruct) Patch(ctx context.Context, id string, entry map[string]interface{}) model.InterfaceResponse[*entity.Vendor] {
	response := model.InterfaceResponse[*entity.Vendor]{
		Status: common.StatusUnProccessableEntity,
	}

	tx := repository.DB.WithContext(ctx)

	data := new(entity.Vendor)
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

// Delete implements VendorRepository.
func (repository *vendorStruct) Delete(ctx context.Context, id string) model.InterfaceResponse[any] {
	var data *entity.Vendor
	response := model.InterfaceResponse[any]{
		Status: common.StatusUnProccessableEntity,
	}
	tx := repository.DB.WithContext(ctx)
	if err := tx.
		Where("id = ?", id).
		First(&data).Error; err != nil {
		if errors.Is(gorm.ErrRecordNotFound, err) {
			repository.Log.Warnf("query failed : %+v", err)
			response.Error = err
			response.Status = common.StatusNotFound
			return response
		}
		response.Error = err
		return response
	}

	if err := tx.Omit(clause.Associations).
		Where("id = ?", id).
		Unscoped().
		Delete(&data).Error; err != nil {
		response.Error = err
		return response
	}
	response.Status = common.StatusOK
	return response
}
func NewVendorRepository(db *gorm.DB, log *logrus.Logger) VendorRepository {
	return &vendorStruct{
		DB:  db,
		Log: log,
	}
}
