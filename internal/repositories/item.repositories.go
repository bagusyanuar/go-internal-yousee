package repositories

import (
	"context"

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
		FindAll(ctx context.Context, queryString model.QueryString[string]) model.InterfaceResponse[[]entity.Item]
		FindByID(ctx context.Context, id string) (*entity.Item, error)
		Create(ctx context.Context, entity *entity.Item) error
		// Patch(ctx context.Context, id string, data map[string]interface{}) error
		// Delete(ctx context.Context, id string) error
	}

	itemStruct struct {
		DB  *gorm.DB
		Log *logrus.Logger
	}
)

// FindAll implements ItemRepository.
func (repository *itemStruct) FindAll(ctx context.Context, queryString model.QueryString[string]) model.InterfaceResponse[[]entity.Item] {
	var data []entity.Item
	metaPagination := new(model.MetaPagination)
	paginate := &common.Pagination{
		Limit: queryString.QueryPagination.PerPage,
		Page:  queryString.QueryPagination.Page,
	}
	response := model.InterfaceResponse[[]entity.Item]{
		Status:         common.StatusUnProccessableEntity,
		MetaPagination: metaPagination,
	}
	tx := repository.DB.WithContext(ctx)

	if queryString.Query != "" {
		q := "%" + queryString.Query + "%"
		tx = tx.Where("name LIKE ?", q)
	}

	if err := tx.
		Preload("Type").
		Preload("City").
		Preload("Vendor").
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
func (repository *itemStruct) FindByID(ctx context.Context, id string) (*entity.Item, error) {
	var entity *entity.Item
	tx := repository.DB.WithContext(ctx)
	if err := tx.
		Preload("Type").
		Preload("City").
		Preload("Vendor").
		Where("id = ?", id).
		First(&entity).Error; err != nil {
		return entity, err
	}
	return entity, nil
}

// Create implements ItemRepository.
func (repository *itemStruct) Create(ctx context.Context, entity *entity.Item) error {
	tx := repository.DB.WithContext(ctx)
	if err := tx.Omit(clause.Associations).Create(entity).Error; err != nil {
		return err
	}
	return nil
}

func NewItemRepository(db *gorm.DB, log *logrus.Logger) ItemRepository {
	return &itemStruct{
		DB:  db,
		Log: log,
	}
}
