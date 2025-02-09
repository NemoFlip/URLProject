package services

import (
	"URLProject/internal/entity"
	"URLProject/internal/errors"
	"URLProject/pkg/di"
	"errors"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	UserRepository di.IUserRepository
}

func NewAuthService(userRepository di.IUserRepository) *AuthService {
	return &AuthService{UserRepository: userRepository}
}

func (as *AuthService) Register(email, password, name string) (string, error) {
	if existedUser, _ := as.UserRepository.FindByEmail(email); existedUser != nil {
		return "", errors.New(customErrors.ErrUserExists)
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	user := &entity.User{
		Name:     name,
		Email:    email,
		Password: string(hashedPassword),
	}
	if err = as.UserRepository.Create(user); err != nil {
		return "", err
	}

	return email, nil
}

func (as *AuthService) Login(email, password string) (string, error) {
	userFromDB, _ := as.UserRepository.FindByEmail(email)
	if userFromDB == nil {
		return "", errors.New(customErrors.ErrWrongCredentials)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(userFromDB.Password), []byte(password)); err != nil {
		return "", errors.New(customErrors.ErrWrongCredentials)
	}
	return email, nil
}
