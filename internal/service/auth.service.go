package service

import (
	"context"
	"errors"

	"github.com/bagusyanuar/go-internal-yousee/internal/model"
	"github.com/bagusyanuar/go-internal-yousee/internal/repositories"
	"gorm.io/gorm"
)

type (
	AuthService interface {
		SignIn(ctx context.Context, request *model.AuthRequest) (*model.AuthResponse, error)
	}

	auth struct {
		AuthRepository repositories.AuthRepository
	}
)

// SignIn implements AuthServiceUsecase.
func (service *auth) SignIn(ctx context.Context, request *model.AuthRequest) (*model.AuthResponse, error) {
	username := request.Username
	response := new(model.AuthResponse)
	user, err := service.AuthRepository.SignIn(ctx, username)
	if err != nil {
		if errors.Is(gorm.ErrRecordNotFound, err) {
			return response, errors.New("user not found")
		}
		return response, err
	}
	response.AccessToken = user.Username
	return response, nil
}

func NewAuthService(authRepository repositories.AuthRepository) AuthService {
	return &auth{
		AuthRepository: authRepository,
	}
}
