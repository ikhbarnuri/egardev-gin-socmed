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
