package main

import (
	"github.com/lk16/cobinhoodGo"
	"log"
	"os"
	"time"
)

// this file is currently here for quick testing

// WARNING: this cancels all your open orders and places some new ones
// afterwards all orders will be cancelled
// USE WITH CARE

func main() {

	apiKey := os.Getenv("COBINHOOD_APIKEY")

	if apiKey == "" {
		log.Fatalf("please set COBINHOOD_APIKEY environment variable")
	}

	// lower timeout is recommended for production use
	client := cobinhood.NewClient(apiKey, 5*time.Second)

	log.Printf("Cancel all orders:")
	err := client.CancelAllOrders()
	if err != nil {
		log.Fatalf("error = %s", err.Error())
	}

	log.Printf("---")
	log.Printf("Get balance:")

	wallet, err := client.GetBalances()
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
	log.Printf("Get trading pairs:")

	tradingPairs, err := client.GetTradingPairs()
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
	log.Printf("Get ticker:")

	ticker, err := client.GetTicker("BTC-USDT")
	if err != nil {
		log.Fatalf("error = %s", err.Error())
	}
	log.Printf("ticker = %+v", ticker)

	log.Printf("---")
	log.Printf("Place two orders:")

	placedOrder, err := client.PlaceOrder("ETH-USDT", "bid", "limit", 0.01, 10)
	if err != nil {
		log.Fatalf("error = %s", err.Error())
	}
	log.Printf("placedOrder = %+v", placedOrder)

	placedOrder, err = client.PlaceOrder("ETH-USDT", "ask", "limit", 500, 0.12)
	if err != nil {
		log.Fatalf("error = %s", err.Error())
	}
	log.Printf("placedOrder = %+v", placedOrder)

	log.Printf("---")
	log.Printf("List orders:")

	openOrders, err := client.GetOpenOrders()
	if err != nil {
		log.Fatalf("error = %s", err.Error())
	}
	for i, item := range openOrders {
		log.Printf("openOrders item = %+v", item)
		if i > 4 {
			break
		}
	}

	log.Printf("---")
	log.Printf("Cancel all orders:")

	err = client.CancelAllOrders()
	if err != nil {
		log.Fatalf("error = %s", err.Error())
	}

	log.Printf("---")
	log.Printf("List orders again:")

	openOrders, err = client.GetOpenOrders()
	if err != nil {
		log.Fatalf("error = %s", err.Error())
	}
	for i, item := range openOrders {
		log.Printf("openOrders item = %+v", item)
		if i > 4 {
			break
		}
	}
	log.Printf("---")
	log.Printf("List orderbook:")
	orderbook, err := client.GetOrderBook("ETH-USDT", 3)
	if err != nil {
		log.Fatalf("error = %s", err.Error())
	}
	for _, item := range orderbook.Bids {
		log.Printf("bid = %+v", item)
	}
	log.Printf("-")
	for _, item := range orderbook.Asks {
		log.Printf("ask = %+v", item)
	}
}
