package repositories

import (
	"context"

	"github.com/bagusyanuar/go-internal-yousee/common"
	"github.com/bagusyanuar/go-internal-yousee/internal/entity"
	"github.com/bagusyanuar/go-internal-yousee/internal/model"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type (
	ProvinceRepository interface {
		FindAll(ctx context.Context, param string, paginateQuery model.PaginationQuery) ([]entity.Province, error)
	}

	province struct {
		DB  *gorm.DB
		Log *logrus.Logger
	}
)

// FindAll implements ProvinceRepository.
func (repository *province) FindAll(ctx context.Context, param string, paginateQuery model.PaginationQuery) ([]entity.Province, error) {
	var provinces []entity.Province

	tx := repository.DB.WithContext(ctx)

	if param != "" {
		q := "%" + param + "%"
		tx = tx.Where("name LIKE ?", q)
	}

	paginate := common.Pagination{
		Limit: paginateQuery.PerPage,
		Page:  paginateQuery.Page,
	}
	if err := tx.Scopes(common.Paginate(provinces, &paginate, tx)).Find(&provinces).Error; err != nil {
		return nil, err
	}
	repository.Log.Warnf("Total Rows : %+v", paginate.TotalRows)
	return provinces, nil
}

func NewProvinceRepository(db *gorm.DB, log *logrus.Logger) ProvinceRepository {
	return &province{
		DB:  db,
		Log: log,
	}
}
