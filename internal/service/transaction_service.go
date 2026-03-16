package service

import (
	"trim/internal/domain"
	"trim/internal/repository"
)

func ListTransactions() ([]domain.Transaction, error) {

	return repository.GetAll()
}

func AddTransaction(t *domain.Transaction) error {

	return repository.Create(t)
}