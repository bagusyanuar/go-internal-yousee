package repositories

import (
	"context"

	"github.com/bagusyanuar/go-internal-yousee/internal/entity"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type (
	AuthRepository interface {
		SignIn(ctx context.Context) (*entity.User, error)
	}

	auth struct {
		DB  *gorm.DB
		Log *logrus.Logger
	}
)

// SignIn implements AuthRepositoryUsecase.
func (repository *auth) SignIn(ctx context.Context) (*entity.User, error) {
	tx := repository.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	user := new(entity.User)
	username := "admin1"
	if err := tx.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func NewAuthRepository(db *gorm.DB, log *logrus.Logger) AuthRepository {
	return &auth{
		DB:  db,
		Log: log,
	}
}
