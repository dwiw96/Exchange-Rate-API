package main

import (
	"context"
	"log"

	"exchange-rate-api/db"
	"exchange-rate-api/pb"
	"exchange-rate-api/tools"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// Load Configuration
	configAPI, err := tools.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load config: ", err)
	}

	conn, err := grpc.Dial(configAPI.ServerAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("Cannot dial server: ", err)
	}
	log.Printf("dial server %s\n", configAPI.ServerAddress)

	// Create a new client object
	curencyClient := pb.NewCurrenciesAPIClient(conn)
	ratesClient := pb.NewExchangeRateAPIClient(conn)

	req := &pb.GetCurrencyRequest{
		Currency: "THB",
	}

	res, err := curencyClient.GetCurrency(context.Background(), req)
	if err != nil {
		log.Fatal("can't get currency, msg:", err)
		return
	}

	log.Printf("get currency: %s", res.Currency)

	req2 := &pb.GetRateRequest{
		FromCurrency: "IDR",
		ToCurrency:   "NOK",
	}

	res2, err := ratesClient.GetRate(context.Background(), req2)
	if err != nil {
		log.Fatal("can't get currency rate, msg:", err)
		return
	}
	newRes2 := db.ExchangeRate{
		ID:               res2.Id,
		CurrencyCodeFrom: res2.CurrencyCodeFrom,
		CurrencyCodeTo:   res2.CurrencyCodeTo,
		Buy:              res2.Buy,
		Sell:             res2.Sell,
		ValidateDate:     res2.ValidateDate.AsTime(),
	}

	log.Println("get rate: ", newRes2)
}
