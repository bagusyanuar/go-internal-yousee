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
	VendorRepository interface {
		FindAll(ctx context.Context, queryString model.QueryString[string]) (model.Response[[]entity.Vendor], error)
	}

	vendor struct {
		DB  *gorm.DB
		Log *logrus.Logger
	}
)

// FindAll implements VendorRepository.
func (repository *vendor) FindAll(ctx context.Context, queryString model.QueryString[string]) (model.Response[[]entity.Vendor], error) {
	var vendors []entity.Vendor

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
	if err := tx.Scopes(common.Paginate(vendors, paginate, tx)).Find(&vendors).Error; err != nil {
		return model.Response[[]entity.Vendor]{}, err
	}

	metaPagination = transformer.ToMetaPagination(paginate)
	return model.Response[[]entity.Vendor]{Data: vendors, Meta: metaPagination}, nil
}

func NewVendorRepository(db *gorm.DB, log *logrus.Logger) VendorRepository {
	return &vendor{
		DB:  db,
		Log: log,
	}
}
