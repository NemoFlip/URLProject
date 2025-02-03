package services

import (
	"URLProject/internal/entity"
	"URLProject/internal/errors"
	"URLProject/internal/repository"
	"errors"
)

type AuthService struct {
	UserRepository *repository.UserRepository
}

func NewAuthService(userRepository *repository.UserRepository) *AuthService {
	return &AuthService{UserRepository: userRepository}
}

func (as *AuthService) Register(email, password, name string) (string, error) {
	if existedUser, _ := as.UserRepository.FindByEmail(email); existedUser != nil {
		return "", errors.New(customErrors.ErrUserExists)
	}

	user := &entity.User{
		Name:     name,
		Email:    email,
		Password: "",
	}
	if err := as.UserRepository.Create(user); err != nil {
		return "", err
	}

	return email, nil
}
