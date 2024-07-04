package models

import "encoding/json"

type Currency int64

const (
	CurrencyUndefined Currency = iota
	CurrencyUSD
	CurrencyEUR
	CurrencyGBP
)

func (c Currency) String() string {
	return [...]string{"Undefined", "USD", "EUR", "GBP"}[c]
}

func (c Currency) IsValid() bool {
	switch c {
	case CurrencyUSD, CurrencyEUR, CurrencyGBP:
		return true
	}
	return false
}

func (c Currency) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.String())
}

type Account struct {
	Id       string   `json:"id"`
	Currency Currency `json:"currency"`
	Balance  float64  `json:"balance"`
	Nav      float64  `json:"NAV"`
}
