package repositories

import (
	"context"

	"github.com/bagusyanuar/go-internal-yousee/internal/entity"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type (
	TypeRepository interface {
		FindAll(ctx context.Context, param string) ([]entity.Type, error)
		FindByID(ctx context.Context, id string) (*entity.Type, error)
		Create(ctx context.Context, entity *entity.Type) error
		Patch(ctx context.Context, id string, entity *entity.Type) error
		Delete(ctx context.Context, id string) error
	}

	itemType struct {
		DB  *gorm.DB
		Log *logrus.Logger
	}
)

// Delete implements TypeRepository.
func (repository *itemType) Delete(ctx context.Context, id string) error {
	entity := new(entity.Type)
	tx := repository.DB.WithContext(ctx)
	if err := tx.Omit(clause.Associations).Where("id = ?", id).Unscoped().Delete(&entity).Error; err != nil {
		return err
	}
	return nil
}

// FindByID implements TypeRepository.
func (repository *itemType) FindByID(ctx context.Context, id string) (*entity.Type, error) {
	entity := new(entity.Type)
	tx := repository.DB.WithContext(ctx)
	if err := tx.Omit(clause.Associations).Where("id = ?", id).First(&entity).Error; err != nil {
		return nil, err
	}
	return entity, nil
}

// Patch implements TypeRepository.
func (repository *itemType) Patch(ctx context.Context, id string, entity *entity.Type) error {
	tx := repository.DB.WithContext(ctx)
	if err := tx.Omit(clause.Associations).Where("id = ?", id).Updates(&entity).Error; err != nil {
		return err
	}
	return nil
}

// Create implements TypeRepository.
func (repository *itemType) Create(ctx context.Context, entity *entity.Type) error {
	tx := repository.DB.WithContext(ctx)
	if err := tx.Create(entity).Error; err != nil {
		return err
	}
	return nil
}

// FindAll implements TypeRepository.
func (repository *itemType) FindAll(ctx context.Context, param string) ([]entity.Type, error) {
	var types []entity.Type

	tx := repository.DB.WithContext(ctx)

	if param != "" {
		q := "%" + param + "%"
		tx = tx.Where("name LIKE ?", q)
	}
	if err := tx.Find(&types).Error; err != nil {
		return nil, err
	}
	return types, nil
}

func NewTypeRepository(db *gorm.DB, log *logrus.Logger) TypeRepository {
	return &itemType{
		DB:  db,
		Log: log,
	}
}
