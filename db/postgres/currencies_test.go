package postgres

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"exchange-rate-api/db"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestListAllCurrencies(t *testing.T) {
	res, err := pgPool.ListAllCurrencies(ctx)
	require.NoError(t, err)
	require.NotEmpty(t, res)

	//--- Reading currencies.json as data test set ---//
	content, err := os.ReadFile("../../assets/currencies.json")
	require.NoError(t, err)
	require.NotEmpty(t, content)

	// convert []byte from the json file read to struct
	currencies := make([]db.Currencies, 26)
	err = json.Unmarshal(content, &currencies)
	require.NoError(t, err)
	require.NotEmpty(t, currencies)

	//--- Compare the return with json file from sraping ---//
	for i := range currencies {
		assert.Equal(t, currencies[i].CurrencyCode, res[i].CurrencyCode)
		assert.Equal(t, currencies[i].CurrencyName, res[i].CurrencyName)
	}
}

func TestGetCurrency(t *testing.T) {
	tests := []struct {
		code struct {
			name string
			code string
		}
		name struct {
			nameTest     string
			nameCurrency string
		}
		ans struct {
			code db.Currencies
			name db.Currencies
		}
	}{
		{
			code: struct {
				name string
				code string
			}{
				"Test Get Currency With Code", "IDR",
			},
			name: struct {
				nameTest     string
				nameCurrency string
			}{
				"Test Get Currency With Name", "SWISS FRANC",
			},
			ans: struct {
				code db.Currencies
				name db.Currencies
			}{
				code: db.Currencies{
					CurrencyCode: "IDR",
					CurrencyName: "INDONESIAN RUPIAH",
				},
				name: db.Currencies{
					CurrencyCode: "CHF",
					CurrencyName: "SWISS FRANC",
				},
			},
		}, {
			code: struct {
				name string
				code string
			}{
				"Test Get Currency With Code", "DKK",
			},
			name: struct {
				nameTest     string
				nameCurrency string
			}{
				"Test Get Currency With Name", "LAOTIAN KIP",
			},
			ans: struct {
				code db.Currencies
				name db.Currencies
			}{
				code: db.Currencies{
					CurrencyCode: "DKK",
					CurrencyName: "DANISH KRONE",
				},
				name: db.Currencies{
					CurrencyCode: "LAK",
					CurrencyName: "LAOTIAN KIP",
				},
			},
		},
	}

	//--- Test get currency with code ---//
	for _, test := range tests {
		t.Run(test.code.name, func(t *testing.T) {
			res, err := pgPool.GetCurrency(ctx, test.code.code)
			require.NoError(t, err)
			require.NotEmpty(t, res)
			assert.Equal(t, test.ans.code.CurrencyCode, res.CurrencyCode)
			assert.Equal(t, test.ans.code.CurrencyName, res.CurrencyName)
		})
	}
	fmt.Println()

	//--- Test get currency with name ---//
	for _, test := range tests {
		t.Run(test.name.nameTest, func(t *testing.T) {
			res, err := pgPool.GetCurrency(ctx, test.name.nameCurrency)
			require.NoError(t, err)
			require.NotEmpty(t, res)
			assert.Equal(t, test.ans.name.CurrencyCode, res.CurrencyCode)
			assert.Equal(t, test.ans.name.CurrencyName, res.CurrencyName)
		})
	}
}
