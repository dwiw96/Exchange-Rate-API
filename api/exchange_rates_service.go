package api

import (
	"context"
	"log"

	"exchange-rate-api/pb"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (server *apiServer) GetRate(ctx context.Context, req *pb.GetRateRequest) (*pb.GetRateResponse, error) {
	log.Printf("recieve get-rates request for data from: %s to: %s\n", req.FromCurrency, req.ToCurrency)
	rates, err := server.DB.GetRates(ctx, req.FromCurrency, req.ToCurrency)
	if err != nil {
		log.Printf("Failed to get exchange rates from: %s to: %s\n", req.FromCurrency, req.ToCurrency)
		return nil, status.Error(codes.Internal, "Failed to get exchange rates data")
	}

	err = checkContext(ctx)
	if err != nil {
		return nil, err
	}

	res := pb.GetRateResponse{
		Id:               rates.ID,
		CurrencyCodeFrom: rates.CurrencyCodeFrom,
		CurrencyCodeTo:   rates.CurrencyCodeTo,
		Buy:              rates.Buy,
		Sell:             rates.Sell,
		ValidateDate:     timestamppb.New(rates.ValidateDate),
	}

	log.Println("rates: ", rates.ValidateDate)
	log.Println("res: ", res.ValidateDate)
	t := res.ValidateDate
	log.Println("convert: ", t.AsTime())

	return &res, nil
}
func (server *apiServer) GetBuy(ctx context.Context, req *pb.GetRateRequest) (*pb.GetBuyResponse, error) {
	log.Printf("recieve get-buy request for data from: %s to: %s\n", req.FromCurrency, req.ToCurrency)
	rates, err := server.DB.GetBuy(ctx, req.FromCurrency, req.ToCurrency)
	if err != nil {
		log.Printf("Failed to get buy rates from: %s to: %s\n", req.FromCurrency, req.ToCurrency)
		return nil, status.Error(codes.Internal, "Failed to get buy rates data")
	}

	err = checkContext(ctx)
	if err != nil {
		return nil, err
	}

	res := pb.GetBuyResponse{
		Buy: rates,
	}

	return &res, nil
}
func (server *apiServer) GetSell(ctx context.Context, req *pb.GetRateRequest) (*pb.GetSellResponse, error) {
	log.Printf("recieve get-sell request for data from: %s to: %s\n", req.FromCurrency, req.ToCurrency)
	rates, err := server.DB.GetSell(ctx, req.FromCurrency, req.ToCurrency)
	if err != nil {
		log.Printf("Failed to get sell rates from: %s to: %s\n", req.FromCurrency, req.ToCurrency)
		return nil, status.Error(codes.Internal, "Failed to get sell rates data")
	}

	err = checkContext(ctx)
	if err != nil {
		return nil, err
	}

	res := pb.GetSellResponse{
		Sell: rates,
	}

	return &res, nil
}
