package api

import (
	"context"
	"log"

	"exchange-rate-api/pb"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *apiServer) ListAllCurrencies(ctx context.Context, empty *empty.Empty) (*pb.ListAllCurrenciesResponse, error) {
	log.Printf("recieve a list-all-currencies request")
	currencies, err := server.DB.ListAllCurrencies(ctx)
	if err != nil {
		log.Println("Failed to get all the currencies list, msg:", err)
		return nil, status.Error(codes.Internal, "Failed to get all the currencies list")
	}

	err = checkContext(ctx)
	if err != nil {
		return nil, err
	}

	res := pb.ListAllCurrenciesResponse{}

	for i := range currencies {
		var temp pb.Currencies
		temp.Id = uint64(currencies[i].ID)
		temp.CurrencyCode = currencies[i].CurrencyCode
		temp.CurrencyName = currencies[i].CurrencyName
		// log.Println("temp:", &temp)
		// log.Println("res:", res)
		// log.Println("res.Currencies:", res.Currencies)
		res.Currencies = append(res.Currencies, &temp)
	}

	return &res, nil
}

func (server *apiServer) GetCurrency(ctx context.Context, req *pb.GetCurrencyRequest) (*pb.GetCurrencyResponse, error) {
	log.Printf("recieve a get-currency request with currency: %s\n", req.Currency)
	currency, err := server.DB.GetCurrency(ctx, req.Currency)
	if err != nil {
		return nil, status.Error(codes.Internal, "Failed to get the currency")
	}

	err = checkContext(ctx)
	if err != nil {
		return nil, err
	}

	res := &pb.GetCurrencyResponse{
		Currency: &pb.Currencies{
			Id:           uint64(currency.ID),
			CurrencyCode: currency.CurrencyCode,
			CurrencyName: currency.CurrencyName,
		},
	}

	return res, nil
}
