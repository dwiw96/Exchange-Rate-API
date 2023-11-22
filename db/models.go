package db

import "time"

type Currencies struct {
	ID           uint64 `json:"id"`
	CurrencyCode string `json:"currency_code"`
	CurrencyName string `json:"currency_name"`
}
type ExchangeRate struct {
	ID               uint64    `json:"id"`
	CurrencyCodeFrom string    `json:"currency_code_from"`
	CurrencyCodeTo   string    `json:"currency_code_to"`
	Buy              float64   `json:"buy"`
	Sell             float64   `json:"sell"`
	ValidateDate     time.Time `json:"validate_date"`
}
