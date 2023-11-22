package postgres

import (
	"context"
	"errors"
	"log"

	"exchange-rate-api/db"
)

// List all exchange rates data

/* Getting rates data for spesicif currency.*/
// Both input are use currency code like 'IDR' or 'USD' in string.
// This func returning ExchangeRate struct.
func (conn *DB) GetRates(ctx context.Context, from_currency, to_currency string) (*db.ExchangeRate, error) {
	var res db.ExchangeRate
	query := `SELECT * FROM exchange_rates WHERE currency_code_from=$1 AND currency_code_to=$2;`
	err := conn.db.QueryRow(ctx, query, from_currency, to_currency).Scan(&res.ID, &res.CurrencyCodeFrom, &res.CurrencyCodeTo, &res.Buy, &res.Sell, &res.ValidateDate)
	if err != nil {
		log.Println("error (db/postgres/client)[1] - Failed to get rates data")
		return nil, err
	}

	if res == (db.ExchangeRate{}) {
		return &res, errors.New("error (db/postgres/client)[2] - result is empty")
	}

	return &res, err
}

// getting buy rates data for spesicif currency
func (conn *DB) GetBuy(ctx context.Context, from_currency, to_currency string) (res float64, err error) {
	query := `SELECT buy FROM exchange_rates WHERE currency_code_from=$1 AND currency_code_to=$2;`
	err = conn.db.QueryRow(ctx, query, from_currency, to_currency).Scan(&res)
	if err != nil {
		log.Println("error (db/postgres/client)[1] - Failed to get buy rates data")
		return 0, err
	}

	if res <= 0 {
		return res, errors.New("error (db/postgres/client)[2] - result is empty")
	}

	return res, err
}

// getting sell rates data for spesicif currency
func (conn *DB) GetSell(ctx context.Context, from_currency, to_currency string) (res float64, err error) {
	query := `SELECT sell FROM exchange_rates WHERE currency_code_from=$1 AND currency_code_to=$2;`
	err = conn.db.QueryRow(ctx, query, from_currency, to_currency).Scan(&res)
	if err != nil {
		log.Println("error (db/postgres/client)[1] - Failed to get sell rates data")
		return 0, err
	}

	if res <= 0 {
		return res, errors.New("error (db/postgres/client)[2] - result is empty")
	}

	return res, err
}
