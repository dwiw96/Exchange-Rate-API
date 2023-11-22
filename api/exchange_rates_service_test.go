package api

import (
	"context"
	"encoding/json"
	"os"
	"testing"
	"time"

	"exchange-rate-api/db"
	"exchange-rate-api/pb"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestGetRate(t *testing.T) {
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

	// get today date's
	date := time.Now()
	date.Round(0)

	tests := []struct {
		name string
		req  pb.GetRateRequest
		ans  pb.GetRateResponse
		code codes.Code
	}{
		{
			name: "success IDR - PHP",
			req: pb.GetRateRequest{
				FromCurrency: "IDR",
				ToCurrency:   "PHP",
			},
			ans: pb.GetRateResponse{
				Id:               19,
				CurrencyCodeFrom: "IDR",
				CurrencyCodeTo:   "PHP",
				Buy:              276.18,
				Sell:             279.06,
				ValidateDate:     timestamppb.New(date),
			},
			code: codes.OK,
		}, {
			name: "fail: empty from-currency",
			req: pb.GetRateRequest{
				ToCurrency: "VND",
			},
			code: codes.Internal,
		}, {
			name: "fail: wrong to-currency",
			req: pb.GetRateRequest{
				FromCurrency: "IDR",
				ToCurrency:   "DBD",
			},
			code: codes.Internal,
		},
	}

	for i := range tests {
		t.Run(tests[i].name, func(t *testing.T) {
			res, err := server.GetRate(context.Background(), &tests[i].req)
			if tests[i].code == codes.OK {
				require.NoError(t, err)
				require.NotNil(t, res)

				assert.Equal(t, tests[i].ans.Id, res.Id)
				assert.Equal(t, tests[i].ans.CurrencyCodeFrom, res.CurrencyCodeFrom)
				assert.Equal(t, tests[i].ans.CurrencyCodeTo, res.CurrencyCodeTo)
				assert.Equal(t, tests[i].ans.Buy, res.Buy)
				assert.Equal(t, tests[i].ans.Sell, res.Sell)
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

func TestGetBuy(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		req  pb.GetRateRequest
		ans  pb.GetBuyResponse
		code codes.Code
	}{
		{
			name: "success_IDR_DKK",
			req: pb.GetRateRequest{
				FromCurrency: "IDR",
				ToCurrency:   "DKK",
			},
			ans:  pb.GetBuyResponse{Buy: 2246.88},
			code: codes.OK,
		}, {
			name: "fail_empty_from_currency",
			req: pb.GetRateRequest{
				ToCurrency: "VND",
			},
			code: codes.Internal,
		}, {
			name: "fail_wrong_to_currency",
			req: pb.GetRateRequest{
				FromCurrency: "IDR",
				ToCurrency:   "DBD",
			},
			code: codes.Internal,
		},
	}

	for i := range tests {
		t.Run(tests[i].name, func(t *testing.T) {
			res, err := server.GetBuy(context.Background(), &tests[i].req)
			if tests[i].code == codes.OK {
				require.NoError(t, err)
				require.NotNil(t, res)

				assert.Equal(t, tests[i].ans.Buy, res.Buy)
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

func TestGetSell(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		req  pb.GetRateRequest
		ans  pb.GetSellResponse
		code codes.Code
	}{
		{
			name: "success_IDR_DKK",
			req: pb.GetRateRequest{
				FromCurrency: "IDR",
				ToCurrency:   "DKK",
			},
			ans:  pb.GetSellResponse{Sell: 2269.79},
			code: codes.OK,
		}, {
			name: "fail_empty_from_currency",
			req: pb.GetRateRequest{
				ToCurrency: "VND",
			},
			code: codes.Internal,
		}, {
			name: "fail_wrong_to_currency",
			req: pb.GetRateRequest{
				FromCurrency: "IDR",
				ToCurrency:   "DBD",
			},
			code: codes.Internal,
		},
	}

	for i := range tests {
		t.Run(tests[i].name, func(t *testing.T) {
			res, err := server.GetSell(context.Background(), &tests[i].req)
			if tests[i].code == codes.OK {
				require.NoError(t, err)
				require.NotNil(t, res)

				assert.Equal(t, tests[i].ans.Sell, res.Sell)
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
