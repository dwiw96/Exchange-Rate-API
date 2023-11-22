package postgres

import (
	"context"
	"fmt"
	"testing"
	"time"

	"log"
	"os"

	"exchange-rate-api/tools"

	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	pgPool *DB
	ctx    context.Context
)

func TestMain(m *testing.M) {
	// Load Configuration
	configAPI, err := tools.LoadConfig("../..")
	if err != nil {
		log.Fatal("Cannot load config: ", err)
	}

	//--- Parseconfig ---//
	config, err := pgxpool.ParseConfig(configAPI.DBAddress) // Using environment variables instead of a connection string.
	// config, err := pgxpool.ParseConfig("postgresql://pg:secret@localhost:5432/exchange?sslmode=disable") // Using environment variables instead of a connection string.
	if err != nil {
		log.Fatal(err)
	}

	//--- pgxpool ---//
	dbpool, err := pgxpool.New(context.Background(), config.ConnString())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}
	defer dbpool.Close()

	pgPool = NewDB(dbpool)

	// create context for timeout duration
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(context.Background(), 1200*time.Second)
	defer cancel()

	os.Exit(m.Run())
}
