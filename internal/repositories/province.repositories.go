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
	ProvinceRepository interface {
		FindAll(ctx context.Context, queryString model.QueryString[string]) (model.Response[[]entity.Province], error)
	}

	province struct {
		DB  *gorm.DB
		Log *logrus.Logger
	}
)

// FindAll implements ProvinceRepository.
func (repository *province) FindAll(ctx context.Context, queryString model.QueryString[string]) (model.Response[[]entity.Province], error) {
	var provinces []entity.Province

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
	if err := tx.Scopes(common.Paginate(provinces, paginate, tx)).Find(&provinces).Error; err != nil {
		return model.Response[[]entity.Province]{}, err
	}

	metaPagination = transformer.ToMetaPagination(paginate)
	return model.Response[[]entity.Province]{Data: provinces, Meta: metaPagination}, nil
}

func NewProvinceRepository(db *gorm.DB, log *logrus.Logger) ProvinceRepository {
	return &province{
		DB:  db,
		Log: log,
	}
}
