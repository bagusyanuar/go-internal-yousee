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
	CityRepository interface {
		FindAll(ctx context.Context, queryString model.QueryString[string]) (model.Response[[]entity.City], error)
		FindByID(ctx context.Context, id string) (*entity.City, error)
	}

	city struct {
		DB  *gorm.DB
		Log *logrus.Logger
	}
)

// FindAll implements CityRepository.
func (repository *city) FindAll(ctx context.Context, queryString model.QueryString[string]) (model.Response[[]entity.City], error) {
	var cities []entity.City

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
	if err := tx.Preload("Province").Scopes(common.Paginate(cities, paginate, tx)).Find(&cities).Error; err != nil {
		return model.Response[[]entity.City]{}, err
	}

	metaPagination = transformer.ToMetaPagination(paginate)
	return model.Response[[]entity.City]{Data: cities, Meta: metaPagination}, nil
}

// FindByID implements CityRepository.
func (repository *city) FindByID(ctx context.Context, id string) (*entity.City, error) {
	entity := new(entity.City)
	tx := repository.DB.WithContext(ctx)
	if err := tx.Preload("Province").Find(&entity).Error; err != nil {
		return entity, err
	}
	return entity, nil
}

func NewCityRepository(db *gorm.DB, log *logrus.Logger) CityRepository {
	return &city{
		DB:  db,
		Log: log,
	}
}
