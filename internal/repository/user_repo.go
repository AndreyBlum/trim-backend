package repository

import (
	"trim/internal/config"
	"trim/internal/domain"
)

func CreateUser(u *domain.User) error {
	result := config.DB.Create(u)
	return result.Error
}

func GetUserByEmail(email string) (*domain.User, error) {
	var user domain.User
	result := config.DB.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func GetUserByID(id uint) (*domain.User, error) {
	var user domain.User
	result := config.DB.First(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
