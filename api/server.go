package api

import (
	"context"
	"log"
	"net"

	pg "exchange-rate-api/db/postgres"
	pb "exchange-rate-api/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type apiServer struct {
	pb.UnimplementedCurrenciesAPIServer
	pb.UnimplementedExchangeRateAPIServer
	DB  *pg.DB
	ctx context.Context
}

func NewApiServer(ctx context.Context, pg *pg.DB) *apiServer {
	server := &apiServer{
		DB:  pg,
		ctx: ctx,
	}

	return server
}

func (server *apiServer) StartServer(address string) {
	log.Printf("start server on port: %s\n", address)
	// Create grpc server
	grpcServer := grpc.NewServer()
	// Register service server on the grpc server
	pb.RegisterCurrenciesAPIServer(grpcServer, server)
	pb.RegisterExchangeRateAPIServer(grpcServer, server)
	// reflection allow the gRPC client to explore what RPCs are available on
	// the server, and how to call them (like some self documentations for the server)
	reflection.Register(grpcServer)

	// Create an address string with the port we get before
	// address := fmt.Sprintf("0.0.0.0:%d", *port)
	// Listen to TCP connection.
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal("Failed to listen TCP connection, msg: ", err)
	}

	// Start listening to the request
	// This func is blocking call, so it should run in parallel
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("Cannot start server, msg: ", err)
	}
}
