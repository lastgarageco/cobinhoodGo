package cobinhoodgo

// Ticker hold the ticker information for the coin pair
type Ticker struct {
	TradingPairID  string
	Timestamp      int64
	Two4HHigh      float64 `json:"24h_high,string"`
	Two4HLow       float64 `json:"24h_low,string"`
	Two4HOpen      float64 `json:"24h_open,string"`
	Two4HVolume    float64 `json:"24h_volume,string"`
	LastTradePrice float64
	HighestBid     float64 `json:"highest_bid,string"`
	LowestAsk      float64 `json:"lowest_ask,string"`
}

type tickerResponse struct {
	Success bool `json:"success"`
	Result  struct {
		Ticker Ticker `json:"ticker"`
	} `json:"result"`
}

// OpenOrder hold Open Orders information
type OpenOrder struct {
	ID            string
	TradingPairID string `json:"trading_pair_id"`
	Side          string
	Type          string
	Price         float64 `json:"price,string"`
	Size          float64 `json:"size,string"`
	Filled        float64 `json:"filled,string"`
	State         string
	Timestamp     int64
	EqPrice       string
	CompletedAt   string
	TradingPair   string
}

type openOrderResponse struct {
	Success bool `json:"success"`
	Result  struct {
		OpenOrders []OpenOrder `json:"orders"`
	} `json:"result"`
}

// Order to hold the information about an order
type OrderStatus struct {
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

type orderStatusResponse struct {
	Success bool `json:"success"`
	Result  struct {
		OrderStatus OrderStatus `json:"ticker"`
	} `json:"result"`
}

// PlaceOrderRequest contains data regarding a new order
type orderRequest struct {
	TradingPairID string `json:"trading_pair_id"`
	Side          string `json:"side"`
	Type          string `json:"type"`
	Price         string `json:"price"`
	Size          string `json:"size"`
}

// PlacedOrderResult holds the result for placing an order
type PlacedOrder struct {
	ID          string
	TradingPair string `json:"trading_pair_id"`
	State       string
	Side        string
	Type        string  `json:"completed_at"`
	Price       float64 `json:"price,string"`
	Size        float64 `json:"size,string"`
	Filled      float64 `json:"filled,string"`
	Timestamp   int64   `json:"timestamp"`
	EqPrice     string  `json:"eq_price"`
}

type placedOrderResponse struct {
	Success bool `json:"success"`
	Result  struct {
		PlacedOrder PlacedOrder `json:"order"`
	} `json:"result"`
}

type Balance struct {
	Currency string  `json:"currency"`
	Type     string  `json:"type"`
	Total    float64 `json:"total,string"`
	OnOrder  float64 `json:"on_order,string"`
	Locked   bool    `json:"locked"`
	UsdValue float64 `json:"usd_value,string"`
	BtcValue float64 `json:"btc_value,string"`
}

type balancesResponse struct {
	Success bool `json:"success"`
	Result  struct {
		Balances []Balance `json:"balances"`
	} `json:"result"`
}

type TradingPair struct {
	BaseCurrency   string  `json:"base_currency_id"`
	MarginEnabled  bool    `json:"margin_enabled"`
	BaseMaxSize    float64 `json:"base_max_size,string"`
	QuoteIncrement float64 `json:"quote_increment,string"`
	QuoteCurrency  string  `json:"quote_currency_id"`
	ID             string  `json:"id"`
	BaseMinSize    float64 `json:"base_min_size,string"`
}

type tradingPairsResponse struct {
	Success bool `json:"success"`
	Result  struct {
		TradingPairs []TradingPair `json:"trading_pairs"`
	} `json:"result"`
}

type StatusMessage struct {
	Success bool `json:"success"`
}
