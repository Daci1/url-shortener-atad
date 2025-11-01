package service

import (
	"fmt"
	"github.com/Daci1/url-shortener-atad/internal/errs"
	"github.com/Daci1/url-shortener-atad/internal/hash"
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

func (s *UserService) RegisterUser(req response.RegisterRequestAttributes) (*response.APIResponse[response.CredentialsResponseAttributes], errs.CustomError) {
	// TODO: add email validation
	password, err := hash.HashPassword(req.Password)
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

	customError := s.userRepository.RegisterUser(*userEntity)
	if customError != nil {
		return nil, customError
	}

	// TODO: create jwt token and refresh token
	attributes := response.ToCredentialsResponseAttributes(req.Username, "", "")
	return response.New("users", attributes), nil
}

func (s *UserService) LoginUser(req response.LoginRequestAttributes) (*response.APIResponse[response.CredentialsResponseAttributes], errs.CustomError) {
	entity, customError := s.userRepository.GetUserByEmail(req.Email)
	if customError != nil {
		return nil, customError
	}

	if !hash.CheckPassword(entity.PasswordHash, req.Password) {
		// TODO: make this the same as when user not found
		return nil, errs.Unauthorized("invalid credentials")
	}

	// TODO: create jwt token and refresh token
	attributes := response.ToCredentialsResponseAttributes(entity.Username, "", "")
	return response.New("users", attributes), nil
}
