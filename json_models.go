package cobinhood

// Ticker holds the ticker information for a tradingpair
type Ticker struct {
	TradingPairID  string  `json:"trading_pair_id"`
	Timestamp      int64   `json:"timestamp"`
	Two4HHigh      float64 `json:"24h_high,string"`
	Two4HLow       float64 `json:"24h_low,string"`
	Two4HOpen      float64 `json:"24h_open,string"`
	Two4HVolume    float64 `json:"24h_volume,string"`
	LastTradePrice float64 `json:"last_trade_price,string"`
	HighestBid     float64 `json:"highest_bid,string"`
	LowestAsk      float64 `json:"lowest_ask,string"`
}

// OpenOrder holds information on open orders
type OpenOrder struct {
	ID            string  `json:"id"`
	TradingPairID string  `json:"trading_pair_id"`
	Side          string  `json:"side"`
	Type          string  `json:"type"`
	Price         float64 `json:"price,string"`
	Size          float64 `json:"size,string"`
	Filled        float64 `json:"filled,string"`
	State         string  `json:"stats"`
	Timestamp     int64   `json:"timestamp"`
	EqPrice       string  `json:"eq_price"`
	CompletedAt   string  `json:"completed_at"`
	TradingPair   string  `json:"trading_pair_id"`
}

type orderRequest struct {
	TradingPairID string `json:"trading_pair_id"`
	Side          string `json:"side"`
	Type          string `json:"type"`
	Price         string `json:"price"`
	Size          string `json:"size"`
}

// PlacedOrder contains the reply after placing an order
type PlacedOrder struct {
	ID          string  `json:"id"`
	TradingPair string  `json:"trading_pair_id"`
	State       string  `json:"state"`
	Side        string  `json:"side"`
	Type        string  `json:"completed_at"`
	Price       float64 `json:"price,string"`
	Size        float64 `json:"size,string"`
	Filled      float64 `json:"filled,string"`
	Timestamp   int64   `json:"timestamp"`
	EqPrice     string  `json:"eq_price"`
}

// Balance contains the balance for a particlar currency
type Balance struct {
	Currency string  `json:"currency"`
	Type     string  `json:"type"`
	Total    float64 `json:"total,string"`
	OnOrder  float64 `json:"on_order,string"`
	Locked   bool    `json:"locked"`
	UsdValue float64 `json:"usd_value,string"`
	BtcValue float64 `json:"btc_value,string"`
}

// TradingPair contains data regarding one trading pair
type TradingPair struct {
	BaseCurrency   string  `json:"base_currency_id"`
	MarginEnabled  bool    `json:"margin_enabled"`
	BaseMaxSize    float64 `json:"base_max_size,string"`
	QuoteIncrement float64 `json:"quote_increment,string"`
	QuoteCurrency  string  `json:"quote_currency_id"`
	ID             string  `json:"id"`
	BaseMinSize    float64 `json:"base_min_size,string"`
}

// Error contains an API error string
type Error struct {
	ErrorCode string `json:"error_code"`
}

// GenericResult contains any of the returned API values
type GenericResult struct {
	Balances     *[]Balance     `json:"balances"`
	TradingPairs *[]TradingPair `json:"trading_pairs"`
	Ticker       *Ticker        `json:"ticker"`
	OpenOrders   *[]OpenOrder   `json:"orders"`
	PlacedOrder  *PlacedOrder   `json:"order"`
}

// GenericResponse contains the result from any API call
type GenericResponse struct {
	Success       bool          `json:"success"`
	GenericResult GenericResult `json:"result"`
	Error         *Error        `json:"error"`
}
