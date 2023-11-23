# exchange-rate-API
gRPC API and Web Scrapping project that scrape exchange rates from BI website.
(https://www.bi.go.id/id/statistik/informasi-kurs/transaksi-bi/default.aspx)

# Overview
This project is gRPC API for getting exchange rate value, like from_currency, to_currency, buy, sell and validate rates.
Use Go programming language, PostgreSQL database, and proto buffer for gRPC.

## Supported Features
* User can get full exchange rates information.
* User can get all exchange rates data.
* User can get only 'buy' or 'sell' value from-to specific currency.

## Installation Guide
* Clone/download this project into your local machine
  ` git clone https://github.com/dwiw96/exchange-rate-API.git `
* The main branch is the most stable branch at any given time, ensure you're working from it.
* This project run database via docker, so ensure that you installed docker in your machine.
* Install postgres image in your docker, see this guide (https://hub.docker.com/_/postgres).

### To Run
For more information, see the Makefile (https://github.com/dwiw96/exchange-rate-API/blob/main/Makefile)
* To run the postgresql server inside docker, run this task:
```
make dockerStart
```
  note: I set database to remove when docker stop running.
* Create database table using migrate library.
```
make migrateUp
```
* To Login into database using sql.
```
make dockerExec
```
* Run local server
  Server is run at http://localhost:9090
```
make runServer
```
* Run client to test the api
```
make runClient
```

### Request
Get exchange rates from: "IDR" - to:"NOK"
```
ratesClient := pb.NewExchangeRateAPIClient(conn)

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
```

## Technologies Used
* [gRPC]
* [Proto buffer]
* [Go Programming Language]
* [Postgresql] sql database for saving all data in this project
* [Docker] (https://www.docker.com/) Docker is a software platform that allows you to build, test, and deploy applications quickly. - amazon
* [PASETO] (https://paseto.io/) Paseto (Platform-Agnostic SEcurity TOkens) is a specification and reference implementation for secure stateless tokens.
