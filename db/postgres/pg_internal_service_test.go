package postgres

import (
	"encoding/json"
	"os"
	"testing"

	"exchange-rate-api/db"

	"github.com/stretchr/testify/require"
)

func TestCreateExchangeRates(t *testing.T) {
	//--- Reading currencies.json as data test set ---//
	content, err := os.ReadFile("../../assets/currencies.json")
	require.NoError(t, err)
	require.NotEmpty(t, content)

	currencies := make([]db.Currencies, 26)
	err = json.Unmarshal(content, &currencies)
	require.NoError(t, err)

	//--- Reading exchange_rates.json data test set ---//
	content, err = os.ReadFile("../../assets/exchange_rates.json")
	require.NoError(t, err)
	require.NotEmpty(t, content)

	exchangeRates := make([]db.ExchangeRate, 25)
	err = json.Unmarshal(content, &exchangeRates)
	require.NoError(t, err)

	//--- Insert currencies & exchange rates data to postgres ---//
	err = pgPool.InsertExchangeRate(ctx, currencies, exchangeRates)
	require.NoError(t, err)

	t.Run("Test bulk insert currencies", func(t *testing.T) {
	})
}
