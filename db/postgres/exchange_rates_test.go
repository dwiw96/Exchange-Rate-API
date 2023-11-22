package postgres

import (
	"fmt"
	"testing"

	"exchange-rate-api/db"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var tests = []struct {
	name string
	from string
	to   string
	ans  db.ExchangeRate
}{
	{
		name: "IDR",
		from: "IDR",
		to:   "AUD",
		ans: db.ExchangeRate{
			CurrencyCodeFrom: "IDR",
			CurrencyCodeTo:   "AUD",
			Buy:              9985.52,
			Sell:             10093.75,
		},
	}, {
		name: "IDR",
		from: "IDR",
		to:   "CNY",
		ans: db.ExchangeRate{
			CurrencyCodeFrom: "IDR",
			CurrencyCodeTo:   "CNY",
			Buy:              2137.14,
			Sell:             2158.74,
		},
	}, {
		name: "IDR",
		from: "IDR",
		to:   "EUR",
		ans: db.ExchangeRate{
			CurrencyCodeFrom: "IDR",
			CurrencyCodeTo:   "EUR",
			Buy:              16665.37,
			Sell:             16839.17,
		},
	}, {
		name: "IDR",
		from: "IDR",
		to:   "LAK",
		ans: db.ExchangeRate{
			CurrencyCodeFrom: "IDR",
			CurrencyCodeTo:   "LAK",
			Buy:              0.75,
			Sell:             0.76,
		},
	}, {
		name: "IDR",
		from: "IDR",
		to:   "VND",
		ans: db.ExchangeRate{
			CurrencyCodeFrom: "IDR",
			CurrencyCodeTo:   "VND",
			Buy:              0.64,
			Sell:             0.65,
		},
	},
}

func TestGetRates(t *testing.T) {
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res, err := pgPool.GetRates(ctx, test.from, test.to)
			require.NoError(t, err)
			require.NotEmpty(t, res)

			assert.Equal(t, test.ans.CurrencyCodeFrom, res.CurrencyCodeFrom)
			assert.Equal(t, test.ans.CurrencyCodeTo, res.CurrencyCodeTo)
			assert.Equal(t, test.ans.Buy, res.Buy)
			assert.Equal(t, test.ans.Sell, res.Sell)
			fmt.Println()
		})
	}
}

func TestGetBuy(t *testing.T) {
	for _, test := range tests {
		t.Run("Test Get Buy", func(t *testing.T) {
			res, err := pgPool.GetBuy(ctx, test.from, test.to)
			require.NoError(t, err)
			require.NotEmpty(t, res)

			assert.Equal(t, test.ans.Buy, res)
			fmt.Println()
		})
	}
}

func TestGetSell(t *testing.T) {
	for _, test := range tests {
		t.Run("Test Get Sell", func(t *testing.T) {
			res, err := pgPool.GetSell(ctx, test.from, test.to)
			require.NoError(t, err)
			require.NotEmpty(t, res)

			assert.Equal(t, test.ans.Sell, res)
			fmt.Println()
		})
	}
}
