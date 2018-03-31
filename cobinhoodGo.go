package cobinhoodgo

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"
	"time"
)

// Cobin holds your Cobinhood configuration
type Cobin struct {
	APIKey string
}

// SetAPIKey set your Coobinhood API Key
func (c *Cobin) SetAPIKey(key string) {
	c.APIKey = key
}

// Wallet hold you wallet balances
type Wallet struct {
	Currency string
	Type     string
	Total    float64
	OnOrder  float64
	Locked   bool
	UsdValue float64
	BtcValue float64
}

// Ticker hold the ticker information for the coin pair
type Ticker struct {
	TradingPairID  string
	Timestamp      int64
	Two4HHigh      string
	Two4HLow       string
	Two4HOpen      string
	Two4HVolume    string
	LastTradePrice float64
	HighestBid     float64
	LowestAsk      float64
}

// OpenOrder hold Open Orders information
type OpenOrder struct {
	ID            string
	TradingPairID string
	Side          string
	Type          string
	Price         float64
	Size          float64
	Filled        float64
	State         string
	Timestamp     int64
	EqPrice       string
	CompletedAt   string
	TradingPair   string
}

// Order to hold the information about an order
type Order struct {
	ID          string
	TradingPair string
	State       string
	Side        string
	Type        string
	Price       float64
	Size        float64
	Filled      float64
	Timestamp   int64
}

// PlaceOrderData is for the necessary data to place an order
type PlaceOrderData struct {
	TradingPairID string
	Side          string
	Type          string
	Price         float64
	Size          float64
}

// PlaceOrderResult holds the result for placing an order
type PlaceOrderResult struct {
	ID          string
	TradingPair string
	State       string
	Side        string
	Type        string
	Price       float64
	Size        float64
	Filled      float64
	Timestamp   int64
	EqPrice     string
}

// GetWallet returns a slice of type Wallet with balances of your account
func GetWallet(c Cobin) ([]Wallet, error) {
	var myWallet []Wallet
	walletBalances := wallet{}

	err := requestCobinhood(c, "GET", "https://api.cobinhood.com/v1/wallet/balances", nil, &walletBalances)
	if err != nil {
		return nil, err
	}

	var w Wallet
	for i := range walletBalances.Result.Balances {
		w.Currency = walletBalances.Result.Balances[i].Currency
		w.Type = walletBalances.Result.Balances[i].Type
		w.Total, _ = strconv.ParseFloat(walletBalances.Result.Balances[i].Total, 64)
		w.OnOrder, _ = strconv.ParseFloat(walletBalances.Result.Balances[i].OnOrder, 64)
		w.Locked = walletBalances.Result.Balances[i].Locked
		w.UsdValue, _ = strconv.ParseFloat(walletBalances.Result.Balances[i].UsdValue, 64)
		w.BtcValue, _ = strconv.ParseFloat(walletBalances.Result.Balances[i].BtcValue, 64)

		myWallet = append(myWallet, w)
	}
	return myWallet, nil
}

// GetTicker returns a slice of type Ticker with the exchange/ticker information for the pair
func GetTicker(c Cobin, coinPair []string) ([]Ticker, error) {
	var tick []Ticker
	cobinhoodTicker := ticker{}

	var t Ticker
	for i := range coinPair {
		err := requestCobinhood(c, "GET", "https://api.cobinhood.com/v1/market/tickers/"+coinPair[i], nil, &cobinhoodTicker)
		if err != nil {
			return nil, err
		}

		t.HighestBid, _ = strconv.ParseFloat(cobinhoodTicker.Result.Ticker.HighestBid, 64)
		t.LastTradePrice, _ = strconv.ParseFloat(cobinhoodTicker.Result.Ticker.LastTradePrice, 64)
		t.LowestAsk, _ = strconv.ParseFloat(cobinhoodTicker.Result.Ticker.LowestAsk, 64)
		t.Timestamp = cobinhoodTicker.Result.Ticker.Timestamp
		t.TradingPairID = cobinhoodTicker.Result.Ticker.TradingPairID
		t.Two4HHigh = cobinhoodTicker.Result.Ticker.Two4HHigh
		t.Two4HLow = cobinhoodTicker.Result.Ticker.Two4HLow
		t.Two4HOpen = cobinhoodTicker.Result.Ticker.Two4HOpen
		t.Two4HVolume = cobinhoodTicker.Result.Ticker.Two4HVolume

		tick = append(tick, t)
	}

	return tick, nil
}

// GetOpenOrders returna slice of type OpenOrder with all your open orders at the exchange
func GetOpenOrders(c Cobin) ([]OpenOrder, error) {
	var myOpenOrders []OpenOrder
	openOrders := openorders{}

	err := requestCobinhood(c, "GET", "https://api.cobinhood.com/v1/trading/orders", nil, &openOrders)
	if err != nil {
		return nil, err
	}

	var oo OpenOrder
	for i := range openOrders.Result.Orders {
		oo.EqPrice = openOrders.Result.Orders[i].EqPrice
		oo.Filled, _ = strconv.ParseFloat(openOrders.Result.Orders[i].Filled, 64)
		oo.ID = openOrders.Result.Orders[i].ID
		oo.Price, _ = strconv.ParseFloat(openOrders.Result.Orders[i].Price, 64)
		oo.Side = openOrders.Result.Orders[i].Side
		oo.Size, _ = strconv.ParseFloat(openOrders.Result.Orders[i].Size, 64)
		oo.State = openOrders.Result.Orders[i].State
		oo.Timestamp = openOrders.Result.Orders[i].Timestamp
		oo.TradingPair = openOrders.Result.Orders[i].TradingPairID
		oo.Type = openOrders.Result.Orders[i].Type

		myOpenOrders = append(myOpenOrders, oo)
	}
	return myOpenOrders, nil
}

// GetOrderStatus returns the status of an Order
func GetOrderStatus(c Cobin, orderID string) (Order, error) {
	var myOrderStatus Order
	orderStatus := order{}

	err := requestCobinhood(c, "GET", "https://api.cobinhood.com/v1/trading/orders/"+orderID, nil, &orderStatus)
	if err != nil {
		return myOrderStatus, err
	}

	myOrderStatus.Filled, _ = strconv.ParseFloat(orderStatus.Result.Order.Filled, 64)
	myOrderStatus.ID = orderStatus.Result.Order.ID
	myOrderStatus.Price, _ = strconv.ParseFloat(orderStatus.Result.Order.Price, 64)
	myOrderStatus.Side = orderStatus.Result.Order.Side
	myOrderStatus.Size, _ = strconv.ParseFloat(orderStatus.Result.Order.Size, 64)
	myOrderStatus.State = orderStatus.Result.Order.State
	myOrderStatus.Timestamp = orderStatus.Result.Order.Timestamp
	myOrderStatus.TradingPair = orderStatus.Result.Order.TradingPair
	myOrderStatus.Type = orderStatus.Result.Order.Type

	return myOrderStatus, nil
}

// PlaceOrder lets you place an order in the exchange
func PlaceOrder(c Cobin, po PlaceOrderData) (PlaceOrderResult, error) {
	var myPOResult PlaceOrderResult
	newOrder := &placeorder{
		TradingPairID: po.TradingPairID,
		Side:          po.Side,
		Type:          po.Type,
		Price:         strconv.FormatFloat(po.Price, 'f', 8, 64),
		Size:          strconv.FormatFloat(po.Size, 'f', 4, 64),
	}

	orderJSON, _ := json.Marshal(newOrder)

	placedOrder := placeorderresult{}

	err := requestCobinhood(c, "POST", "https://api.cobinhood.com/v1/trading/orders", bytes.NewReader(orderJSON), &placedOrder)
	if err != nil {
		return myPOResult, err
	}

	myPOResult.EqPrice = placedOrder.Result.Order.EqPrice
	myPOResult.Filled, _ = strconv.ParseFloat(placedOrder.Result.Order.Filled, 64)
	myPOResult.ID = placedOrder.Result.Order.ID
	myPOResult.Price, _ = strconv.ParseFloat(placedOrder.Result.Order.Price, 64)
	myPOResult.Side = placedOrder.Result.Order.Side
	myPOResult.Size, _ = strconv.ParseFloat(placedOrder.Result.Order.Size, 64)
	myPOResult.State = placedOrder.Result.Order.State
	myPOResult.Timestamp = placedOrder.Result.Order.Timestamp
	myPOResult.TradingPair = placedOrder.Result.Order.TradingPair
	myPOResult.Type = placedOrder.Result.Order.Type

	return myPOResult, nil
}

// send request to Cobinhood
func requestCobinhood(c Cobin, postType string, apiURL string, body io.Reader, target interface{}) error {
	if c.APIKey == "" {
		return errors.New("Api Key can't be empty")
	}

	client := &http.Client{Timeout: 30 * time.Second}
	req, err := http.NewRequest(postType, apiURL, body)

	req.Header.Add("Authorization", c.APIKey)
	req.Header.Add("nonce", strconv.FormatInt(time.Now().UnixNano()/1000000, 10))

	resp, err := client.Do(req)

	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(target)
}
