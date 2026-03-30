package service

import (
	"errors"
	"trim/internal/domain"
	"trim/internal/repository"

	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(email, password string) (*domain.User, error) {
	if email == "" || password == "" {
		return nil, errors.New("email and password are required")
	}

	_, err := repository.GetUserByEmail(email)
	if err == nil {
		return nil, errors.New("user already exists")
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &domain.User{
		Email:        email,
		PasswordHash: string(passwordHash),
	}

	if err := repository.CreateUser(user); err != nil {
		return nil, err
	}

	return user, nil
}

func GetUserByID(id uint) (*domain.User, error) {
	return repository.GetUserByID(id)
}

func AuthenticateUser(email, password string) (*domain.User, error) {
	user, err := repository.GetUserByEmail(email)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		return nil, errors.New("invalid credentials")
	}

	return user, nil
}
