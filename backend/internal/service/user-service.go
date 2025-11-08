package service

import (
	"fmt"

	"github.com/Daci1/url-shortener-atad/internal/errs"
	"github.com/Daci1/url-shortener-atad/internal/security"
	"github.com/Daci1/url-shortener-atad/internal/types"

	"time"

	"github.com/Daci1/url-shortener-atad/internal/db"
	"github.com/Daci1/url-shortener-atad/internal/server/response"
	"github.com/google/uuid"
)

type UserService struct {
	userRepository *db.UserRepository
}

func NewUserService(userRepo *db.UserRepository) *UserService {
	return &UserService{
		userRepository: userRepo,
	}
}

func (s *UserService) RegisterUser(req types.RegisterRequestAttributes) (*types.APIResponse[types.CredentialsResponseAttributes], errs.CustomError) {
	// TODO: add email validation
	password, err := security.HashPassword(req.Password)
	if err != nil {
		return nil, errs.Internal(fmt.Sprintf("Error while encrypting password: %s", err.Error()))
	}
	userEntity := &db.UserEntity{
		Id:           uuid.NewString(),
		Email:        req.Email,
		Username:     req.Username,
		PasswordHash: password,
		CreatedAt:    time.Now(),
	}

	entity, customError := s.userRepository.RegisterUser(userEntity)
	if customError != nil {
		return nil, customError
	}

	tokenPairs, err := security.GenerateTokens(entity.Id, entity.Username)
	if err != nil {
		return nil, errs.Internal(fmt.Sprintf("Error while creating token pairs: %s", tokenPairs))
	}

	attributes := response.ToCredentialsResponseAttributes(req.Username, tokenPairs.Token, tokenPairs.RefreshToken)
	return response.New("users", attributes), nil
}

func (s *UserService) LoginUser(req types.LoginRequestAttributes) (*types.APIResponse[types.CredentialsResponseAttributes], errs.CustomError) {
	entity, customError := s.userRepository.GetUserByEmail(req.Email)
	if customError != nil {
		return nil, customError
	}

	if !security.CheckPassword(entity.PasswordHash, req.Password) {
		// TODO: make this the same as when user not found
		return nil, errs.Unauthorized("invalid credentials")
	}

	tokenPairs, err := security.GenerateTokens(entity.Id, entity.Username)
	if err != nil {
		return nil, errs.Internal(fmt.Sprintf("Error while creating token pairs: %s", tokenPairs))
	}

	attributes := response.ToCredentialsResponseAttributes(entity.Username, tokenPairs.Token, tokenPairs.RefreshToken)
	return response.New("users", attributes), nil
}
