package service

import (
	"egardev-gin-socmed/dto"
	"egardev-gin-socmed/entity"
	"egardev-gin-socmed/errorHandler"
	"egardev-gin-socmed/helper"
	"egardev-gin-socmed/repository"
	"time"
)

type AuthService interface {
	Register(req *dto.RegisterRequest) error
	Login(req *dto.LoginRequest) (*dto.LoginResponse, error)
}

type authService struct {
	repository repository.AuthRepository
}

func NewAuthService(r repository.AuthRepository) *authService {
	return &authService{
		repository: r,
	}
}

func (s *authService) Register(req *dto.RegisterRequest) error {
	if emailExist := s.repository.EmailExist(req.Email); emailExist {
		return &errorHandler.BadRequestError{Message: "email already registered"}
	}

	if req.Password != req.PasswordConfirmation {
		return &errorHandler.BadRequestError{Message: "password not match"}
	}

	passwordHash, err := helper.HashPassword(req.Password)
	if err != nil {
		return &errorHandler.InternalServerError{Message: err.Error()}
	}

	user := entity.User{
		Name:      req.Name,
		Email:     req.Email,
		Password:  passwordHash,
		Gender:    req.Gender,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}

	if err := s.repository.Register(&user); err != nil {
		return &errorHandler.InternalServerError{Message: err.Error()}
	}

	return nil
}

func (s authService) Login(req *dto.LoginRequest) (*dto.LoginResponse, error) {
	var response dto.LoginResponse

	user, err := s.repository.GetUseByEmail(req.Email)
	if err != nil {
		return nil, &errorHandler.NotFoundError{Message: "wrong email or password"}
	}

	if err := helper.VerifyPassword(user.Password, req.Password); err != nil {
		return nil, &errorHandler.NotFoundError{Message: "wring email or password"}
	}

	token, err := helper.GenerateToken(user)
	if err != nil {
		return nil, &errorHandler.InternalServerError{Message: err.Error()}
	}

	response = dto.LoginResponse{
		ID:    user.ID,
		Name:  user.Name,
		Token: token,
	}

	return &response, nil
}
