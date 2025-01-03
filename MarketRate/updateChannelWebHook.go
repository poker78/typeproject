package marketRateType

type UpdateChannelWebHook struct {
	Name        string      `json:"name"`
	Symbol      string      `json:"symbol"`
	Price       float64     `json:"price"`
	PriceChange PriceChange `json:"priceChange"`
}

type PriceChange struct {
	IRT float64 `json:"irt"`
	EUR float64 `json:"eur"`
	JPY float64 `json:"jpy"`
	GBP float64 `json:"gbp"`
	ZAR float64 `json:"zar"`
}
