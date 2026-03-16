package repository

import (
	"trim/internal/config"
	"trim/internal/domain"
)

func GetAll() ([]domain.Transaction, error) {

	var transactions []domain.Transaction

	result := config.DB.Find(&transactions)

	return transactions, result.Error
}

func Create(t *domain.Transaction) error {

	result := config.DB.Create(t)

	return result.Error
}