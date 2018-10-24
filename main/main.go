package main

import (
	"github.com/lk16/cobinhoodGo"
	"log"
	"os"
)

// this file is currently here for quick testing
// unit tests may be provided later

func main() {

	apiKey := os.Getenv("COBINHOOD_APIKEY")

	if apiKey == "" {
		log.Fatalf("please set COBINHOOD_APIKEY environment variable")
	}

	cobin := &cobinhoodgo.Cobin{}
	cobin.SetAPIKey(apiKey)

	wallet, err := cobin.GetBalances()
	if err != nil {
		log.Fatalf("error = %s", err.Error())
	}
	for i, item := range wallet {
		log.Printf("wallet item = %+v", item)
		if i > 4 {
			break
		}
	}

	log.Printf("---")

	tradingPairs, err := cobin.GetTradingPairs()
	if err != nil {
		log.Fatalf("error = %s", err.Error())
	}
	for i, item := range tradingPairs {
		log.Printf("tradingPairs item = %+v", item)
		if i > 4 {
			break
		}
	}

	log.Printf("---")

	ticker, err := cobin.GetTicker("BTC-USDT")
	if err != nil {
		log.Fatalf("error = %s", err.Error())
	}
	log.Printf("ticker = %+v", ticker)
}
