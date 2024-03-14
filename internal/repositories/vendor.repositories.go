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
	VendorRepository interface {
		FindAll(ctx context.Context, queryString model.QueryString[string]) (model.Response[[]entity.Vendor], error)
		FindByID(ctx context.Context, id string) (*entity.Vendor, error)
		Create(ctx context.Context, entity *entity.Vendor) error
		Patch(ctx context.Context, id string, data *entity.Vendor) error
		Delete(ctx context.Context, id string) error
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

// FindByID implements VendorRepository.
func (repository *vendor) FindByID(ctx context.Context, id string) (*entity.Vendor, error) {
	entity := new(entity.Vendor)
	tx := repository.DB.WithContext(ctx)
	if err := tx.Where("id = ?", id).First(&entity).Error; err != nil {
		return entity, err
	}
	return entity, nil
}

// Create implements VendorRepository.
func (repository *vendor) Create(ctx context.Context, entity *entity.Vendor) error {
	tx := repository.DB.WithContext(ctx)
	if err := tx.Create(entity).Error; err != nil {
		return err
	}
	return nil
}

// Patch implements VendorRepository.
func (repository *vendor) Patch(ctx context.Context, id string, data *entity.Vendor) error {
	tx := repository.DB.WithContext(ctx)

	v := new(entity.Vendor)
	if err := tx.Where("id = ?", id).First(&v).Error; err != nil {
		return err
	}

	dataMap := map[string]interface{}{
		"name":     data.Name,
		"phone":    data.Phone,
		"brand":    data.Brand,
		"email":    data.Email,
		"picName":  data.PICName,
		"picPhone": data.PICPhone,
	}
	if err := tx.Model(&v).
		Omit(clause.Associations).
		Where("id = ?", id).
		Updates(&dataMap).Error; err != nil {
		return err
	}
	return nil
}

// Delete implements VendorRepository.
func (repository *vendor) Delete(ctx context.Context, id string) error {
	entity := new(entity.Type)
	tx := repository.DB.WithContext(ctx)
	if err := tx.Omit(clause.Associations).Where("id = ?", id).Unscoped().Delete(&entity).Error; err != nil {
		return err
	}
	return nil
}
func NewVendorRepository(db *gorm.DB, log *logrus.Logger) VendorRepository {
	return &vendor{
		DB:  db,
		Log: log,
	}
}
