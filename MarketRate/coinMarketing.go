package marketRateType

type Quote struct {
	USD struct {
		Price float64 `json:"price"`
	} `json:"USD"`
}

type Coin struct {
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
	Slug   string `json:"slug"`
	Quote  Quote  `json:"quote"`
}

type CoinResponse struct {
	Status struct {
		Timestamp string `json:"timestamp"`
	} `json:"status"`
	Data []Coin `json:"data"`
}
