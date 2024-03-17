package service

import (
	"context"
	"errors"
	"time"

	"github.com/bagusyanuar/go-internal-yousee/common"
	"github.com/bagusyanuar/go-internal-yousee/internal/entity"
	"github.com/bagusyanuar/go-internal-yousee/internal/model"
	"github.com/bagusyanuar/go-internal-yousee/internal/repositories"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

type (
	AuthService interface {
		SignIn(ctx context.Context, request *model.AuthRequest) (*model.AuthResponse, error)
	}

	auth struct {
		AuthRepository repositories.AuthRepository
		JWT            *common.JWT
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

	accessToken, err := service.createToken(service.JWT, user)

	if err != nil {
		return response, errors.New("failed to generate access token")
	}
	response.AccessToken = accessToken
	return response, nil
}

func (service *auth) createToken(cfg *common.JWT, user *entity.User) (string, error) {
	JWTSignInMethod := jwt.SigningMethodHS256
	exp := time.Now().Add(time.Minute * time.Duration(cfg.Exp))
	claims := common.JWTClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    cfg.Issuer,
			ExpiresAt: jwt.NewNumericDate(exp),
		},
		UserID: user.ID,
	}

	token := jwt.NewWithClaims(JWTSignInMethod, claims)
	return token.SignedString([]byte(cfg.SignatureKey))
}

func NewAuthService(authRepository repositories.AuthRepository, jwt *common.JWT) AuthService {
	return &auth{
		AuthRepository: authRepository,
		JWT:            jwt,
	}
}
