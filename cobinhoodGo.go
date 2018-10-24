// Package cobinhoodgo implements the basic calls of cobinhood API
// defined by their documentation
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

const (
	apiBase = "https://api.cobinhood.com/v1/"
)

// Cobin holds your Cobinhood configuration
type Cobin struct {
	apiKey string
}

// SetAPIKey set your Coobinhood API Key
func (c *Cobin) SetAPIKey(key string) {
	c.apiKey = key
}

// GetBalances returns a slice of all non-zero balances of your account
func (c *Cobin) GetBalances() ([]Balance, error) {

	var response balancesResponse
	err := c.request("GET", "wallet/balances", nil, &response)
	if err != nil {
		return nil, err
	}

	return response.Result.Balances, nil
}

// GetTradingPairs returns all available trading pairs
func (c *Cobin) GetTradingPairs() ([]TradingPair, error) {

	var response tradingPairsResponse
	err := c.request("GET", "market/trading_pairs", nil, &response)
	if err != nil {
		return nil, err
	}

	return response.Result.TradingPairs, nil
}

// GetTicker returns a slice of type Ticker with the exchange/ticker information for the pair
func (c *Cobin) GetTicker(tradingPairID string) (Ticker, error) {

	var response tickerResponse
	err := c.request("GET", "market/tickers/"+tradingPairID, nil, &response)
	if err != nil {
		return Ticker{}, err
	}

	return response.Result.Ticker, nil
}

// GetOpenOrders returns a slice of type OpenOrder with all your open orders at the exchange
func (c *Cobin) GetOpenOrders() ([]OpenOrder, error) {
	var myOpenOrders []OpenOrder
	openOrders := openorders{}

	err := c.request("GET", "trading/orders", nil, &openOrders)
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
func (c *Cobin) GetOrderStatus(orderID string) (Order, error) {
	var myOrderStatus Order
	orderStatus := order{}

	err := c.request("GET", "trading/orders/"+orderID, nil, &orderStatus)
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

// CancelOrder cancel an order
func (c *Cobin) CancelOrder(orderID string) (bool, error) {
	status := false
	statusMsg := statusmessage{}

	err := c.request("DELETE", "trading/orders/"+orderID, nil, &statusMsg)
	if err != nil {
		return false, err
	}

	if statusMsg.Success == true {
		status = true
	}

	return status, nil
}

// PlaceOrder lets you place an order in the exchange
func (c *Cobin) PlaceOrder(po PlaceOrderData) (PlaceOrderResult, error) {
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

	err := c.request("POST", "trading/orders", bytes.NewReader(orderJSON), &placedOrder)
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

// send requested acction to Cobinhood
func (c *Cobin) request(postType string, apiURL string, body io.Reader, target interface{}) error {

	if c.apiKey == "" {
		return errors.New("Api Key can't be empty")
	}

	client := &http.Client{Timeout: 30 * time.Second}
	req, err := http.NewRequest(postType, apiBase+apiURL, body)

	req.Header.Add("Authorization", c.apiKey)
	req.Header.Add("nonce", strconv.FormatInt(time.Now().UnixNano()/1000000, 10))

	resp, err := client.Do(req)

	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// TODO intercept failure in json value here

	return json.NewDecoder(resp.Body).Decode(target)
}
