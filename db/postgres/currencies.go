package postgres

import (
	"context"
	"log"

	"exchange-rate-api/db"
)

// getting list of all currencies code and name in the database
func (conn *DB) ListAllCurrencies(ctx context.Context) ([]db.Currencies, error) {
	query := `SELECT * FROM currencies;`
	rows, err := conn.db.Query(ctx, query)
	if err != nil {
		log.Println("error (db/postgres/currencies)[1] - Failed to get all currencies data")
		return nil, err
	}
	defer rows.Close()

	res := make([]db.Currencies, 26, 50)
	idx := 0
	for rows.Next() {
		if idx <= len(res) {
			err := rows.Scan(&res[idx].ID, &res[idx].CurrencyCode, &res[idx].CurrencyName)
			if err != nil {
				log.Println("error (db/postgres/currencies)[2] Failed to scan data, msg:", err)
				return nil, err
			}
		} else {
			var temp db.Currencies
			err := rows.Scan(&temp.ID, &temp.CurrencyCode, &temp.CurrencyName)
			if err != nil {
				log.Println("error (db/postgres/currencies)[3] Failed to scan data, msg:", err)
				return nil, err
			}
			res = append(res, temp)
		}
		idx++
	}

	return res, err
}

/* Getting currencies name and code based on name or code */
// The input can take both currency code and name, this func will automatically
// check the input.
func (conn *DB) GetCurrency(ctx context.Context, currency string) (*db.Currencies, error) {
	queryCode := `SELECT * FROM currencies WHERE currency_code=$1;`
	queryName := `SELECT * FROM currencies WHERE currency_name=$1;`
	query := queryCode

	if len(currency) > 3 {
		query = queryName
	}

	var res db.Currencies

	err := conn.db.QueryRow(ctx, query, currency).Scan(&res.ID, &res.CurrencyCode, &res.CurrencyName)
	if err != nil {
		log.Println("error (db/postgres/currencies)[1] Failed to get currency, msg:", err)
		return nil, err
	}

	return &res, err
}
