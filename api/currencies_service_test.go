package api

import (
	"context"
	"encoding/json"
	"os"
	"testing"

	"exchange-rate-api/db"
	"exchange-rate-api/pb"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestListAllCurrencies(t *testing.T) {
	t.Parallel()

	//--- Reading currencies.json as data test set ---//
	content, err := os.ReadFile("../assets/currencies.json")
	require.NoError(t, err)
	require.NotEmpty(t, content)

	// convert []byte from the json file read to struct
	currencies := make([]db.Currencies, 26)
	err = json.Unmarshal(content, &currencies)
	require.NoError(t, err)
	require.NotEmpty(t, currencies)
	t.Log("ans:", currencies[0].ID)
	t.Log("ans:", currencies[0].CurrencyCode)
	t.Log("ans:", currencies[0].CurrencyName)

	tests := []struct {
		name string
		ans  []db.Currencies
		code codes.Code
	}{
		{
			name: "success",
			ans:  currencies,
		},
	}

	for i := range tests {
		t.Run(tests[i].name, func(t *testing.T) {
			res, err := server.ListAllCurrencies(context.Background(), nil)
			require.NoError(t, err)
			require.NotNil(t, res)
			for j := range res.Currencies {
				assert.Equal(t, tests[i].ans[j].ID, res.Currencies[j].Id)
				assert.Equal(t, tests[i].ans[j].CurrencyCode, res.Currencies[j].CurrencyCode)
				assert.Equal(t, tests[i].ans[j].CurrencyName, res.Currencies[j].CurrencyName)
			}
		})
	}

}
func TestGetCurrency(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		currency pb.GetCurrencyRequest
		ans      pb.GetCurrencyResponse
		code     codes.Code
	}{
		{
			name: "success_IDR",
			currency: pb.GetCurrencyRequest{
				Currency: "IDR",
			},
			ans: pb.GetCurrencyResponse{
				Currency: &pb.Currencies{
					Id:           uint64(1),
					CurrencyCode: "IDR",
					CurrencyName: "INDONESIAN RUPIAH",
				},
			},
			code: codes.OK,
		}, {
			name: "success_KWD",
			currency: pb.GetCurrencyRequest{
				Currency: "KWD",
			},
			ans: pb.GetCurrencyResponse{
				Currency: &pb.Currencies{
					Id:           uint64(14),
					CurrencyCode: "KWD",
					CurrencyName: "KUWAITI DINAR",
				},
			},
			code: codes.OK,
		}, {
			name: "failure_no_currency",
			currency: pb.GetCurrencyRequest{
				Currency: "",
			},
			code: codes.Internal,
		},
	}

	for i := range tests {
		// server := NewApiServer(context.Background(), server)
		t.Run(tests[i].name, func(t *testing.T) {
			res, err := server.GetCurrency(context.Background(), &tests[i].currency)
			if tests[i].code == codes.OK {
				require.NoError(t, err)
				require.NotNil(t, res)
				require.NotEmpty(t, res.Currency.Id)

				assert.Equal(t, tests[i].ans.Currency.Id, res.Currency.Id)
				assert.Equal(t, tests[i].ans.Currency.CurrencyCode, res.Currency.CurrencyCode)
				assert.Equal(t, tests[i].ans.Currency.CurrencyName, res.Currency.CurrencyName)
			} else {
				require.Error(t, err)
				require.Nil(t, res)
				st, ok := status.FromError(err)
				require.True(t, ok)
				require.Equal(t, tests[i].code, st.Code())
			}
		})
	}
}
