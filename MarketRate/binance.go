package marketRateType


type BinanceCoin struct {
	EventType              string `json:"e"` // Event type
	EventTime              int64  `json:"E"` // Event time
	Symbol                 string `json:"s"` // Symbol
	PriceChange            string `json:"p"` // Price change
	PriceChangePercent     string `json:"P"` // Price change percent
	WeightedAveragePrice   string `json:"w"` // Weighted average price
	FirstTradePrice        string `json:"x"` // First trade price
	LastPrice              string `json:"c"` // Last price
	LastQuantity           string `json:"Q"` // Last quantity
	BestBidPrice           string `json:"b"` // Best bid price
	BestBidQuantity        string `json:"B"` // Best bid quantity
	BestAskPrice           string `json:"a"` // Best ask price
	BestAskQuantity        string `json:"A"` // Best ask quantity
	OpenPrice              string `json:"o"` // Open price
	HighPrice              string `json:"h"` // High price
	LowPrice               string `json:"l"` // Low price
	TotalTradedBaseVolume  string `json:"v"` // Total traded base volume
	TotalTradedQuoteVolume string `json:"q"` // Total traded quote volume
	StatisticsOpenTime     int64  `json:"O"` // Statistics open time
	StatisticsCloseTime    int64  `json:"C"` // Statistics close time
	FirstTradeID           int64  `json:"F"` // First trade ID
	LastTradeID            int64  `json:"L"` // Last trade ID
	TotalNumberOfTrades    int64  `json:"n"` // Total number of trades
}
