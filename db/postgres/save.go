package postgres

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"exchange-rate-api/db"
)

func (conn *DB) SaveDataToDB(ctx context.Context) error {
	log.Println("save data to postgreSQL")
	//--- Reading currencies.json as data test set ---//
	content, err := os.ReadFile("./assets/currencies.json")
	if err != nil {
		log.Println("error (db/postgres/save)[1] - Failed to read currencies data from json file")
		return fmt.Errorf("Failed to read currencies data from json file, msg: %v\n", err)
	}
	if content == nil {
		return fmt.Errorf("the currencies return data from json file is empty or nil")
	}

	currencies := make([]db.Currencies, 26)
	err = json.Unmarshal(content, &currencies)
	if err != nil {
		return fmt.Errorf("Failed to unmarshall currencies json to struct, msg: %v\n", err)
	}
	if currencies == nil {
		return fmt.Errorf("the unmarshalled currencies data is empty or nil")
	}

	//--- Reading exchange_rates.json data test set ---//
	content, err = os.ReadFile("./assets/exchange_rates.json")
	if err != nil {
		return fmt.Errorf("Failed to read exchange rates data from json file, msg: %v\n", err)
	}
	if content == nil {
		return fmt.Errorf("the return exchange rates data from json file is empty or nil")
	}

	exchangeRates := make([]db.ExchangeRate, 25)
	err = json.Unmarshal(content, &exchangeRates)
	if err != nil {
		return fmt.Errorf("Failed to unmarshall exchange rates json to struct, msg: %v\n", err)
	}
	if currencies == nil {
		return fmt.Errorf("the unmarshalled exchange rates data is empty or nil")
	}

	//--- Insert currencies & exchange rates data to postgres ---//
	err = conn.InsertExchangeRate(ctx, currencies, exchangeRates)
	if err != nil {
		return fmt.Errorf("Failed to insert data to postgreSQL, msg: %v\n", err)
	}

	return nil
}
