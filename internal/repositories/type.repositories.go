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
	TypeRepository interface {
		FindAll(ctx context.Context, queryString model.QueryString[string]) model.InterfaceResponse[[]entity.Type]
		FindByID(ctx context.Context, id string) model.InterfaceResponse[*entity.Type]
		Create(ctx context.Context, data *entity.Type) model.InterfaceResponse[*entity.Type]
		Patch(ctx context.Context, id string, entry map[string]interface{}) model.InterfaceResponse[*entity.Type]
		Delete(ctx context.Context, id string) model.InterfaceResponse[any]
	}

	typeStruct struct {
		DB  *gorm.DB
		Log *logrus.Logger
	}
)

// FindAll implements TypeRepository.
func (repository *typeStruct) FindAll(ctx context.Context, queryString model.QueryString[string]) model.InterfaceResponse[[]entity.Type] {
	var data []entity.Type
	metaPagination := new(model.MetaPagination)
	paginate := &common.Pagination{
		Limit: queryString.QueryPagination.PerPage,
		Page:  queryString.QueryPagination.Page,
	}
	response := model.InterfaceResponse[[]entity.Type]{
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
		repository.Log.Warnf("query failed : %+v", err)
		response.Error = err
		return response
	}
	response.Status = common.StatusOK
	response.Data = data
	response.MetaPagination = transformer.ToMetaPagination(paginate)
	return response
}

// FindByID implements TypeRepository.
func (repository *typeStruct) FindByID(ctx context.Context, id string) model.InterfaceResponse[*entity.Type] {
	var data *entity.Type
	response := model.InterfaceResponse[*entity.Type]{
		Status: common.StatusUnProccessableEntity,
	}

	tx := repository.DB.WithContext(ctx)
	if err := tx.
		Where("id = ?", id).
		First(&data).Error; err != nil {
		if errors.Is(gorm.ErrRecordNotFound, err) {
			repository.Log.Warnf("query failed : %+v", err)
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

// Create implements TypeRepository.
func (repository *typeStruct) Create(ctx context.Context, data *entity.Type) model.InterfaceResponse[*entity.Type] {
	response := model.InterfaceResponse[*entity.Type]{
		Status: common.StatusUnProccessableEntity,
	}

	tx := repository.DB.WithContext(ctx)
	if err := tx.Create(&data).Error; err != nil {
		repository.Log.Warnf("query failed : %+v", err)
		response.Error = err
		return response
	}

	response.Status = common.StatusCreated
	return response
}

// Patch implements TypeRepository.
func (repository *typeStruct) Patch(ctx context.Context, id string, entry map[string]interface{}) model.InterfaceResponse[*entity.Type] {
	response := model.InterfaceResponse[*entity.Type]{
		Status: common.StatusUnProccessableEntity,
	}
	tx := repository.DB.WithContext(ctx)

	data := new(entity.Type)
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

// Delete implements TypeRepository.
func (repository *typeStruct) Delete(ctx context.Context, id string) model.InterfaceResponse[any] {
	var data *entity.Type
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

	if err := tx.
		Omit(clause.Associations).
		Unscoped().
		Delete(&data).Error; err != nil {
		response.Error = err
		return response
	}
	response.Status = common.StatusOK
	return response
}

func NewTypeRepository(db *gorm.DB, log *logrus.Logger) TypeRepository {
	return &typeStruct{
		DB:  db,
		Log: log,
	}
}
