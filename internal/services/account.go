package services

import (
	"oanda-mock-api/internal/models"
	"oanda-mock-api/internal/repositories"
)

func GetAccount(id string) (models.Account, error) {
	return repositories.GetAccount(id)
}
