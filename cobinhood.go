package cobinhood

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"
)

const (
	apiBase = "https://api.cobinhood.com/v1/"
)

// Client holds your Cobinhood configuration
type Client struct {
	apiKey  string
	Timeout time.Duration
}

// NewClient creates a new cobinhood client
func NewClient(apiKey string, timeout time.Duration) *Client {
	return &Client{
		apiKey:  apiKey,
		Timeout: timeout}
}

// GetBalances returns a slice of all non-zero balances of your account
func (c *Client) GetBalances() ([]Balance, error) {

	result, err := c.request("GET", "wallet/balances", nil, true)
	if err != nil {
		return nil, err
	}

	balances := result.Balances
	if balances == nil {
		return nil, errors.New("balances is nil")
	}
	return *balances, nil
}

// GetTradingPairs returns all available trading pairs
func (c *Client) GetTradingPairs() ([]TradingPair, error) {

	result, err := c.request("GET", "market/trading_pairs", nil, false)
	if err != nil {
		return nil, err
	}

	tradingPairs := result.TradingPairs
	if tradingPairs == nil {
		return nil, errors.New("tradingPairs is nil")
	}

	return *tradingPairs, nil
}

// GetTicker returns a slice of type Ticker with the exchange/ticker information for the pair
func (c *Client) GetTicker(tradingPairID string) (Ticker, error) {

	result, err := c.request("GET", "market/tickers/"+tradingPairID, nil, false)
	if err != nil {
		return Ticker{}, err
	}

	ticker := result.Ticker
	if ticker == nil {
		return Ticker{}, errors.New("ticker is nil")
	}

	return *ticker, nil
}

// GetOpenOrders returns a slice of type OpenOrder with all your open orders at the exchange
func (c *Client) GetOpenOrders() ([]OpenOrder, error) {

	result, err := c.request("GET", "trading/orders", nil, true)
	if err != nil {
		return nil, err
	}

	openOrders := result.OpenOrders
	if openOrders == nil {
		return nil, errors.New("openOrders is nil")
	}

	return *openOrders, nil
}

// PlaceOrder places an order
func (c *Client) PlaceOrder(tradingPairID, side, tradeType string, price, size float64) (PlacedOrder, error) {
	request := &orderRequest{
		TradingPairID: tradingPairID,
		Side:          side,
		Type:          tradeType,
		Price:         strconv.FormatFloat(price, 'f', 8, 64),
		Size:          strconv.FormatFloat(size, 'f', 4, 64)}

	requestJSON, _ := json.Marshal(request)

	result, err := c.request("POST", "trading/orders", bytes.NewReader(requestJSON), true)
	if err != nil {
		return PlacedOrder{}, err
	}

	placedOrder := result.PlacedOrder
	if placedOrder == nil {
		return PlacedOrder{}, errors.New("placedOrder is nil")
	}

	return *placedOrder, nil
}

// CancelOrder cancels an order
func (c *Client) CancelOrder(orderID string) error {
	_, err := c.request("DELETE", "trading/orders/"+orderID, nil, true)
	return err
}

// CancelAllOrders cancels all open orders. It only returns the last error if any
func (c *Client) CancelAllOrders() error {
	openOrders, err := c.GetOpenOrders()
	if err != nil {
		return err
	}

	for _, openOrder := range openOrders {
		err = c.CancelOrder(openOrder.ID)
		// cancelling is more important than returning the error
	}

	return err
}

// GetOrderBook gets the orderbook for a market
func (c *Client) GetOrderBook(tradingPairID string, limit int) (OrderBook, error) {

	url := fmt.Sprintf("market/orderbooks/%s?limit=%d", tradingPairID, limit)
	result, err := c.request("GET", url, nil, false)
	if err != nil {
		return OrderBook{}, err
	}

	orderbook := result.OrderBook
	if orderbook == nil {
		return OrderBook{}, errors.New("orderbook is nil")
	}

	return *orderbook, nil
}

func (c *Client) request(method string, apiURL string, body io.Reader, private bool) (*GenericResult, error) {

	client := &http.Client{Timeout: c.Timeout}
	req, err := http.NewRequest(method, apiBase+apiURL, body)

	if private {
		if c.apiKey == "" {
			return nil, errors.New("API key cannot be empty for private requests")
		}

		req.Header.Add("Authorization", c.apiKey)
		req.Header.Add("nonce", strconv.FormatInt(time.Now().UnixNano()/1000000, 10))
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var response GenericResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, err
	}

	if response.Error != nil {
		return nil, errors.New(response.Error.ErrorCode)
	}

	return &response.GenericResult, nil
}
