package repositories

import (
	"context"
	"errors"

	"github.com/bagusyanuar/go-internal-yousee/common"
	"github.com/bagusyanuar/go-internal-yousee/internal/entity"
	"github.com/bagusyanuar/go-internal-yousee/internal/model"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type (
	AuthRepository interface {
		SignIn(ctx context.Context, username string) model.InterfaceResponse[*entity.User]
	}

	authStuct struct {
		DB  *gorm.DB
		Log *logrus.Logger
	}
)

// SignIn implements AuthRepositoryUsecase.
func (repository *authStuct) SignIn(ctx context.Context, username string) model.InterfaceResponse[*entity.User] {
	var data *entity.User
	response := model.InterfaceResponse[*entity.User]{
		Status: common.StatusUnProccessableEntity,
	}
	tx := repository.DB.WithContext(ctx).Begin()

	if err := tx.
		Where("username = ?", username).
		First(&data).Error; err != nil {
		if errors.Is(gorm.ErrRecordNotFound, err) {
			repository.Log.Warnf("query failed : %+v", err)
			response.Status = common.StatusNotFound
			response.Error = common.ErrUserNotFound
			return response
		}
		response.Error = err
		return response
	}
	response.Status = common.StatusOK
	response.Data = data
	return response
}

func NewAuthRepository(db *gorm.DB, log *logrus.Logger) AuthRepository {
	return &authStuct{
		DB:  db,
		Log: log,
	}
}
