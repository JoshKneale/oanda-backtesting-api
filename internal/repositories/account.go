package repositories

import (
	"oanda-mock-api/internal/models"
)

func GetAccount(id string) (models.Account, error) {
	return models.Account{
		Id:       id,
		Currency: models.CurrencyEUR,
		Balance:  1000.0,
		Nav:      1000.0,
	}, nil
}
