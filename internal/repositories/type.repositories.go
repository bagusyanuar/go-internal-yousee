package repositories

import (
	"context"

	"github.com/bagusyanuar/go-internal-yousee/internal/entity"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type (
	TypeRepository interface {
		FindAll(ctx context.Context) ([]entity.Type, error)
		Create(ctx context.Context, entity *entity.Type) error
	}

	itemType struct {
		DB  *gorm.DB
		Log *logrus.Logger
	}
)

// Create implements TypeRepository.
func (repository *itemType) Create(ctx context.Context, entity *entity.Type) error {
	tx := repository.DB.WithContext(ctx)
	if err := tx.Create(entity).Error; err != nil {
		return err
	}
	return nil
}

// FindAll implements TypeRepository.
func (repository *itemType) FindAll(ctx context.Context) ([]entity.Type, error) {
	var types []entity.Type

	tx := repository.DB.WithContext(ctx)
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
