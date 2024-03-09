package service

import (
	"context"

	"github.com/bagusyanuar/go-internal-yousee/internal/repositories"
)

type (
	AuthService interface {
		SignIn(ctx context.Context) (string, error)
	}

	auth struct {
		AuthRepository repositories.AuthRepository
	}
)

// SignIn implements AuthServiceUsecase.
func (service *auth) SignIn(ctx context.Context) (string, error) {
	user, err := service.AuthRepository.SignIn(ctx)
	if err != nil {
		return "", err
	}
	return user.Username, nil
}

func NewAuthService(authRepository repositories.AuthRepository) AuthService {
	return &auth{
		AuthRepository: authRepository,
	}
}
