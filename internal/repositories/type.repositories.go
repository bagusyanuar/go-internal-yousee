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
		FindAll(ctx context.Context, queryString model.QueryString[string]) (model.Response[[]entity.Type], int, error)
		FindByID(ctx context.Context, id string) (*entity.Type, int, error)
		Create(ctx context.Context, entity *entity.Type) (int, error)
		Patch(ctx context.Context, id string, entity *entity.Type) error
		Delete(ctx context.Context, id string) error
	}

	typeStruct struct {
		DB  *gorm.DB
		Log *logrus.Logger
	}
)

// FindAll implements TypeRepository.
func (repository *typeStruct) FindAll(ctx context.Context, queryString model.QueryString[string]) (model.Response[[]entity.Type], int, error) {
	var data []entity.Type
	code := common.StatusUnProccessableEntity
	metaPagination := new(model.MetaPagination)

	tx := repository.DB.WithContext(ctx)

	if queryString.Query != "" {
		q := "%" + queryString.Query + "%"
		tx = tx.Where("name LIKE ?", q)
	}

	paginate := &common.Pagination{
		Limit: queryString.QueryPagination.PerPage,
		Page:  queryString.QueryPagination.Page,
	}

	if err := tx.
		Scopes(common.Paginate(data, paginate, tx)).
		Preload("Abc").
		Find(&data).Error; err != nil {
		repository.Log.Warnf("query failed : %+v", err)
		return model.Response[[]entity.Type]{}, code, err
	}
	code = common.StatusOK
	metaPagination = transformer.ToMetaPagination(paginate)
	return model.Response[[]entity.Type]{
		Data: data,
		Meta: metaPagination,
	}, code, nil
}

// FindByID implements TypeRepository.
func (repository *typeStruct) FindByID(ctx context.Context, id string) (*entity.Type, int, error) {
	var data *entity.Type
	code := common.StatusUnProccessableEntity
	tx := repository.DB.WithContext(ctx)
	if err := tx.
		Omit(clause.Associations).
		Preload("ASD").
		Where("id = ?", id).
		First(&data).Error; err != nil {
		if errors.Is(gorm.ErrRecordNotFound, err) {
			repository.Log.Warnf("query failed : %+v", err)
			code = common.StatusNotFound
			return nil, code, common.ErrRecordNotFound
		}
		return nil, code, err
	}
	code = common.StatusOK
	return data, code, nil
}

// Create implements TypeRepository.
func (repository *typeStruct) Create(ctx context.Context, entity *entity.Type) (int, error) {
	code := common.StatusUnProccessableEntity
	tx := repository.DB.WithContext(ctx)
	if err := tx.Create(entity).Error; err != nil {
		repository.Log.Warnf("query failed : %+v", err)
		return code, err
	}
	code = common.StatusOK
	return code, nil
}

// Delete implements TypeRepository.
func (repository *typeStruct) Delete(ctx context.Context, id string) error {
	entity := new(entity.Type)
	tx := repository.DB.WithContext(ctx)
	if err := tx.Omit(clause.Associations).Where("id = ?", id).Unscoped().Delete(&entity).Error; err != nil {
		return err
	}
	return nil
}

// Patch implements TypeRepository.
func (repository *typeStruct) Patch(ctx context.Context, id string, entity *entity.Type) error {
	tx := repository.DB.WithContext(ctx)
	if err := tx.Omit(clause.Associations).Where("id = ?", id).Updates(&entity).Error; err != nil {
		return err
	}
	return nil
}

func NewTypeRepository(db *gorm.DB, log *logrus.Logger) TypeRepository {
	return &typeStruct{
		DB:  db,
		Log: log,
	}
}
