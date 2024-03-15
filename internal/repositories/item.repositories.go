package repositories

import (
	"context"

	"github.com/bagusyanuar/go-internal-yousee/common"
	"github.com/bagusyanuar/go-internal-yousee/internal/entity"
	"github.com/bagusyanuar/go-internal-yousee/internal/model"
	"github.com/bagusyanuar/go-internal-yousee/internal/model/transformer"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type (
	ItemRepository interface {
		FindAll(ctx context.Context, queryString model.QueryString[string]) (model.Response[[]entity.Item], error)
		// FindByID(ctx context.Context, id string) (*entity.Vendor, error)
		// Create(ctx context.Context, entity *entity.Vendor) error
		// Patch(ctx context.Context, id string, data map[string]interface{}) error
		// Delete(ctx context.Context, id string) error
	}

	itemStruct struct {
		DB  *gorm.DB
		Log *logrus.Logger
	}
)

// FindAll implements ItemRepository.
func (repository *itemStruct) FindAll(ctx context.Context, queryString model.QueryString[string]) (model.Response[[]entity.Item], error) {
	var items []entity.Item
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
	if err := tx.Scopes(common.Paginate(items, paginate, tx)).Find(&items).Error; err != nil {
		return model.Response[[]entity.Item]{}, err
	}

	metaPagination = transformer.ToMetaPagination(paginate)
	return model.Response[[]entity.Item]{Data: items, Meta: metaPagination}, nil
}

func NewItemRepository(db *gorm.DB, log *logrus.Logger) ItemRepository {
	return &itemStruct{
		DB:  db,
		Log: log,
	}
}
