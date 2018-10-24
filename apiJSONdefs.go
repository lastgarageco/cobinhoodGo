package cobinhoodgo

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

type openorders struct {
	Success bool `json:"success"`
	Result  struct {
		Limit     int `json:"limit"`
		Page      int `json:"page"`
		TotalPage int `json:"total_page"`
		Orders    []struct {
			ID            string      `json:"id"`
			TradingPairID string      `json:"trading_pair_id"`
			Side          string      `json:"side"`
			Type          string      `json:"type"`
			Price         string      `json:"price"`
			Size          string      `json:"size"`
			Filled        string      `json:"filled"`
			State         string      `json:"state"`
			Timestamp     int64       `json:"timestamp"`
			EqPrice       string      `json:"eq_price"`
			CompletedAt   interface{} `json:"completed_at"`
			TradingPair   string      `json:"trading_pair"`
		} `json:"orders"`
	} `json:"result"`
}

type order struct {
	Success bool `json:"success"`
	Result  struct {
		Order struct {
			ID          string `json:"id"`
			TradingPair string `json:"trading_pair"`
			State       string `json:"state"`
			Side        string `json:"side"`
			Type        string `json:"type"`
			Price       string `json:"price"`
			Size        string `json:"size"`
			Filled      string `json:"filled"`
			Timestamp   int64  `json:"timestamp"`
		} `json:"order"`
	} `json:"result"`
}

type ticker struct {
	Success bool `json:"success"`
	Result  struct {
		Ticker struct {
			TradingPairID  string `json:"trading_pair_id"`
			Timestamp      int64  `json:"timestamp"`
			Two4HHigh      string `json:"24h_high"`
			Two4HLow       string `json:"24h_low"`
			Two4HOpen      string `json:"24h_open"`
			Two4HVolume    string `json:"24h_volume"`
			LastTradePrice string `json:"last_trade_price"`
			HighestBid     string `json:"highest_bid"`
			LowestAsk      string `json:"lowest_ask"`
		} `json:"ticker"`
	} `json:"result"`
}

type placeorder struct {
	TradingPairID string `json:"trading_pair_id"`
	Side          string `json:"side"`
	Type          string `json:"type"`
	Price         string `json:"price"`
	Size          string `json:"size"`
}

type placeorderresult struct {
	Success bool `json:"success"`
	Result  struct {
		Order struct {
			ID          string `json:"id"`
			TradingPair string `json:"trading_pair"`
			State       string `json:"state"`
			Side        string `json:"side"`
			Type        string `json:"type"`
			Price       string `json:"price"`
			Size        string `json:"size"`
			Filled      string `json:"filled"`
			Timestamp   int64  `json:"timestamp"`
			EqPrice     string `json:"eq_price"`
		} `json:"order"`
	} `json:"result"`
}

type statusmessage struct {
	Success bool `json:"success"`
}
