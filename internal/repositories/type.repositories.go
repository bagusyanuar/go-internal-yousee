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
		FindAll(ctx context.Context, queryString model.QueryString[string]) (model.Response[[]entity.Type], error)
		FindByID(ctx context.Context, id string) (*entity.Type, error)
		Create(ctx context.Context, entity *entity.Type) error
		Patch(ctx context.Context, id string, entity *entity.Type) error
		Delete(ctx context.Context, id string) error
	}

	typeStruct struct {
		DB  *gorm.DB
		Log *logrus.Logger
	}
)

// FindAll implements TypeRepository.
func (repository *typeStruct) FindAll(ctx context.Context, queryString model.QueryString[string]) (model.Response[[]entity.Type], error) {
	var types []entity.Type
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
		Scopes(common.Paginate(types, paginate, tx)).
		Find(&types).Error; err != nil {
		return model.Response[[]entity.Type]{}, err
	}
	metaPagination = transformer.ToMetaPagination(paginate)
	return model.Response[[]entity.Type]{Data: types, Meta: metaPagination}, nil
}

// FindByID implements TypeRepository.
func (repository *typeStruct) FindByID(ctx context.Context, id string) (*entity.Type, error) {
	var entity *entity.Type
	tx := repository.DB.WithContext(ctx)
	if err := tx.
		Omit(clause.Associations).
		Where("id = ?", id).
		First(&entity).Error; err != nil {
		if errors.Is(gorm.ErrRecordNotFound, err) {
			return nil, common.ErrRecordNotFound
		}
		return nil, err
	}
	return entity, nil
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

// Create implements TypeRepository.
func (repository *typeStruct) Create(ctx context.Context, entity *entity.Type) error {
	tx := repository.DB.WithContext(ctx)
	if err := tx.Create(entity).Error; err != nil {
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
