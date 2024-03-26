package service

import (
	"context"
	"time"

	"github.com/bagusyanuar/go-internal-yousee/common"
	"github.com/bagusyanuar/go-internal-yousee/internal/entity"
	"github.com/bagusyanuar/go-internal-yousee/internal/model"
	"github.com/bagusyanuar/go-internal-yousee/internal/repositories"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
)

type (
	AuthService interface {
		SignIn(ctx context.Context, request *model.AuthRequest) (model.InterfaceResponse[*model.AuthResponse], *entity.User)
		ValidateFormRequest(ctx context.Context, request *model.AuthRequest) model.InterfaceResponse[any]
	}

	authStruct struct {
		AuthRepository repositories.AuthRepository
		JWT            *common.JWT
		Validator      *validator.Validate
	}
)

// SignIn implements AuthServiceUsecase.
func (service *authStruct) SignIn(ctx context.Context, request *model.AuthRequest) (model.InterfaceResponse[*model.AuthResponse], *entity.User) {
	response := model.InterfaceResponse[*model.AuthResponse]{
		Status: common.StatusInternalServerError,
		Error:  common.ErrUnknown,
	}

	username := request.Username

	repositoryResponse := service.AuthRepository.SignIn(ctx, username)
	if repositoryResponse.Error != nil {
		response.Status = repositoryResponse.Status
		response.Error = repositoryResponse.Error
		return response, nil
	}

	user := repositoryResponse.Data
	accessToken, err := service.createToken(service.JWT, user)

	if err != nil {
		response.Error = common.ErrGenerateToken
		return response, nil
	}

	response.Status = common.StatusOK
	response.Error = nil
	response.Data = &model.AuthResponse{
		AccessToken: accessToken,
	}
	return response, user
}

// ValidateFormRequest implements AuthService.
func (service *authStruct) ValidateFormRequest(ctx context.Context, request *model.AuthRequest) model.InterfaceResponse[any] {
	response := model.InterfaceResponse[any]{
		Status: common.StatusInternalServerError,
		Error:  common.ErrValidateRequest,
	}

	err, msg := common.Validate(service.Validator, request)
	if err != nil {
		response.Status = common.StatusBadRequest
		response.Error = err
		response.Data = msg
		return response
	}
	response.Status = common.StatusOK
	response.Error = nil
	return response
}

func (service *authStruct) createToken(cfg *common.JWT, user *entity.User) (string, error) {
	JWTSignInMethod := jwt.SigningMethodHS256
	exp := time.Now().Add(time.Minute * time.Duration(cfg.Exp))
	claims := common.JWTClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    cfg.Issuer,
			ExpiresAt: jwt.NewNumericDate(exp),
		},
		// UserID:   user.ID,
		Username: user.Username,
		Role:     "admin",
	}

	token := jwt.NewWithClaims(JWTSignInMethod, claims)
	return token.SignedString([]byte(cfg.SignatureKey))
}

func NewAuthService(authRepository repositories.AuthRepository, jwt *common.JWT, validator *validator.Validate) AuthService {
	return &authStruct{
		AuthRepository: authRepository,
		JWT:            jwt,
		Validator:      validator,
	}
}
