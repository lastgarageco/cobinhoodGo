// Package cobinhoodgo implements the basic calls of cobinhood API
// defined by their documentation
package cobinhoodgo

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"log"
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

	var response openOrderResponse
	err := c.request("GET", "trading/orders", nil, &response)
	if err != nil {
		return nil, err
	}

	return response.Result.OpenOrders, nil
}

// PlaceOrder lets you place an order in the exchange
func (c *Cobin) PlaceOrder(tradingPairID, side, tradeType string, price, size float64) (PlacedOrder, error) {
	request := &orderRequest{
		TradingPairID: tradingPairID,
		Side:          side,
		Type:          tradeType,
		Price:         strconv.FormatFloat(price, 'f', 8, 64),
		Size:          strconv.FormatFloat(size, 'f', 4, 64)}

	requestJSON, _ := json.Marshal(request)

	log.Printf("requestJSON = %+v", string(requestJSON))

	var response placedOrderResponse
	err := c.request("POST", "trading/orders", bytes.NewReader(requestJSON), &response)
	if err != nil {
		return PlacedOrder{}, err
	}

	return response.Result.PlacedOrder, nil
}

// CancelOrder cancel an order
func (c *Cobin) CancelOrder(orderID string) error {

	var emptyStruct struct{}
	return c.request("DELETE", "trading/orders/"+orderID, nil, &emptyStruct)
}

// send requested acction to Cobinhood
func (c *Cobin) request(method string, apiURL string, body io.Reader, target interface{}) error {

	if c.apiKey == "" {
		return errors.New("Api Key can't be empty")
	}

	client := &http.Client{Timeout: time.Second}
	req, err := http.NewRequest(method, apiBase+apiURL, body)

	req.Header.Add("Authorization", c.apiKey)
	req.Header.Add("nonce", strconv.FormatInt(time.Now().UnixNano()/1000000, 10))

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	// TODO remove bodyBytes
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	//log.Printf("bodyBytes = %s", string(bodyBytes))

	bodyBuffer := bytes.NewBuffer(bodyBytes)

	defer resp.Body.Close()

	// TODO intercept failure in json value here

	return json.NewDecoder(bodyBuffer).Decode(target)
}
