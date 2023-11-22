package postgres

import (
	"context"
	"fmt"
	"log"

	"exchange-rate-api/db"

	"github.com/jackc/pgx/v5"
)

// Insert currencies and exchange rates data from scraping into postgres
// Currencies will be insert first, because 'currency_code' column is referenced
// by 'exchange_rates' table. Then, exchange rates data insert with 'currency_code_from'
// column became Foreign Key <<FK>> referencing 'currency_code' in 'currencies' table.
// This function will be automatically called one a day.
func (conn *DB) InsertExchangeRate(ctx context.Context, currencies []db.Currencies, exchangeRates []db.ExchangeRate) error {
	rows := [][]interface{}{}

	//--- insert into currencies ---//
	for i := range currencies {
		rows = append(rows, []interface{}{currencies[i].CurrencyCode, currencies[i].CurrencyName})
	}

	copyCount, err := conn.db.CopyFrom(
		ctx, pgx.Identifier{"currencies"},
		[]string{"currency_code", "currency_name"},
		pgx.CopyFromRows(rows),
	)
	if err != nil {
		log.Printf("--- (database/server)[1] - bulk insertion to 'currencies' failed, msg:\n%s", err)
		return err
	}
	if copyCount != int64(len(currencies)) {
		log.Printf("--- (database/server)[2] - 'currencies'returned length from copyFrom wrong, \ncopyCount:%d", copyCount)
		return fmt.Errorf("'currencies' returned length: %d, should be: %d\n", copyCount, len(currencies))
	}

	//--- Insert into exchange_rates ---//
	rows = [][]interface{}{}
	for i := range exchangeRates {
		rows = append(rows, []interface{}{exchangeRates[i].CurrencyCodeFrom, exchangeRates[i].CurrencyCodeTo, exchangeRates[i].Buy, exchangeRates[i].Sell})
	}

	copyCount, err = conn.db.CopyFrom(
		ctx, pgx.Identifier{"exchange_rates"},
		[]string{"currency_code_from", "currency_code_to", "buy", "sell"},
		pgx.CopyFromRows(rows),
	)
	if err != nil {
		log.Printf("--- (database/server)[3] bulk insertion from 'exchange_rates' failed, msg:\n%s", err)
		return err
	}
	if copyCount != int64(len(exchangeRates)) {
		log.Printf("--- (database/server)[4] 'exchange_rates' returned length from copyFrom wrong, \ncopyCount:%d", copyCount)
		return fmt.Errorf("returned length: %d, should be: %d\n", copyCount, len(exchangeRates))
	}

	return nil
}
