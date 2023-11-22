package pkg

import (
	"encoding/json"
	"os"
	"testing"

	"exchange-rate-api/db"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestWriteJsonFile(t *testing.T) {
	t.Run("currencies test", func(t *testing.T) {
		currencies := []db.Currencies{
			{
				CurrencyCode: "IDR",
				CurrencyName: "INDONESIAN RUPIAH",
			}, {
				CurrencyCode: "CDN",
				CurrencyName: "CANADIAN DOLLAR",
			}, {
				CurrencyCode: "DKK",
				CurrencyName: "DANISH KRONE",
			},
		}

		err := WriteJsonFile("currencies_test.json", currencies)
		require.NoError(t, err)

		// Reading into struct type from a JSON file
		content, err := os.ReadFile("currencies_test.json")
		require.NoError(t, err)

		res := []db.Currencies{}

		err = json.Unmarshal(content, &res)
		require.NoError(t, err)
		assert.Equal(t, 3, len(res))

		for i := 0; i < len(currencies)-1; i++ {
			assert.Equal(t, currencies[i].CurrencyCode, res[i].CurrencyCode)
			assert.Equal(t, currencies[i].CurrencyName, res[i].CurrencyName)
		}
	})

	t.Run("exchange_rates test", func(t *testing.T) {
		rates := []db.ExchangeRate{
			{
				CurrencyCodeFrom: "IDR",
				CurrencyCodeTo:   "USD",
				Buy:              14500,
				Sell:             15000,
			}, {
				CurrencyCodeFrom: "AUD",
				CurrencyCodeTo:   "SGD",
				Buy:              11000,
				Sell:             11500,
			}, {
				CurrencyCodeFrom: "JPY",
				CurrencyCodeTo:   "EUR",
				Buy:              17000,
				Sell:             17500,
			},
		}

		err := WriteJsonFile("exchange_rates_test.json", rates)
		require.NoError(t, err)

		// Reading into struct type from a JSON file
		content, err := os.ReadFile("exchange_rates_test.json")
		require.NoError(t, err)

		res := []db.ExchangeRate{}

		err = json.Unmarshal(content, &res)
		require.NoError(t, err)
		assert.Equal(t, 3, len(res))

		for i := 0; i < len(rates)-1; i++ {
			assert.Equal(t, rates[i].CurrencyCodeTo, res[i].CurrencyCodeTo)
			assert.Equal(t, rates[i].Buy, res[i].Buy)
			assert.Equal(t, rates[i].Sell, res[i].Sell)
		}
	})
}
