package pkg

import (
	"encoding/json"
	"os"
	"testing"

	"exchange-rate-api/db"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestScraping(t *testing.T) {
	currencies, exchangeRates, err := Scrapping()
	require.NoError(t, err)
	require.NotEmpty(t, currencies)
	assert.Equal(t, 26, len(currencies))
	require.NotEmpty(t, exchangeRates)
	assert.Equal(t, 25, len(exchangeRates))

	t.Run("Test currencies scraping", func(t *testing.T) {
		// Reading into struct type from a JSON file
		content, err := os.ReadFile("../assets/currencies.json")
		require.NoError(t, err)

		jsonData := []db.ExchangeRate{}

		err = json.Unmarshal(content, &jsonData)
		require.NoError(t, err)
		require.NotEmpty(t, jsonData)
		assert.Equal(t, 26, len(jsonData))
	})

	t.Run("Test exchange_rates scraping", func(t *testing.T) {
		// Reading into struct type from a JSON file
		content, err := os.ReadFile("../assets/exchange_rates.json")
		require.NoError(t, err)

		jsonData := []db.ExchangeRate{}

		err = json.Unmarshal(content, &jsonData)
		require.NoError(t, err)
		require.NotEmpty(t, jsonData)
		assert.Equal(t, 25, len(jsonData))
	})
}
